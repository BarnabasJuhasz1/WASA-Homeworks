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

	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

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
	// var ConversationUsers []util.User
	// // add the user logged in to the conversation automatically
	// ConversationUsers = append(ConversationUsers, LoggedInUser)

	// for _, userNameAtI := range requestBody.Participants {

	// 	// check if all the users exist, if so, add them to a list
	// 	userToAdd, userExistsError := rt.db.GetUser(userNameAtI)
	// 	if userExistsError != nil {
	// 		ctx.Logger.Debugln("User", userNameAtI, "is not in the database!")

	// 		w.WriteHeader(http.StatusNotFound)
	// 		return
	// 	}
	// 	ConversationUsers = append(ConversationUsers, userToAdd)
	// }

	var ConversationUsers []string
	// add the user logged in to the conversation automatically
	ConversationUsers = append(ConversationUsers, LoggedInUser.Username)

	for _, userNameAtI := range requestBody.Participants {

		// check if all the users exist, if so, add them to a list
		userToAdd, userExistsError := rt.db.GetUser(userNameAtI)
		if userExistsError != nil {
			ctx.Logger.Debugln("User", userNameAtI, "is not in the database!")

			w.WriteHeader(http.StatusNotFound)
			return
		}
		ConversationUsers = append(ConversationUsers, userToAdd.Username)
	}

	var emptyMessages []util.Message
	basicPic, _ := util.GetBasicGroupPicture()
	// create the conversation with all the users
	conversation := util.Conversation{
		// Id: len(util.AllConversations),
		// ConversationGroup: util.Group{
		// 	Participants: ConversationUsers,
		// 	GroupName:    "New Group",
		// 	GroupPicture: basicPic,
		// },

		Type:         requestBody.ConvType,
		GroupName:    "New Group",
		GroupPicture: basicPic,
		Participants: ConversationUsers,
		Messages:     emptyMessages,
	}

	// add all users to the new conversation (users' perspective)
	// for _, userToAddAtI := range ConversationUsers {

	// 	// add user to the group (users' perspective)
	// 	userToAddAtI.MyConversations = append(userToAddAtI.MyConversations, conversation.Id)

	// 	// update users map by reassigning the struct
	// 	// util.AllUsers[userToAddAtI.Username] = userToAddAtI

	// }
	// add new conversation to conversations map

	id, insertingErr := rt.db.InsertConversation(conversation)
	if insertingErr != nil {
		ctx.Logger.Errorln("Error inserting new conversation into database! ", insertingErr)
	}

	conversation.Id = id
	util.AllConversations[conversation.Id] = conversation

	ctx.Logger.Debugln("-----Func createConversation Finished-----")
	encodeErr := json.NewEncoder(w).Encode(conversation)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
