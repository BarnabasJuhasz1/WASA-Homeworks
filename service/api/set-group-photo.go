package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func setGroupPhoto Called-----")

	//username := ps.ByName("Username")

	//make sure user is logged in
	if !isUserLoggedIn(w) {
		return
	}

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

	if Conversation.Type == UserType {
		fmt.Println("Cannot change photo of one-on-one conversations!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Read the request body
	var requestBody struct {
		GroupPicture string `json:"GroupPicture"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if Conversation.ConversationGroup.GroupPicture == requestBody.GroupPicture {
		fmt.Println("The group picture is already set to this picture.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//change the group picture to the payload picture
	Conversation.ConversationGroup.GroupPicture = requestBody.GroupPicture

	//update conversations map by reassigning the struct
	AllConversations[Conversation.Id] = Conversation

	fmt.Println("-----Func setGroupPhoto Finished-----")
	json.NewEncoder(w).Encode(Conversation)

}
