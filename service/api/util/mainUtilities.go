package util

import (
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetLoggedInUser(w http.ResponseWriter, ctx reqcontext.RequestContext) User {

	if !ctx.Authenticated {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return User{}
	}

	FoundUser, err := AllUsers[ctx.UserID]
	if err {
		ctx.Logger.Warningln("User with ID: ", ctx.UserID, " requested but not found in database!")
		w.WriteHeader(http.StatusNotFound)
		return User{}
	}

	return FoundUser
}

// gets the conversation from the path.
// if there is an error, the second value of the tuple will be set to true
func GetConversationFromPath(w http.ResponseWriter, ps httprouter.Params, ctx reqcontext.RequestContext) (Conversation, bool) {

	conversationIDString := ps.ByName("ConversationID")

	// make sure the conversationID is correct
	conversationID, convErr := strconv.Atoi(conversationIDString)
	if convErr != nil || conversationID < 0 { //|| messageID > len() {
		ctx.Logger.Debugln("Invalid conversationID in path! ", convErr)

		w.WriteHeader(http.StatusBadRequest)
		return Conversation{}, true
	}

	// make sure the conversation exists
	ConversationStruct, existsConv := AllConversations[conversationID]
	if !existsConv {
		ctx.Logger.Debugln("Invalid conversationID in path! ", existsConv)

		w.WriteHeader(http.StatusBadRequest)
		return Conversation{}, true
	}

	return ConversationStruct, false
}

// Function to check if a user is found in a list of users
func IsUserFoundInList(users []User, userNameToCheck string) bool {
	for _, user := range users {
		if user.Username == userNameToCheck {
			return true
		}
	}
	return false
}

func UserBelongsToConversation(conv Conversation, user User) bool {
	// check if the person already belongs to the group or not
	return IsUserFoundInList(conv.ConversationGroup.Participants, user.Username)

}

// gets the conversation from the path.
// if the conversation is found, the second value of the tuple will be set to true
func GetOneOnOneConversationWithUser(userA User, userB User) (Conversation, bool) {

	// loop through all the conversations of userA
	for _, userAconvIDi := range userA.MyConversations {

		ConversationAti, existsConv := AllConversations[userAconvIDi]

		// make sure conversation exists
		if existsConv {

			// if it is a one on one conversation
			if ConversationAti.Type == UserType {

				// if the other user matches with userB
				if ConversationAti.ConversationGroup.Participants[0].Username == userB.Username {
					return ConversationAti, true
				}
			}
		}
	}
	return Conversation{}, false
}
