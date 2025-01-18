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

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func forwardMessage Called-----")

	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	// get the conversation from path
	OriginalConversation, convErr := GetConversationFromPath(rt, w, ps, ctx)
	if convErr {
		ctx.Logger.Errorln("Error getting the conversation from path!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// make sure the logged in user belongs to the conversation
	if !util.UserBelongsToConversation(OriginalConversation, LoggedInUser) {
		ctx.Logger.Debugln("User is not in the conversation!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// get messageID from path
	messageIDString := ps.ByName("MessageID")
	// make sure the messageID is correct
	messageID, messageErr := strconv.Atoi(messageIDString)
	if messageErr != nil || messageID < 0 || messageID >= len(OriginalConversation.Messages) {
		ctx.Logger.Debugln("Invalid MessageID in path! ", messageErr)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read the request body
	var requestBody struct {
		ForwardToConvID int `json:"ForwardToConvID"`
	}
	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// make sure the recipient exists in the database as a user
	RecipientConversation, dbErr := rt.db.GetConversation(requestBody.ForwardToConvID)
	if dbErr != nil {
		ctx.Logger.Debugln("Error trying to retrieve recipient ConversationID!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var emptyReactions []util.Reaction
	// send the message to the recipient by creating a new message with the same content of the original message
	// and modify one-on-one conversation by adding the new message
	RecipientConversation.Messages = append(RecipientConversation.Messages, util.Message{
		Id:              len(RecipientConversation.Messages),
		Sender:          LoggedInUser.Id,
		Content:         OriginalConversation.Messages[messageID].Content,
		Timestamp:       time.Now().Format("2006-01-02 15:04:05"),
		Status:          util.SingleCheckmark,
		EmojiReactions:  emptyReactions,
		OriginMessageId: -1,
	})

	// update conversations map by reassigning the struct
	// util.AllConversations[ConvWithRecipient.Id] = ConvWithRecipient
	dberr := rt.db.UpdateConversation(RecipientConversation.Id, RecipientConversation)
	if dberr != nil {
		ctx.Logger.Errorln("Failed to update conversation:", dberr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encodeErr := json.NewEncoder(w).Encode(OriginalConversation.Messages[messageID])

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func forwardMessage Finished-----")

}
