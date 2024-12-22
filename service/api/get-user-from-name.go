package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserFromName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func getConversation Called-----")

	// LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	// Retrieve the "Username" query parameter from the URL
	username := r.URL.Query().Get("Username")
	if username == "" {
		http.Error(w, "UserID is required", http.StatusBadRequest)
		return
	}

	userToReturn, userExistsError := rt.db.GetUserFromName(username)
	if userExistsError != nil {
		ctx.Logger.Debugln("User ", username, " is not in the database!")

		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx.Logger.Debugln("-----Func getConversation Finished-----")

	encodeErr := json.NewEncoder(w).Encode(userToReturn)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
