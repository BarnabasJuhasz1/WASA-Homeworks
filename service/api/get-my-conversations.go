package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func getMyConversations Called-----")

	// get the conversations of a user
	var myConversations []util.Conversation
	for _, conversationId := range util.GetLoggedInUser(w, ctx).MyConversations {

		myConversations = append(myConversations, util.AllConversations[conversationId])
	}

	ctx.Logger.Debugln("-----Func getMyConversations Finished-----")
	json.NewEncoder(w).Encode(myConversations)

}
