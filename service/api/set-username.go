package api

import (
	"encoding/json"
	"net/http"

	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func setUsername Called-----")

	// Read the request body
	var requestBody struct {
		NewUsername string `json:"NewUsername"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if util.GetLoggedInUser(w, ctx).Username == requestBody.NewUsername {
		ctx.Logger.Debugln("Your username is already ", util.GetLoggedInUser(w, ctx).Username)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(requestBody.NewUsername) < 3 || len(requestBody.NewUsername) > 16 {
		ctx.Logger.Debugln("Username too short or too long ", requestBody.NewUsername)

		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	_, newExists := util.AllUsers[requestBody.NewUsername]
	UserLoggedIn := util.GetLoggedInUser(w, ctx)

	if newExists {
		ctx.Logger.Debugln("Username ", requestBody.NewUsername, " occupied by someone else!")

		w.WriteHeader(http.StatusNotAcceptable)
		return
	} else {
		oldUserName := UserLoggedIn.Username
		// update the username of the user logged in
		// Remark: since UserLoggedIn is a pointer, the "AllUsers" map is updated as well
		UserLoggedIn.Username = requestBody.NewUsername
		ctx.UserID = requestBody.NewUsername
		// add same user with the new name
		util.AllUsers[requestBody.NewUsername] = UserLoggedIn
		// delete old entry in users
		delete(util.AllUsers, oldUserName)

		// fmt.Println("User ", UserLoggedIn.Username, " renamed sucessfully to ", requestBody.NewUsername, "!")
	}

	ctx.Logger.Debugln("-----Func setUsername Finished-----")

	encodeErr := json.NewEncoder(w).Encode(UserLoggedIn)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
