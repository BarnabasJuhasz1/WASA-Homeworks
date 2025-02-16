package database

import (
	"encoding/json"
	"sapienza/wasatext/service/api/util"
)

// InsertUser updates the user table by creating a new user entry
func (db *appdbimpl) InsertUser(newUser util.User) (int, error) {

	// create primitive types to store values from the user struct
	newUserName := newUser.Username
	newUserProfilePicture := newUser.ProfilePicture
	// create empty conversations slice
	newConversations := []int{}

	// convert the slice to JSON
	newConversationsJSON, err := json.Marshal(newConversations)
	if err != nil {
		return -1, err
	}

	// make query to insert the new user to the users table
	result, err := db.c.Exec("INSERT INTO users (username, profile_picture, conversations) VALUES (?, ?, ?) ON CONFLICT(id) DO NOTHING", newUserName, newUserProfilePicture, newConversationsJSON)
	if err != nil {
		return -1, err
	}

	// retrieve the last inserted row ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), err
}
