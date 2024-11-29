package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func forwardMessage Called-----")

	// make sure user is logged in
	if !isUserLoggedIn(w) {
		return
	}

	// get the conversation from path
	OriginalConversation, convErr := getConversationFromPath(w, ps)
	if convErr {

		return
	}

	// make sure the logged in user belongs to the conversation
	if !userBelongsToConversation(w, OriginalConversation, *UserLoggedIn) {
		fmt.Println("User is not in the conversation!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// get messageID from path
	messageIDString := ps.ByName("MessageID")
	// make sure the messageID is correct
	messageID, messageErr := strconv.Atoi(messageIDString)
	if messageErr != nil || messageID < 0 || messageID >= len(OriginalConversation.Messages) {
		fmt.Println("Invalid MessageID in path! ", messageErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read the request body
	var requestBody struct {
		RecipientUsername string `json:"RecipientUsername"`
	}
	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// make sure the recipient exists in the database as a user
	Recipient, userExists := AllUsers[requestBody.RecipientUsername]
	if !userExists {
		fmt.Println("User ", requestBody.RecipientUsername, " is not in the database!")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// find the one-on-one conversation you have with the recipient
	ConvWithRecipient, exists := getOneOnOneConversationWithUser(*UserLoggedIn, Recipient)
	if !exists {

		fmt.Println("New conversation created between: ", UserLoggedIn.Username, " and ", Recipient.Username)

		// if no such conversation exist, then create a one-on-one conversation with the person (or you cannot forward the message?)
		var emptyMessages []Message
		// create the conversation with the recipient
		ConvWithRecipient = Conversation{
			Id: len(AllConversations),
			ConversationGroup: Group{
				Participants: []User{Recipient},
				GroupName:    "New Conversation",
			},
			Type:     UserType,
			Messages: emptyMessages,
		}

		// add the new conversation to the users
		UserLoggedIn.MyConversations = append(UserLoggedIn.MyConversations, ConvWithRecipient.Id)
		Recipient.MyConversations = append(Recipient.MyConversations, ConvWithRecipient.Id)
	}

	var emptyReactions []Reaction
	// send the message to the recipient by creating a new message with the same content of the original message
	// and modify one-on-one conversation by adding the new message
	ConvWithRecipient.Messages = append(OriginalConversation.Messages, Message{
		Id:              len(ConvWithRecipient.Messages),
		Sender:          *UserLoggedIn,
		Content:         OriginalConversation.Messages[messageID].Content,
		Timestamp:       time.Now().Format("2006-01-02 15:04:05"),
		Status:          UserName,
		EmojiReactions:  emptyReactions,
		OriginMessageId: -1,
	})

	// update conversations map by reassigning the struct
	AllConversations[ConvWithRecipient.Id] = ConvWithRecipient

	fmt.Println("-----Func forwardMessage Finished-----")
	json.NewEncoder(w).Encode(OriginalConversation.Messages[messageID])

}
