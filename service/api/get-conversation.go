package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func getConversation Called-----")

	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	// get the conversation from path
	Conversation, convErr := GetConversationFromPath(rt, w, ps, ctx)
	if convErr {
		return
	}

	// make sure the logged in user belongs to the conversation
	if !util.UserBelongsToConversation(Conversation, LoggedInUser) {
		ctx.Logger.Debugln("User is not in the conversation!")

		w.WriteHeader(http.StatusForbidden)
		return
	}

	// if logged in user is not part of the people who have seen the last messages
	// of the conversation, then add him to it
	for i := len(Conversation.Messages) - 1; i >= 0; i-- {
		message := Conversation.Messages[i]

		// if the message has not been read by everyone yet, and it has not been read by the logged in user
		if message.Status != util.DoubleCheckmark && !contains(message.ReadBy, LoggedInUser.Id) {

			message.ReadBy = append(message.ReadBy, LoggedInUser.Id)

			// if everyone has read this message, flag it with double checkmark
			if len(message.ReadBy) == len(Conversation.Participants) {
				message.Status = util.DoubleCheckmark
			}

			// replace old message in conversation with new one
			Conversation.Messages[i] = message

		} else {
			// if we found a message that has been read by the player then break,
			// because everything before must have also been read by the player
			break
		}

	}

	// update conversation
	dberr := rt.db.UpdateConversation(Conversation.Id, Conversation)
	if dberr != nil {
		ctx.Logger.Errorln("Failed to update conversation:", dberr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encodeErr := json.NewEncoder(w).Encode(Conversation)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func getConversation Finished-----")
}

// function to check if an int ID is present in a list of int IDs
func contains(userConversationIds []int, conversationId int) bool {
	for _, item := range userConversationIds {
		if item == conversationId {
			return true
		}
	}
	return false
}
