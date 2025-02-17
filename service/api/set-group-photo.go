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

	// make sure user did not try to change the picture of a one-on-one conversation
	if Conversation.Type == util.UserType {
		ctx.Logger.Debugln("Cannot change photo of one-on-one conversations!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Read the request body
	var requestBody struct {
		GroupPicture []byte `json:"GroupPicture"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// change the group picture to the payload picture
	Conversation.GroupPicture = requestBody.GroupPicture

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

	ctx.Logger.Debugln("-----Func setGroupPhoto Finished-----")
}
