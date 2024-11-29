package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func setGroupName Called-----")

	// make sure user is logged in
	if !isUserLoggedIn(w) {
		return
	}

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

	if Conversation.Type == UserType {
		fmt.Println("Cannot change name of one-on-one conversations!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Read the request body
	var requestBody struct {
		GroupName string `json:"GroupName"`
	}

	err2 := json.NewDecoder(r.Body).Decode(&requestBody)
	if err2 != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if len(requestBody.GroupName) < 3 || len(requestBody.GroupName) > 16 {
		fmt.Println("Group name too short or too long ", requestBody.GroupName)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if Conversation.ConversationGroup.GroupName == requestBody.GroupName {
		fmt.Println("New Group name matches old one! ", requestBody.GroupName)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// modify conversation by changing the group name
	Conversation.ConversationGroup.GroupName = requestBody.GroupName
	// update the allConversations map by reassigning the struct
	AllConversations[Conversation.Id] = Conversation

	fmt.Println("-----Func setGroupName Finished-----")
	json.NewEncoder(w).Encode(Conversation)

}
