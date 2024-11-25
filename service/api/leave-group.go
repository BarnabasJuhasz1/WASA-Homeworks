package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) leaveGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func leaveGroup Called-----")

	//make sure user is logged in
	if !isUserLoggedIn(w) {
		return
	}

	// conversationIDString := ps.ByName("ConversationID")

	// //Check conversationID in path
	// conversationID, err := strconv.Atoi(conversationIDString)
	// if err != nil || conversationID < 0 {
	// 	fmt.Println("Invalid conversationID in path! ", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// Conversation, existsConv := AllConversations[conversationID]
	// if !existsConv {
	// 	fmt.Println("Invalid conversationID in path! ", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	//get the conversation from path
	Conversation, convErr := getConversationFromPath(w, ps)
	if convErr {
		return
	}

	//make sure the logged in user belongs to the conversation
	if !userBelongsToConversation(w, Conversation, *UserLoggedIn) {
		fmt.Println("User is not in the conversation!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//delete user from group (group's perspective)
	Conversation.ConversationGroup.Participants = deleteUserFromList(Conversation.ConversationGroup.Participants, *UserLoggedIn)

	//delete user from group (users's perspective)
	//Remark: since "UserLoggedIn" holds a pointer, the "Users" map is also updated
	UserLoggedIn.MyConversations = deleteConversationIdFromList(UserLoggedIn.MyConversations, Conversation.Id)

	//update the allConversations map by reassigning the struct
	AllConversations[Conversation.Id] = Conversation

	fmt.Println("-----Func leaveGroup Finished-----")
	json.NewEncoder(w).Encode(UserLoggedIn.MyConversations)
}

// Function to delete a user from a list of users
func deleteUserFromList(users []User, userToDelete User) []User {

	var updatedUsers []User

	for _, user := range users {
		//we only add the user to the new list if it has a different name then the user to delete
		if !(user.Username == userToDelete.Username) {
			updatedUsers = append(updatedUsers, user)
		}
	}

	return updatedUsers
}

func deleteConversationIdFromList(conversationsOfUser []int, conversationToLeave int) []int {
	var updatedConversations []int

	for _, conversationId := range conversationsOfUser {
		//we only add the user to the new list if it has a different name then the user to delete
		if conversationId != conversationToLeave {
			updatedConversations = append(updatedConversations, conversationId)
		}
	}

	return updatedConversations
}
