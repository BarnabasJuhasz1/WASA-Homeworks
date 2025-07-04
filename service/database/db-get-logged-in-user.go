package database

import (
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"
)

func (db *appdbimpl) GetLoggedInUser(w http.ResponseWriter, ctx reqcontext.RequestContext) util.User {

	// double check to make sure user is authenticated
	if !ctx.Authenticated {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return util.User{}
	}

	var returnedUser util.User
	// make query to retrieve row from users table with the given username
	err := db.c.QueryRow("SELECT id, username, profile_picture FROM users WHERE username = ?", ctx.UserID).Scan(&returnedUser.Id, &returnedUser.Username, &returnedUser.ProfilePicture)

	if err != nil {
		ctx.Logger.Warningln("User with name: ", ctx.UserID, " requested but not found in database!")
		w.WriteHeader(http.StatusNotFound)
		return util.User{}
	}

	return returnedUser
}
