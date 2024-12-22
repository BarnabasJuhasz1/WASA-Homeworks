package database

import (
	"encoding/json"
	"sapienza/wasatext/service/api/util"
)

// UpdateUser updates the user table entry if the existing user successfully changed its username or icon
func (db *appdbimpl) UpdateConversation(id int, conversation util.Conversation) error {

	// Serialize the []string slice (the participants)
	participantsJson, jsonErr := json.Marshal(conversation.Participants)
	if jsonErr != nil {
		return jsonErr
	}

	// Serialize the []Messages slice (the messages)
	messagesJson, jsonErr2 := json.Marshal(conversation.Messages)
	if jsonErr2 != nil {
		return jsonErr2
	}

	_, err2 := db.c.Exec("UPDATE conversations SET group_name = ?, group_picture = ?, participants = ?, Messages = ? WHERE id = ?",
		conversation.GroupName,
		conversation.GroupPicture,
		participantsJson,
		messagesJson,
		conversation.Id)

	return err2
}
