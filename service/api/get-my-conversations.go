package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	//"math/rand"
)

// Start a new game generating the secret number
// and return the created game id
func (rt *_router) getMyConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("Func getMyConversations Called")

	username := ps.ByName("Username")

	user, exists := Users[username]
	if !exists {
		fmt.Println("User ", username, " is not in the database!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//get the conversations of a user
	var myConversations []Conversation
	for _, conversationId := range user.MyConversations {

		myConversations = append(myConversations, AllConversations[conversationId])
	}

	fmt.Println("User ", username, " sucessfully received its conversations!")

	json.NewEncoder(w).Encode(myConversations)

}
