package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func setGroupPhoto Called-----")

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

	if Conversation.Type == util.UserType {
		ctx.Logger.Debugln("Cannot change photo of one-on-one conversations!")

		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Read the request body
	var requestBody struct {
		GroupPicture string `json:"GroupPicture"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if Conversation.ConversationGroup.GroupPicture == requestBody.GroupPicture {
		ctx.Logger.Debugln("The group picture is already set to this picture.")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// change the group picture to the payload picture
	Conversation.ConversationGroup.GroupPicture = requestBody.GroupPicture

	// update conversations map by reassigning the struct
	util.AllConversations[Conversation.Id] = Conversation

	ctx.Logger.Debugln("-----Func setGroupPhoto Finished-----")

	json.NewEncoder(w).Encode(Conversation)

}
