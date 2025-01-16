package database

import (
	"sapienza/wasatext/service/api/util"
)

// UpdateUser updates the user table entry if the existing user successfully changed its username or icon
func (db *appdbimpl) UpdateUser(newUser util.User, existingUsername string) error {

	newUserName := newUser.Username
	newUserProfilePicture := newUser.ProfilePicture

	_, err := db.c.Exec("UPDATE users SET username = ?, profile_picture = ? WHERE id = ?", newUserName, newUserProfilePicture, newUser.Id)
	if err != nil {
		return err
	}

	// if the user changed its username
	// if newUserName != existingUsername {
	// 	// we also need to change the username where this user is stored as a participant of conversations
	// 	myConvs, err2 := db.GetMyConversations(newUserName)
	// 	if err2 != nil {
	// 		return err2
	// 	}

	// 	// loop through all the conversations of the user
	// 	for _, myConv := range myConvs {

	// 		// replace the old username with the new one
	// 		for j, convUsername := range myConv.Participants {
	// 			if convUsername == existingUsername {
	// 				myConv.Participants[j] = newUserName
	// 			}
	// 		}

	// 		// update the conversation in the database
	// 		db.UpdateConversation(myConv.Id, myConv)
	// 	}
	// }

	return err
}
