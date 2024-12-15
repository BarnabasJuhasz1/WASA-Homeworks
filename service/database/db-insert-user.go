package database

import (
	"sapienza/wasatext/service/api/util"
)

// InsertUser updates the user table by creating a new user entry
func (db *appdbimpl) InsertUser(newUser util.User) error {

	newUserName := newUser.Username
	newUserProfilePicture := newUser.ProfilePicture

	_, err := db.c.Exec("INSERT INTO users (username, profile_picture) VALUES (?, ?) ON CONFLICT(username) DO NOTHING", newUserName, newUserProfilePicture)

	return err
}
