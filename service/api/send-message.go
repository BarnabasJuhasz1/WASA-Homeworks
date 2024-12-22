package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func sendMessage Called-----")
	// LoggedInUser := util.GetLoggedInUser(w, ctx, rt.db)
	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	conversationIDString := ps.ByName("ConversationID")

	conversationID, err := strconv.Atoi(conversationIDString)
	if err != nil || conversationID < 0 {
		ctx.Logger.Debugln("Invalid conversationID in path! ", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Conversation, existsConv := util.AllConversations[conversationID]
	// Conversation, convErr := rt.db.GetConversation(conversationID)
	// if convErr != nil {
	// 	ctx.Logger.Debugln("Invalid conversationID in path! ", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	myConversationIDs, convErr := rt.db.GetMyConversationIDs(LoggedInUser.Id)
	if convErr != nil {
		ctx.Logger.Debugln("Invalid conversationID in path! ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// make sure the user asks for a conversation it is part of
	// if !contains(LoggedInUser.MyConversations, conversationID) {
	ctx.Logger.Debugln("checking if: ", myConversationIDs, " contains ", conversationID)

	if !contains(myConversationIDs, conversationID) {
		ctx.Logger.Debugln("User tried to send a message to a conversation it is not part of! ", err)

		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Read the request body
	var requestBody struct {
		// SenderUsername string `json:"SenderUsername"`
		MessageContent string `json:"MessageContent"`
	}

	err2 := json.NewDecoder(r.Body).Decode(&requestBody)
	if err2 != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	if len(requestBody.MessageContent) == 0 || len(requestBody.MessageContent) > 10000 {
		ctx.Logger.Debugln("Message too long or too short!")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// query to get the actual conversation struct
	Conversation, convErr2 := rt.db.GetConversation(conversationID)
	if convErr2 != nil {
		ctx.Logger.Debugln("Invalid conversationID in path! ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SenderUser, existsUser := Users[requestBody.SenderUsername]
	// if !existsUser {
	// 	fmt.Println("User ", requestBody.SenderUsername, " is not in the database!")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	var emptyReactions []util.Reaction

	// modify conversation by adding the new message
	Conversation.Messages = append(Conversation.Messages, util.Message{
		Id:              len(Conversation.Messages),
		Sender:          LoggedInUser.Id,
		Content:         requestBody.MessageContent,
		Timestamp:       time.Now().Format("2006-01-02 15:04:05"),
		Status:          util.UserName,
		EmojiReactions:  emptyReactions,
		OriginMessageId: -1,
	})

	// update the allConversations map by reassigning the struct
	// util.AllConversations[conversationID] = Conversation

	ctx.Logger.Debugln("id: ", conversationID, " participants: ", Conversation.Participants)

	updateConvErr := rt.db.UpdateConversation(conversationID, Conversation)
	if updateConvErr != nil {
		ctx.Logger.Errorln("Error updating conversation: ", updateConvErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func sendMessage Finished-----")

	encodeErr := json.NewEncoder(w).Encode(Conversation)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
