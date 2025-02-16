package database

import (
	"encoding/json"
	"sapienza/wasatext/service/api/util"
)

// InsertConversation updates the conversation table by creating a new conversation entry
func (db *appdbimpl) InsertConversation(newConversation util.Conversation) (int, error) {

	// serialize the []string slice (the participants)
	participantsJson, jsonErr := json.Marshal(newConversation.Participants)
	if jsonErr != nil {
		return 0, jsonErr
	}

	// serialize the []Messages slice (the messages)
	messagesJson, jsonErr2 := json.Marshal([]util.Message{})
	if jsonErr2 != nil {
		return 0, jsonErr2
	}

	// make query to insert the serialized conversation struct to the conversations table
	result, err := db.c.Exec("INSERT INTO conversations (type, group_name, group_picture, participants, messages) VALUES (?, ?, ?, ?, ?)",
		newConversation.Type, newConversation.GroupName, newConversation.GroupPicture, participantsJson, messagesJson)
	if err != nil {
		return 0, err
	}

	// retrieve the last inserted row ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), err
}
