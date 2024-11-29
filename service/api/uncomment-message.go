package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Handles the request to uncomment a message (remove the emoji reaction you placed on a message)
func (rt *_router) uncommentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func uncommentMessage Called-----")

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

	// loop through all the emoji reactions to this message
	for i, ReactionAti := range Conversation.Messages[messageID].EmojiReactions {

		// if the user has a reaction to this message already, replace that reaction with the new one
		if ReactionAti.UserWhoReacted.Username == UserLoggedIn.Username {

			// remove your emoji reactions from the list of emoji reactions
			Conversation.Messages[messageID].EmojiReactions = append(Conversation.Messages[messageID].EmojiReactions[:i], Conversation.Messages[messageID].EmojiReactions[i+1:]...)
			break

			// if we checked all the reactions and none of them were created by the user
		} else if i == len(Conversation.Messages[messageID].EmojiReactions)-1 {

			fmt.Println("There is no emoji reaction placed by user on the specified message!")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// update conversations map by reassigning the struct
	AllConversations[Conversation.Id] = Conversation

	fmt.Println("-----Func uncommentMessage Finished-----")
	json.NewEncoder(w).Encode(Conversation.Messages[messageID+1])

}
