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

	// update profile picture
	LoggedInUser.ProfilePicture = requestBody.ProfilePicture

	// update db
	dberr := rt.db.UpdateUser(LoggedInUser, LoggedInUser.Username)
	if dberr != nil {
		rt.baseLogger.Errorln("Saving new User into DB error! ", dberr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// encode response
	encodeErr := json.NewEncoder(w).Encode(LoggedInUser.ProfilePicture)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func setMyPhoto Finished-----")
}
