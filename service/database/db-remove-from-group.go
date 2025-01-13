package database

import (
	"encoding/json"
)

// RemoveConversationIDFromUser updates the users table conversations column which maps users to conversation IDs
func (db *appdbimpl) RemoveConversationIDFromUser(userID int, removeFromConversationId int) error {

	var oldConversationIDsJson []byte

	// fmt.Println("trying to remove user: ", username, " from conversation: ", addToConversationId)
	// get the current conversation IDs
	//err := db.c.QueryRow("SELECT conversations FROM user_to_conversations WHERE username = ?", username).Scan(&oldConversationIDsJson)
	err := db.c.QueryRow("SELECT conversations FROM users WHERE id = ?", userID).Scan(&oldConversationIDsJson)
	if err != nil {
		return err
	}

	var conversationIDs []int
	// Deserialize the []int slice (the old conversationIDs)
	if jsonErr := json.Unmarshal(oldConversationIDsJson, &conversationIDs); jsonErr != nil {
		return jsonErr
	}

	var newConversationIDs []int
	for _, id := range conversationIDs {
		if id != removeFromConversationId {
			newConversationIDs = append(newConversationIDs, id)
		}
	}

	// Serialize the []int slice (the newConversationIDs)
	conversationIDsJson, jsonErr := json.Marshal(newConversationIDs)
	if jsonErr != nil {
		return jsonErr
	}

	// _, err2 := db.c.Exec("UPDATE user_to_conversations SET conversations = ? WHERE username = ?", conversationIDsJson, username)
	_, err2 := db.c.Exec("UPDATE users SET conversations = ? WHERE id = ?", conversationIDsJson, userID)

	return err2
}
