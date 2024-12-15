package database

import (
	"encoding/json"
	"sapienza/wasatext/service/api/util"
)

// GetConversation gets the conversation with a given ID from the conversation table
func (db *appdbimpl) GetConversation(id int) (util.Conversation, error) {
	var returnedConversation util.Conversation

	// Intermediate variables to hold JSON data from the database
	var participantsJson []byte
	var messagesJson []byte

	err := db.c.QueryRow("SELECT id, type, group_name, group_picture, participants, messages FROM conversations WHERE id = ?", id).Scan(
		&returnedConversation.Id,
		&returnedConversation.Type,
		&returnedConversation.GroupName,
		&returnedConversation.GroupPicture,
		&participantsJson,
		&messagesJson)

	if err != nil {
		return returnedConversation, err
	}

	// Deserialize the []string slice (the participants)
	if jsonErr := json.Unmarshal(participantsJson, &returnedConversation.Participants); jsonErr != nil {
		return returnedConversation, jsonErr
	}

	// Deserialize the []Messages slice (the messages)
	if jsonErr := json.Unmarshal(messagesJson, &returnedConversation.Messages); jsonErr != nil {
		return returnedConversation, jsonErr
	}

	return returnedConversation, nil
}
