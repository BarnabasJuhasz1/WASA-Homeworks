package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func setMyPhoto Called-----")

	// Read the request body
	var requestBody struct {
		ProfilePicture string `json:"ProfilePicture"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	UserLoggedIn := util.GetLoggedInUser(w, ctx)
	if UserLoggedIn.ProfilePicture == requestBody.ProfilePicture {
		ctx.Logger.Debugln("Your profile picture is already ", UserLoggedIn.ProfilePicture)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	UserLoggedIn.ProfilePicture = requestBody.ProfilePicture
	util.AllUsers[UserLoggedIn.Username] = UserLoggedIn

	ctx.Logger.Debugln("-----Func setMyPhoto Finished-----")

	encodeErr := json.NewEncoder(w).Encode(UserLoggedIn)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
