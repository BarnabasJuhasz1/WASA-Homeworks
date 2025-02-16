package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func addToGroup Called-----")

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

	// Read the request body
	var requestBody struct {
		UserIDToAdd int `json:"UserIDToAdd"`
	}

	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// make query to retrieve user to add and to make sure user exists
	userToAdd, userExistsError := rt.db.GetUser(requestBody.UserIDToAdd)
	if userExistsError != nil {
		ctx.Logger.Debugln("User ", requestBody.UserIDToAdd, " is not in the database!")

		w.WriteHeader(http.StatusNotFound)
		return
	}

	// make sure user to add does not already belong to the conversation
	if util.UserBelongsToConversation(Conversation, userToAdd) {
		ctx.Logger.Debugln("User is already in the conversation!")

		w.WriteHeader(http.StatusForbidden)
		return
	}

	// add user to the group (group's perspective)
	Conversation.Participants = append(Conversation.Participants, userToAdd.Id)

	// make query to add user to the group (user's perspective)
	dbAddErr := rt.db.AddConversationIDToUser(userToAdd.Id, Conversation.Id)
	if dbAddErr != nil {
		ctx.Logger.Errorln("Failed to add user to conversation:", dbAddErr)
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
	encodeErr := json.NewEncoder(w).Encode(Conversation)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func addToGroup Finished-----")
}
