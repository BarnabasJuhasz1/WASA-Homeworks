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

	// make query to get all conversation IDs corresponding to given userID
	err := db.c.QueryRow("SELECT conversations FROM users WHERE id = ?", userID).Scan(&myConversationIDsJson)
	if err != nil {
		return returnedMyConversations, err
	}

	var myConversationIDs []int
	// deserialize the []int slice (the myConversationIDs)
	if jsonErr := json.Unmarshal(myConversationIDsJson, &myConversationIDs); jsonErr != nil {
		return returnedMyConversations, jsonErr
	}

	// loop through all the conversation IDs retrieved
	for _, conversationID := range myConversationIDs {

		// make query to retrieve conversation corresponding to ID
		conversation, err := db.GetConversation(conversationID)
		if err != nil {
			return returnedMyConversations, err
		} else {
			// append conversation to the list
			returnedMyConversations = append(returnedMyConversations, conversation)
		}
	}

	return returnedMyConversations, nil
}
