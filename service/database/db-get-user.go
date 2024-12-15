package database

import (
	"sapienza/wasatext/service/api/util"
)

// GetUser returns the user with the given name from the database if it exists
func (db *appdbimpl) GetUser(username string) (util.User, error) {

	var returnedUser util.User

	err := db.c.QueryRow("SELECT username, profile_picture FROM users WHERE username = ?", username).Scan(&returnedUser.Username, &returnedUser.ProfilePicture)

	return returnedUser, err
}
