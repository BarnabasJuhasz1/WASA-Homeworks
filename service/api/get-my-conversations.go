package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func getMyConversations Called-----")

	// username := ps.ByName("Username")

	// make sure user is logged in
	if !isUserLoggedIn(w) {
		return
	}

	// user, exists := AllUsers[username]
	// if !exists {
	// 	fmt.Println("User ", username, " is not in the database!")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// get the conversations of a user
	var myConversations []Conversation
	for _, conversationId := range UserLoggedIn.MyConversations {

		myConversations = append(myConversations, AllConversations[conversationId])
	}

	fmt.Println("User ", UserLoggedIn.Username, " sucessfully received its conversations!")

	fmt.Println("-----Func getMyConversations Finished-----")
	json.NewEncoder(w).Encode(myConversations)

}
