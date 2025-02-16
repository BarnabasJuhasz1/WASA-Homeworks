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
	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	conversationIDString := ps.ByName("ConversationID")

	// make sure conversation ID is a non-negative integer
	conversationID, err := strconv.Atoi(conversationIDString)
	if err != nil || conversationID < 0 {
		ctx.Logger.Debugln("Invalid conversationID in path! ", err)

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

	// make sure the logged in user belongs to the conversation
	if !util.UserBelongsToConversation(Conversation, LoggedInUser) {
		ctx.Logger.Debugln("User is not part of the conversation!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Read the request body
	var requestBody struct {
		MessageContent string `json:"MessageContent"`
	}

	// decode body
	err2 := json.NewDecoder(r.Body).Decode(&requestBody)
	if err2 != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	// make sure message is not empty
	if len(requestBody.MessageContent) == 0 {
		ctx.Logger.Debugln("Message is too short!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var emptyReactions []util.Reaction
	// message read by the sender automatically
	newMessageReadBy := []int{LoggedInUser.Id}

	// create the new message struct
	newMessage := util.Message{
		Id:              len(Conversation.Messages),
		Sender:          LoggedInUser.Id,
		Content:         requestBody.MessageContent,
		Timestamp:       time.Now().Format("2006-01-02 15:04:05"),
		Status:          util.SingleCheckmark,
		EmojiReactions:  emptyReactions,
		OriginMessageId: -1,
		HasBeenDeleted:  false,
		WasForwarded:    false,
		ReadBy:          newMessageReadBy,
	}
	// modify conversation by adding the new message
	Conversation.Messages = append(Conversation.Messages, newMessage)

	// update conversation in db
	dberr := rt.db.UpdateConversation(conversationID, Conversation)
	if dberr != nil {
		ctx.Logger.Errorln("Failed to update conversation:", dberr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// encode response
	encodeErr := json.NewEncoder(w).Encode(newMessage)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func sendMessage Finished-----")

}
