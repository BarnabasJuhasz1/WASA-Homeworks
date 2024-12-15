package database

import (
	"sapienza/wasatext/service/api/util"
)

// UpdateUser updates the user table entry if the existing user successfully changed its username or icon
func (db *appdbimpl) UpdateUser(newUser util.User, existingUsername string) error {

	newUserName := newUser.Username
	newUserProfilePicture := newUser.ProfilePicture

	_, err := db.c.Exec("UPDATE users SET username = ?, profile_picture = ? WHERE username = ?", newUserName, newUserProfilePicture, existingUsername)

	return err
}
