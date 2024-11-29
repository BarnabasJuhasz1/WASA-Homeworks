package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func commentMessage Called-----")

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

	// get messageID from path
	messageIDString := ps.ByName("MessageID")
	// make sure the messageID is correct
	messageID, messageErr := strconv.Atoi(messageIDString)
	if messageErr != nil || messageID < 0 || messageID >= len(Conversation.Messages) {
		fmt.Println("Invalid MessageID in path! ", messageErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read the request body
	var requestBody struct {
		TypeOfReaction    ReactionType `json:"ReactionType"`
		ContentOfReaction string       `json:"Content"`
	}
	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// create the instance of the reaction struct
	ReactionToMake := Reaction{
		UserWhoReacted: *UserLoggedIn,
		Type:           requestBody.TypeOfReaction,
		Content:        requestBody.ContentOfReaction,
	}

	// if the reaction is an emoji reaction, just append it to the list of reactions on the message
	if ReactionToMake.Type == EmojiReaction {

		// loop through all the emoji reactions to this message
		for i, ReactionAti := range Conversation.Messages[messageID].EmojiReactions {

			// if the user has a reaction to this message already, replace that reaction with the new one
			if ReactionAti.UserWhoReacted.Username == UserLoggedIn.Username {

				// replace old reaction content with new one
				ReactionAti.Content = ReactionToMake.Content
				// update reaction in the message
				Conversation.Messages[messageID].EmojiReactions[i] = ReactionAti
				break

				// if we checked all the reactions and none of them were created by the user
			} else if i == len(Conversation.Messages[messageID].EmojiReactions)-1 {

				Conversation.Messages[messageID].EmojiReactions = append(Conversation.Messages[messageID].EmojiReactions, ReactionToMake)
			}
		}

		// if the reaction is a reply (=comment), then create a new message
	} else if ReactionToMake.Type == MessageReaction {

		var emptyReactions []Reaction
		// create a new message with the content of the reaction,
		// and add it to the conversation
		Conversation.Messages = append(Conversation.Messages, Message{
			Id:              len(Conversation.Messages),
			Sender:          *UserLoggedIn,
			Content:         ReactionToMake.Content,
			Timestamp:       time.Now().Format("2006-01-02 15:04:05"),
			Status:          UserName,
			EmojiReactions:  emptyReactions,
			OriginMessageId: messageID,
		})
	}

	// update conversations map by reassigning the struct
	AllConversations[Conversation.Id] = Conversation

	fmt.Println("-----Func commentMessage Finished-----")
	json.NewEncoder(w).Encode(Conversation.Messages[messageID+1])

}
