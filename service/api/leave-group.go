package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) leaveGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func leaveGroup Called-----")
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

	// delete user from group (group's perspective)
	Conversation.Participants = deleteIDFromList(Conversation.Participants, LoggedInUser.Id)

	// make query to delete user from group (users's perspective)
	removeErr := rt.db.RemoveConversationIDFromUser(LoggedInUser.Id, Conversation.Id)
	if removeErr != nil {
		ctx.Logger.Errorln("Failed to remove user from conversation:", removeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// update db
	dberr := rt.db.UpdateConversation(Conversation.Id, Conversation)
	if dberr != nil {
		ctx.Logger.Errorln("Failed to update conversation:", dberr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// encode response
	encodeErr := json.NewEncoder(w).Encode(LoggedInUser)
	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func leaveGroup Finished-----")
}

// function to delete a given ID from a list of IDs
func deleteIDFromList(userIDs []int, userIDToDelete int) []int {
	result := []int{}
	for _, userID := range userIDs {
		if userID != userIDToDelete {
			result = append(result, userID)
		}
	}
	return result
}
