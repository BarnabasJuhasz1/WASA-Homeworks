package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func deleteMessage Called-----")

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

	//get messageID from path
	messageIDString := ps.ByName("MessageID")
	//make sure the messageID is correct
	messageID, messageErr := strconv.Atoi(messageIDString)
	if messageErr != nil || messageID < 0 || messageID > len(Conversation.Messages) {
		fmt.Println("Invalid MessageID in path! ", messageErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//make sure you can only delete messages sent by you
	if Conversation.Messages[messageID].Sender.Username != UserLoggedIn.Username {
		fmt.Println("Tried to delete a message not sent by you!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//delete message content
	Conversation.Messages[messageID].Content = "Message content was deleted."
	//update conversations map by reassigning the struct
	AllConversations[Conversation.Id] = Conversation

	fmt.Println("-----Func deleteMessage Finished-----")
	json.NewEncoder(w).Encode(Conversation)

}
