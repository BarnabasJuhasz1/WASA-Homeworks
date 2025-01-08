package database

import (
	"encoding/json"
	"sapienza/wasatext/service/api/util"
)

// GetMyConversations gets all the conversation of a given user
func (db *appdbimpl) GetMyConversations(userID int) ([]util.Conversation, error) {

	var returnedMyConversations []util.Conversation

	// get the conversation IDs the user belongs to
	var myConversationIDsJson []byte

	// err := db.c.QueryRow("SELECT conversations FROM user_to_conversations WHERE username = ?", username).Scan(&myConversationIDsJson)
	err := db.c.QueryRow("SELECT conversations FROM users WHERE id = ?", userID).Scan(&myConversationIDsJson)
	if err != nil {
		return returnedMyConversations, err
	}

	var myConversationIDs []int
	// Deserialize the []int slice (the myConversationIDs)
	if jsonErr := json.Unmarshal(myConversationIDsJson, &myConversationIDs); jsonErr != nil {
		return returnedMyConversations, jsonErr
	}

	for _, conversationID := range myConversationIDs {

		conversation, err := db.GetConversation(conversationID)
		if err != nil {

			return returnedMyConversations, err

		} else {

			returnedMyConversations = append(returnedMyConversations, conversation)
		}
	}

	return returnedMyConversations, nil
}
