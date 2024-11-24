package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	//"math/rand"
)

// Start a new game generating the secret number
// and return the created game id
func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("Func getConversation Called")

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

	json.NewEncoder(w).Encode(AllConversations[conversationID])

}
