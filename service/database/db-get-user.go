package database

import (
	"sapienza/wasatext/service/api/util"
)

// GetUser returns the user with the given userID from the database if it exists
func (db *appdbimpl) GetUser(userID int) (util.User, error) {

	var returnedUser util.User

	err := db.c.QueryRow("SELECT id, username, profile_picture FROM users WHERE id = ?", userID).Scan(&returnedUser.Id, &returnedUser.Username, &returnedUser.ProfilePicture)

	return returnedUser, err
}

// GetUserFromName returns the user with the given name from the database if it exists
func (db *appdbimpl) GetUserFromName(username string) (util.User, error) {

	var returnedUser util.User

	err := db.c.QueryRow("SELECT id, username, profile_picture FROM users WHERE username = ?", username).Scan(&returnedUser.Id, &returnedUser.Username, &returnedUser.ProfilePicture)

	return returnedUser, err
}

// GetAllUsers returns the userIDs in the database
func (db *appdbimpl) GetAllUsers() ([]int, error) {
	var userIDs []int

	rows, err := db.c.Query("SELECT id FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		rowErr := rows.Scan(&id)
		if rowErr != nil {
			return nil, rowErr
		}
		userIDs = append(userIDs, id)
	}

	err2 := rows.Err()
	if err2 != nil {
		return nil, err2
	}

	return userIDs, nil
}
