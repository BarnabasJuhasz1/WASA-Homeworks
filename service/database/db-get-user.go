package database

import (
	"sapienza/wasatext/service/api/util"
)

// GetUser returns the user with the given name from the database if it exists
func (db *appdbimpl) GetUser(userID int) (util.User, error) {

	var returnedUser util.User

	err := db.c.QueryRow("SELECT id, username, profile_picture FROM users WHERE id = ?", userID).Scan(&returnedUser.Id, &returnedUser.Username, &returnedUser.ProfilePicture)

	return returnedUser, err
}

// GetUser returns the user with the given name from the database if it exists
func (db *appdbimpl) GetUserFromName(username string) (util.User, error) {

	var returnedUser util.User

	err := db.c.QueryRow("SELECT id, username, profile_picture FROM users WHERE username = ?", username).Scan(&returnedUser.Id, &returnedUser.Username, &returnedUser.ProfilePicture)

	return returnedUser, err
}
