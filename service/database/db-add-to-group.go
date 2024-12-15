package database

import "encoding/json"

// AddToGroup updates the user_to_group table which maps users to conversation IDs
func (db *appdbimpl) AddToGroup(username string, addToConversationId int) error {

	var oldConversationIDsJson []byte

	// get the current conversation IDs
	err := db.c.QueryRow("SELECT conversations FROM user_to_conversations WHERE username = ?", username).Scan(&oldConversationIDsJson)
	if err != nil {
		return err
	}

	var conversationIDs []int
	// Deserialize the []int slice (the old conversationIDs)
	if jsonErr := json.Unmarshal(oldConversationIDsJson, &conversationIDs); jsonErr != nil {
		return jsonErr
	}

	conversationIDs = append(conversationIDs, addToConversationId)

	// Serialize the []int slice (the conversationIDs)
	conversationIDsJson, jsonErr := json.Marshal(conversationIDs)
	if jsonErr != nil {
		return jsonErr
	}

	_, err2 := db.c.Exec("UPDATE user_to_conversations SET conversations = ? WHERE username = ?", conversationIDsJson, username)

	return err2
}
