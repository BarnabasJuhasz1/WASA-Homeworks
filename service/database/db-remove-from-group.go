package database

import (
	"encoding/json"
)

// RemoveConversationIDFromUser updates the users table conversations column which maps users to conversation IDs
func (db *appdbimpl) RemoveConversationIDFromUser(userID int, removeFromConversationId int) error {

	var oldConversationIDsJson []byte

	// fmt.Println("trying to remove user: ", username, " from conversation: ", addToConversationId)

	// make query to get the current conversation IDs of the given user
	err := db.c.QueryRow("SELECT conversations FROM users WHERE id = ?", userID).Scan(&oldConversationIDsJson)
	if err != nil {
		return err
	}

	var conversationIDs []int
	// deserialize the []int slice (the old conversationIDs)
	if jsonErr := json.Unmarshal(oldConversationIDsJson, &conversationIDs); jsonErr != nil {
		return jsonErr
	}

	var newConversationIDs []int
	// create the new conversationID list without the ID we wanted to remove
	for _, id := range conversationIDs {
		if id != removeFromConversationId {
			newConversationIDs = append(newConversationIDs, id)
		}
	}

	// serialize the []int slice (the newConversationIDs)
	conversationIDsJson, jsonErr := json.Marshal(newConversationIDs)
	if jsonErr != nil {
		return jsonErr
	}

	// make query to update the conversationIDs of the user
	_, err2 := db.c.Exec("UPDATE users SET conversations = ? WHERE id = ?", conversationIDsJson, userID)

	return err2
}
