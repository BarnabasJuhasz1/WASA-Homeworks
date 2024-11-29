package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func addToGroup Called-----")

	//make sure user is logged in
	if !isUserLoggedIn(w) {
		return
	}

	// conversationIDString := ps.ByName("ConversationID")

	// conversationID, convErr := strconv.Atoi(conversationIDString)
	// if convErr != nil || conversationID < 0 {
	// 	fmt.Println("Invalid conversationID in path! ", convErr)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// Conversation, existsConv := AllConversations[conversationID]
	// //if the conversation does not exist
	// if !existsConv {
	// 	fmt.Println("Invalid conversationID in path! ", existsConv)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// get the conversation from path
	Conversation, convErr := getConversationFromPath(w, ps)
	if convErr {
		return
	}

	// make sure the logged in user belongs to the conversation
	if !userBelongsToConversation(w, Conversation, *UserLoggedIn) {
		fmt.Println("User is not in the conversation!")
		w.WriteHeader(http.StatusForbidden)
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

	userToAdd, userExists := AllUsers[requestBody.UserNameToAdd]
	if !userExists {
		fmt.Println("User ", requestBody.UserNameToAdd, " is not in the database!")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if the person already belongs to the group or not
	// if isUserFoundInList(Conversation.ConversationGroup.Participants, requestBody.UserNameToAdd) {

	// 	fmt.Println("User ", requestBody.UserNameToAdd, " is already in the group!")
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return

	// }

	// make sure user to add does not already belong to the conversation
	if userBelongsToConversation(w, Conversation, userToAdd) {
		fmt.Println("User is already in the conversation!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// add user to the group (group's perspective)
	Conversation.ConversationGroup.Participants = append(Conversation.ConversationGroup.Participants, userToAdd)

	// add user to the group (user's perspective)
	userToAdd.MyConversations = append(userToAdd.MyConversations, Conversation.Id)

	// update users map by reassigning the struct
	AllUsers[requestBody.UserNameToAdd] = userToAdd
	// update conversations map by reassigning the struct
	AllConversations[Conversation.Id] = Conversation

	fmt.Println("-----Func addToGroup Finished-----")
	json.NewEncoder(w).Encode(Conversation)

}
