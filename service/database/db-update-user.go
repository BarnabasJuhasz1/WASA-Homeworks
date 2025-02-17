package database

import (
	"sapienza/wasatext/service/api/util"
)

// UpdateUser updates the user table entry if the existing user successfully changed its username or icon
func (db *appdbimpl) UpdateUser(newUser util.User, existingUsername string) error {

	// create primitive types to hold the values of the user struct
	newUserName := newUser.Username
	newUserProfilePicture := newUser.ProfilePicture

	// make query to update the username of the user in the users table
	_, err := db.c.Exec("UPDATE users SET username = ?, profile_picture = ? WHERE id = ?", newUserName, newUserProfilePicture, newUser.Id)
	if err != nil {
		return err
	}

	return err
}
