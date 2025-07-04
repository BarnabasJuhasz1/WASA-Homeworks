package database

import (
	"encoding/json"
)

// GetMyConversations gets all the conversation of a given user
func (db *appdbimpl) GetMyConversationIDs(userID int) ([]int, error) {

	var returnedMyConversationIDs []int

	// get the conversation IDs the user belongs to
	var myConversationIDsJson []byte

	// make query to get all conversation IDs corresponding to given userID
	err := db.c.QueryRow("SELECT conversations FROM users WHERE id = ?", userID).Scan(&myConversationIDsJson)
	if err != nil {
		return returnedMyConversationIDs, err
	}

	// deserialize the []int slice (the myConversationIDs)
	if jsonErr := json.Unmarshal(myConversationIDsJson, &returnedMyConversationIDs); jsonErr != nil {
		return returnedMyConversationIDs, jsonErr
	}

	return returnedMyConversationIDs, nil
}
