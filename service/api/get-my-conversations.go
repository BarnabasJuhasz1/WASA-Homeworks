package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func getMyConversations Called-----")

	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	// // get the conversations of a user
	// var myConversations []util.Conversation
	// for _, conversationId := range LoggedInUser.MyConversations {

	// 	myConversations = append(myConversations, util.AllConversations[conversationId])
	// }
	myConversations, err := rt.db.GetMyConversations(LoggedInUser.Username)
	if err != nil {
		ctx.Logger.Errorln("Failed to retrieve user conversations!", err)
	}

	ctx.Logger.Debugln("-----Func getMyConversations Finished-----")
	encodeErr := json.NewEncoder(w).Encode(myConversations)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
