package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) leaveGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func leaveGroup Called-----")

	// conversationIDString := ps.ByName("ConversationID")

	// //Check conversationID in path
	// conversationID, err := strconv.Atoi(conversationIDString)
	// if err != nil || conversationID < 0 {
	// 	fmt.Println("Invalid conversationID in path! ", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// Conversation, existsConv := AllConversations[conversationID]
	// if !existsConv {
	// 	fmt.Println("Invalid conversationID in path! ", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// get the conversation from path
	Conversation, convErr := util.GetConversationFromPath(w, ps, ctx)
	if convErr {
		return
	}

	// make sure the logged in user belongs to the conversation
	if !util.UserBelongsToConversation(Conversation, util.GetLoggedInUser(w, ctx)) {
		ctx.Logger.Debugln("User is not in the conversation!")

		w.WriteHeader(http.StatusForbidden)
		return
	}

	// delete user from group (group's perspective)
	Conversation.ConversationGroup.Participants = deleteUserFromList(Conversation.ConversationGroup.Participants, util.GetLoggedInUser(w, ctx))

	// delete user from group (users's perspective)
	UserLoggedIn := util.GetLoggedInUser(w, ctx)
	UserLoggedIn.MyConversations = deleteConversationIdFromList(UserLoggedIn.MyConversations, Conversation.Id)
	util.AllUsers[util.GetLoggedInUser(w, ctx).Username] = UserLoggedIn

	// update the allConversations map by reassigning the struct
	util.AllConversations[Conversation.Id] = Conversation

	ctx.Logger.Debugln("-----Func leaveGroup Finished-----")

	encodeErr := json.NewEncoder(w).Encode(util.GetLoggedInUser(w, ctx).MyConversations)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Function to delete a user from a list of users
func deleteUserFromList(users []util.User, userToDelete util.User) []util.User {

	var updatedUsers []util.User

	for _, user := range users {
		// we only add the user to the new list if it has a different name then the user to delete
		if !(user.Username == userToDelete.Username) {
			updatedUsers = append(updatedUsers, user)
		}
	}

	return updatedUsers
}

func deleteConversationIdFromList(conversationsOfUser []int, conversationToLeave int) []int {
	var updatedConversations []int

	for _, conversationId := range conversationsOfUser {
		// we only add the user to the new list if it has a different name then the user to delete
		if conversationId != conversationToLeave {
			updatedConversations = append(updatedConversations, conversationId)
		}
	}

	return updatedConversations
}
