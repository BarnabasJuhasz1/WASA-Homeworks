package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	//"math/rand"
)

func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("Func getConversation Called")

	conversationIDString := ps.ByName("ConversationID")

	//fmt.Println("User ", username, " sucessfully received its conversations!")
	conversationID, err := strconv.Atoi(conversationIDString)
	if err != nil || conversationID < 0 {
		fmt.Println("Invalid conversationID in path! ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//make sure the user asks for a conversation it is part of
	if !contains(UserLoggedIn.MyConversations, conversationID) {
		fmt.Println("User tried to read a conversation it is not part of! ", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(AllConversations[conversationID])

}

// function to check if an int ID is present in a list of int IDs
func contains(userConversationIds []int, conversationId int) bool {
	for _, item := range userConversationIds {
		if item == conversationId {
			return true
		}
	}
	return false
}
