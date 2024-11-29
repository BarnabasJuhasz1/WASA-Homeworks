package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func sendMessage Called-----")

	// make sure user is logged in
	if !isUserLoggedIn(w) {
		return
	}

	conversationIDString := ps.ByName("ConversationID")

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

	// make sure the user asks for a conversation it is part of
	if !contains(UserLoggedIn.MyConversations, conversationID) {
		fmt.Println("User tried to send a message to a conversation it is not part of! ", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Read the request body
	var requestBody struct {
		// SenderUsername string `json:"SenderUsername"`
		MessageContent string `json:"MessageContent"`
	}

	err2 := json.NewDecoder(r.Body).Decode(&requestBody)
	if err2 != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	if len(requestBody.MessageContent) == 0 || len(requestBody.MessageContent) > 10000 {
		fmt.Println("Message too long or too short!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SenderUser, existsUser := Users[requestBody.SenderUsername]
	// if !existsUser {
	// 	fmt.Println("User ", requestBody.SenderUsername, " is not in the database!")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	var emptyReactions []Reaction

	// modify conversation by adding the new message
	Conversation.Messages = append(Conversation.Messages, Message{
		Id: len(Conversation.Messages),
		// Sender:    SenderUser,
		Sender:          *UserLoggedIn,
		Content:         requestBody.MessageContent,
		Timestamp:       time.Now().Format("2006-01-02 15:04:05"),
		Status:          UserName,
		EmojiReactions:  emptyReactions,
		OriginMessageId: -1,
	})
	// update the allConversations map by reassigning the struct
	AllConversations[conversationID] = Conversation

	fmt.Println("-----Func sendMessage Finished-----")
	json.NewEncoder(w).Encode(AllConversations[conversationID])

}
