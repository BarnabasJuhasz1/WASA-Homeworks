package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func getConversation Called-----")

	// LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	// Retrieve the "Username" query parameter from the URL
	userIDStr := r.URL.Query().Get("UserID")
	if userIDStr == "" {
		http.Error(w, "UserID is required", http.StatusBadRequest)
		return
	}
	// Convert the string to an integer
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid UserID: must be an integer", http.StatusBadRequest)
		return
	}

	userToReturn, userExistsError := rt.db.GetUser(userID)
	if userExistsError != nil {
		ctx.Logger.Debugln("User ", userID, " is not in the database!")

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
