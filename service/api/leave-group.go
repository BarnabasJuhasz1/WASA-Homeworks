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

	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

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
	Conversation, convErr := GetConversationFromPath(rt, w, ps, ctx)
	if convErr {
		return
	}

	// make sure the logged in user belongs to the conversation
	if !util.UserBelongsToConversation(Conversation, LoggedInUser) {
		ctx.Logger.Debugln("User is not in the conversation!")

		w.WriteHeader(http.StatusForbidden)
		return
	}

	// delete user from group (group's perspective)
	// Conversation.Participants = deleteUserFromList(Conversation.Participants, LoggedInUser)
	Conversation.Participants = deleteIDFromList(Conversation.Participants, LoggedInUser.Id)

	// delete user from group (users's perspective)
	// LoggedInUser.MyConversations = deleteConversationIdFromList(LoggedInUser.MyConversations, Conversation.Id)
	// util.AllUsers[LoggedInUser.Username] = LoggedInUser

	// update the allConversations map by reassigning the struct
	//util.AllConversations[Conversation.Id] = Conversation
	rt.db.UpdateConversation(Conversation.Id, Conversation)

	ctx.Logger.Debugln("-----Func leaveGroup Finished-----")

	// encodeErr := json.NewEncoder(w).Encode(LoggedInUser.MyConversations)
	encodeErr := json.NewEncoder(w).Encode(LoggedInUser)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Function to delete a user from a list of users
// func deleteUserFromList(users []util.User, userToDelete util.User) []util.User {

// 	var updatedUsers []util.User

// 	for _, user := range users {
// 		// we only add the user to the new list if it has a different name then the user to delete
// 		if !(user.Username == userToDelete.Username) {
// 			updatedUsers = append(updatedUsers, user)
// 		}
// 	}

// 	return updatedUsers
// }

func deleteIDFromList(userIDs []int, userIDToDelete int) []int {
	result := []int{}
	for _, userID := range userIDs {
		if userID != userIDToDelete {
			result = append(result, userID)
		}
	}
	return result
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
