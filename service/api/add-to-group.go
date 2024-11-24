package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	//"math/rand"
)

// Start a new game generating the secret number
// and return the created game id
func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("Func addToGroup Called")

	conversationIDString := ps.ByName("ConversationID")

	conversationID, convErr := strconv.Atoi(conversationIDString)
	if convErr != nil || conversationID < 0 {
		fmt.Println("Invalid conversationID in path! ", convErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Conversation, existsConv := AllConversations[conversationID]
	//if the conversation does not exist
	if !existsConv {
		fmt.Println("Invalid conversationID in path! ", existsConv)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read the request body
	var requestBody struct {
		UserNameToAdd string `json:"UserNameToAdd"`
	}

	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	userToAdd, userExists := Users[requestBody.UserNameToAdd]
	if !userExists {
		fmt.Println("User ", requestBody.UserNameToAdd, " is not in the database!")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//check if the person already belongs to the group or not
	if isUserFound(Conversation.ConversationGroup.Participants, requestBody.UserNameToAdd) {

		fmt.Println("User ", requestBody.UserNameToAdd, " is already in the group!")
		w.WriteHeader(http.StatusForbidden)
		return

	}

	//add user to the group (group's perspective)
	Conversation.ConversationGroup.Participants = append(Conversation.ConversationGroup.Participants, userToAdd)

	//add user to the group (user's perspective)
	userToAdd.MyConversations = append(userToAdd.MyConversations, Conversation.Id)

	//update users map by reassigning the struct
	Users[requestBody.UserNameToAdd] = userToAdd
	//update conversations map by reassigning the struct
	AllConversations[conversationID] = Conversation

	json.NewEncoder(w).Encode(Conversation)

}

// Function to check if a user is found in a list of users
func isUserFound(users []User, userNameToCheck string) bool {
	for _, user := range users {
		if user.Username == userNameToCheck {
			return true
		}
	}
	return false
}
