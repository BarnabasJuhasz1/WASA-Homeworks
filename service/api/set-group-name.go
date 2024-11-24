package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	//"math/rand"
)

func (rt *_router) setGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("Func setGroupName Called")

	conversationIDString := ps.ByName("ConversationID")

	// user, exists := Users[username]
	// if !exists {
	// 	fmt.Println("User ", username, " is not in the database!")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	//get the conversation with a specific ID
	// var myConversations []Conversation
	// for _, conversationId := range user.MyConversations {

	// 	myConversations = append(myConversations, AllConversations[conversationId])
	// }

	//fmt.Println("User ", username, " sucessfully received its conversations!")
	conversationID, err := strconv.Atoi(conversationIDString)
	if err != nil || conversationID < 0 {
		fmt.Println("Invalid conversationID in path! ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Conversation, existsConv := AllConversations[conversationID]
	if !existsConv {
		fmt.Println("Invalid conversationID in path! ", err)
		w.WriteHeader(http.StatusBadRequest)
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

	//modify conversation by changing the group name
	Conversation.ConversationGroup.GroupName = requestBody.GroupName
	//update the allConversations map by reassigning the struct
	AllConversations[conversationID] = Conversation

	json.NewEncoder(w).Encode(AllConversations[conversationID])

}
