package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func setMyPhoto Called-----")
	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	// Read the request body
	var requestBody struct {
		ProfilePicture []byte `json:"ProfilePicture"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// UserLoggedIn := LoggedInUser
	// if UserLoggedIn.ProfilePicture == requestBody.ProfilePicture {
	// 	ctx.Logger.Debugln("Your profile picture is already ", UserLoggedIn.ProfilePicture)

	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	LoggedInUser.ProfilePicture = requestBody.ProfilePicture
	rt.db.UpdateUser(LoggedInUser, LoggedInUser.Username)
	// util.AllUsers[UserLoggedIn.Username] = UserLoggedIn

	ctx.Logger.Debugln("-----Func setMyPhoto Finished-----")

	encodeErr := json.NewEncoder(w).Encode(LoggedInUser)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
