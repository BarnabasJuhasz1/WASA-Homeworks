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

	// make db query to get the conversations of a user
	myConversations, err := rt.db.GetMyConversations(LoggedInUser.Id)
	if err != nil {
		ctx.Logger.Errorln("Failed to retrieve user conversations!", err)
	}

	// encode response
	encodeErr := json.NewEncoder(w).Encode(myConversations)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func getMyConversations Finished-----")
}
