package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	//"math/rand"
)

func (rt *_router) getMyConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("Func getMyConversations Called")

	//username := ps.ByName("Username")

	if UserLoggedIn == nil {
		fmt.Println("User is not logged in!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// user, exists := AllUsers[username]
	// if !exists {
	// 	fmt.Println("User ", username, " is not in the database!")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	//get the conversations of a user
	var myConversations []Conversation
	for _, conversationId := range UserLoggedIn.MyConversations {

		myConversations = append(myConversations, AllConversations[conversationId])
	}

	fmt.Println("User ", UserLoggedIn.Username, " sucessfully received its conversations!")
	json.NewEncoder(w).Encode(myConversations)

}
