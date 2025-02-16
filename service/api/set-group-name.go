package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func setGroupName Called-----")
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

	// make sure user did not try to change the name of a one-on-one conversation
	if Conversation.Type == util.UserType {
		ctx.Logger.Debugln("Cannot change name of one-on-one conversations!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Read the request body
	var requestBody struct {
		GroupName string `json:"GroupName"`
	}

	err2 := json.NewDecoder(r.Body).Decode(&requestBody)
	if err2 != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// make sure the new group name respects the length boundaries
	if len(requestBody.GroupName) < 3 || len(requestBody.GroupName) > 16 {
		ctx.Logger.Debugln("Group name too short or too long ", requestBody.GroupName)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// if then new group name is the same as the old one
	if Conversation.GroupName == requestBody.GroupName {
		ctx.Logger.Debugln("New Group name matches old one! ", requestBody.GroupName)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// modify conversation by changing the group name
	Conversation.GroupName = requestBody.GroupName
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

	ctx.Logger.Debugln("-----Func setGroupName Finished-----")
}
