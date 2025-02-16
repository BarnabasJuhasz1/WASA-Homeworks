package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Handles the request to uncomment a message (remove the emoji reaction you placed on a message)
func (rt *_router) uncommentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func uncommentMessage Called-----")
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

	// get messageID from path
	messageIDString := ps.ByName("MessageID")
	// make sure the messageID is correct
	messageID, messageErr := strconv.Atoi(messageIDString)
	if messageErr != nil || messageID < 0 || messageID >= len(Conversation.Messages) {
		ctx.Logger.Debugln("Invalid MessageID in path! ", messageErr)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// loop through all the emoji reactions to this message
	for i, ReactionAti := range Conversation.Messages[messageID].EmojiReactions {

		// if the user has a reaction to this message already, replace that reaction with the new one
		if ReactionAti.UserWhoReacted == LoggedInUser.Id {

			// remove your emoji reactions from the list of emoji reactions
			Conversation.Messages[messageID].EmojiReactions = append(Conversation.Messages[messageID].EmojiReactions[:i], Conversation.Messages[messageID].EmojiReactions[i+1:]...)
			break

			// if we checked all the reactions and none of them were created by the user
		} else if i == len(Conversation.Messages[messageID].EmojiReactions)-1 {

			ctx.Logger.Debugln("There is no emoji reaction placed by user on the specified message!")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// update db
	dberr := rt.db.UpdateConversation(Conversation.Id, Conversation)
	if dberr != nil {
		ctx.Logger.Errorln("Failed to update conversation:", dberr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// encode response
	encodeErr := json.NewEncoder(w).Encode(Conversation.Messages[messageID])

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func uncommentMessage Finished-----")
}
