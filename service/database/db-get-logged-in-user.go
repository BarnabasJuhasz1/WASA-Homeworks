package database

import (
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"
)

func (db *appdbimpl) GetLoggedInUser(w http.ResponseWriter, ctx reqcontext.RequestContext) util.User {

	if !ctx.Authenticated {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return util.User{}
	}

	var returnedUser util.User
	err := db.c.QueryRow("SELECT username, profile_picture FROM users WHERE username = ?", ctx.UserID).Scan(&returnedUser.Username, &returnedUser.ProfilePicture)

	if err != nil {
		ctx.Logger.Warningln("User with ID: ", ctx.UserID, " requested but not found in database!")
		w.WriteHeader(http.StatusNotFound)
		return util.User{}
	}

	return returnedUser
}
