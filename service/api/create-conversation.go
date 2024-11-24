package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	//"math/rand"
)

// Start a new game generating the secret number
// and return the created game id
func (rt *_router) createConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("Func createConversation Called")

	// Read the request body
	var requestBody struct {
		ConvType     ConversationType `json:"ConversationType"`
		Participants []string         `json:"Participants"`
	}

	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	//check all the userNames provided are in the database
	var ConversationUsers []User
	for _, userNameAtI := range requestBody.Participants {

		//check if all the users exist, if so, add them to a list
		userToAdd, userExists := Users[userNameAtI]
		if !userExists {
			fmt.Println("User ", userNameAtI, " is not in the database!")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ConversationUsers = append(ConversationUsers, userToAdd)
	}

	var emptyMessages []Message
	//create the conversation with all the users
	conversation := Conversation{
		Id: len(AllConversations),
		ConversationGroup: Group{
			Participants: ConversationUsers,
			GroupName:    "New Group",
		},
		Type:     requestBody.ConvType,
		Messages: emptyMessages,
	}

	//add all users to the new conversation (users' perspective)
	for _, userToAddAtI := range ConversationUsers {

		//add user to the group (users' perspective)
		userToAddAtI.MyConversations = append(userToAddAtI.MyConversations, conversation.Id)

		//update users map by reassigning the struct
		Users[userToAddAtI.Username] = userToAddAtI
	}
	//add new conversation to conversations map
	AllConversations[conversation.Id] = conversation

	json.NewEncoder(w).Encode(conversation)

}
