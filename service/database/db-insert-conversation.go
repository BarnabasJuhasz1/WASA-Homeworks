package database

import (
	"encoding/json"
	"sapienza/wasatext/service/api/util"
)

// InsertConversation updates the conversation table by creating a new conversation entry
func (db *appdbimpl) InsertConversation(newConversation util.Conversation) (int, error) {

	// Serialize the []string slice (the participants)
	participantsJson, jsonErr := json.Marshal(newConversation.Participants)
	if jsonErr != nil {
		return 0, jsonErr
	}

	// Serialize the []Messages slice (the messages)
	messagesJson, jsonErr2 := json.Marshal(newConversation.Messages)
	if jsonErr2 != nil {
		return 0, jsonErr2
	}

	// _, err := db.c.Exec("INSERT INTO conversations (id, type, group_name, group_picture, participants, messages) VALUES (?, ?, ?, ?, ?, ?) ON CONFLICT(id) DO NOTHING",
	result, err := db.c.Exec("INSERT INTO conversations (type, group_name, group_picture, participants, messages) VALUES (?, ?, ?, ?, ?)",
		newConversation.Type, newConversation.GroupName, newConversation.GroupPicture, participantsJson, messagesJson)
	if err != nil {
		return 0, err
	}

	// Retrieve the last inserted row ID
	lastInsertID, err := result.LastInsertId() // This method returns the last insert ID.
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), err
}

// type UserSlice []util.User

// // Implement driver.Valuer to serialize into JSON
// func (u UserSlice) Value() (driver.Value, error) {
// 	return json.Marshal(u)
// }

// // Implement sql.Scanner to deserialize from JSON
// func (u *UserSlice) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return fmt.Errorf("type assertion to []byte failed")
// 	}
// 	return json.Unmarshal(bytes, u)
// }
