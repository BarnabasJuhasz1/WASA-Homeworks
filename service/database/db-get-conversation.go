package database

import (
	"encoding/json"
	"sapienza/wasatext/service/api/util"
)

// GetConversation gets the conversation with a given ID from the conversation table
func (db *appdbimpl) GetConversation(id int) (util.Conversation, error) {

	// initialize conversation struct to be returned
	var returnedConversation util.Conversation

	// intermediate variables to hold JSON data from the database
	var participantsJson []byte
	var messagesJson []byte

	// query the row with the given ID
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

	// deserialize the []string slice (the participants)
	if jsonErr := json.Unmarshal(participantsJson, &returnedConversation.Participants); jsonErr != nil {
		return returnedConversation, jsonErr
	}

	// deserialize the []Messages slice (the messages)
	if jsonErr := json.Unmarshal(messagesJson, &returnedConversation.Messages); jsonErr != nil {
		return returnedConversation, jsonErr
	}

	return returnedConversation, nil
}
