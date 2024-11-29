package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func deleteMessage Called-----")

	// get the conversation from path
	Conversation, convErr := util.GetConversationFromPath(w, ps, ctx)
	if convErr {
		return
	}

	// make sure the logged in user belongs to the conversation
	if !util.UserBelongsToConversation(Conversation, util.GetLoggedInUser(w, ctx)) {
		ctx.Logger.Debugln("User is not in the conversation!")

		w.WriteHeader(http.StatusForbidden)
		return
	}

	// get messageID from path
	messageIDString := ps.ByName("MessageID")
	// make sure the messageID is correct
	messageID, messageErr := strconv.Atoi(messageIDString)
	if messageErr != nil || messageID < 0 || messageID > len(Conversation.Messages) {
		ctx.Logger.Debugln("Invalid MessageID in path! ", messageErr)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// make sure you can only delete messages sent by you
	if Conversation.Messages[messageID].Sender.Username != util.GetLoggedInUser(w, ctx).Username {
		ctx.Logger.Debugln("Tried to delete a message not sent by you!")

		w.WriteHeader(http.StatusForbidden)
		return
	}

	// delete message content
	Conversation.Messages[messageID].Content = "Message content was deleted."
	// update conversations map by reassigning the struct
	util.AllConversations[Conversation.Id] = Conversation

	ctx.Logger.Debugln("-----Func deleteMessage Finished-----")

	json.NewEncoder(w).Encode(Conversation)

}
