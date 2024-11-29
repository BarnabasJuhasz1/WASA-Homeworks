package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func createConversation Called-----")

	// Read the request body
	var requestBody struct {
		ConvType     util.ConversationType `json:"ConversationType"`
		Participants []string              `json:"Participants"`
	}

	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// check all the userNames provided are in the database
	var ConversationUsers []util.User
	for _, userNameAtI := range requestBody.Participants {

		// check if all the users exist, if so, add them to a list
		userToAdd, userExists := util.AllUsers[userNameAtI]
		if !userExists {
			ctx.Logger.Debugln("User ", userNameAtI, " is not in the database!")

			w.WriteHeader(http.StatusNotFound)
			return
		}
		ConversationUsers = append(ConversationUsers, userToAdd)
	}

	var emptyMessages []util.Message
	// create the conversation with all the users
	conversation := util.Conversation{
		Id: len(util.AllConversations),
		ConversationGroup: util.Group{
			Participants: ConversationUsers,
			GroupName:    "New Group",
		},
		Type:     requestBody.ConvType,
		Messages: emptyMessages,
	}

	// add all users to the new conversation (users' perspective)
	for _, userToAddAtI := range ConversationUsers {

		// add user to the group (users' perspective)
		userToAddAtI.MyConversations = append(userToAddAtI.MyConversations, conversation.Id)

		// update users map by reassigning the struct
		util.AllUsers[userToAddAtI.Username] = userToAddAtI
	}
	// add new conversation to conversations map
	util.AllConversations[conversation.Id] = conversation

	ctx.Logger.Debugln("-----Func createConversation Finished-----")
	json.NewEncoder(w).Encode(conversation)

}
