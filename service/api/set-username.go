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
	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)
	ctx.Logger.Debugln("User trying to change name :", LoggedInUser.Username)

	// Read the request body
	var requestBody struct {
		NewUsername string `json:"NewUsername"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if LoggedInUser.Username == requestBody.NewUsername {
		ctx.Logger.Debugln("Your username is already ", LoggedInUser.Username)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(requestBody.NewUsername) < 3 || len(requestBody.NewUsername) > 16 {
		ctx.Logger.Debugln("Username too short or too long ", requestBody.NewUsername)

		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	_, userExistsError := rt.db.GetUserFromName(requestBody.NewUsername)

	if userExistsError == nil {
		ctx.Logger.Debugln("Username ", requestBody.NewUsername, " occupied by someone else!")

		w.WriteHeader(http.StatusNotAcceptable)
		return
	} else {
		oldUserName := LoggedInUser.Username
		// update the username of the user logged in
		// Remark: since UserLoggedIn is a pointer, the "AllUsers" map is updated as well
		LoggedInUser.Username = requestBody.NewUsername
		ctx.UserID = requestBody.NewUsername
		updateErr := util.UpdateUsername(oldUserName, requestBody.NewUsername)
		if updateErr != nil {
			rt.baseLogger.Errorln("Error updating token of user after name change! ", updateErr)
		}
		// add same user with the new name
		// util.AllUsers[requestBody.NewUsername] = LoggedInUser
		// delete old entry in users
		// delete(util.AllUsers, oldUserName)

		dberr := rt.db.UpdateUser(LoggedInUser, oldUserName)
		if dberr != nil {
			rt.baseLogger.Errorln("Saving new User into DB error! ", dberr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	ctx.Logger.Debugln("-----Func setUsername Finished-----")

	encodeErr := json.NewEncoder(w).Encode(LoggedInUser.Username)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
