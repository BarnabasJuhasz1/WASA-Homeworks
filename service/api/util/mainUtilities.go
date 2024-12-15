package util

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sapienza/wasatext/service/api/reqcontext"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// func GetLoggedInUser(w http.ResponseWriter, ctx reqcontext.RequestContext, ADB database.AppDatabase) User {

// 	if !ctx.Authenticated {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return User{}
// 	}

// 	// FoundUser, err := AllUsers[ctx.UserID]
// 	FoundUser, err := ADB.GetUser(ctx.UserID)
// 	if err != nil {
// 		ctx.Logger.Warningln("User with ID: ", ctx.UserID, " requested but not found in database!")
// 		w.WriteHeader(http.StatusNotFound)
// 		return User{}
// 	}

// 	return FoundUser
// }

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

// Function to check if a user is found in a list of usernames
func IsUserFoundInList(usernames []string, userNameToCheck string) bool {
	for _, user := range usernames {
		if user == userNameToCheck {
			return true
		}
	}
	return false
}

func UserBelongsToConversation(conv Conversation, user User) bool {
	// check if the person already belongs to the group or not
	return IsUserFoundInList(conv.Participants, user.Username)

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
				if ConversationAti.Participants[0] == userB.Username {
					return ConversationAti, true
				}
			}
		}
	}
	return Conversation{}, false
}

func GetBasicUserPicture() ([]byte, error) {
	return GetPicture("/default/basic_user_picture.png")
}

func GetBasicGroupPicture() ([]byte, error) {
	return GetPicture("/default/basic_group_picture.png")
}

func GetPicture(imagePath string) ([]byte, error) {

	cwd, _ := os.Getwd()
	filePath := filepath.Join(filepath.Dir(filepath.Dir(cwd)), "/public/images", imagePath)

	// Open the file
	imageFile, err := os.Open(filePath)
	if err != nil {
		return []byte{}, errors.New("failed to open image file: " + err.Error())
	}
	defer imageFile.Close()

	// Read the file contents into a byte slice
	binaryData, err := io.ReadAll(imageFile)
	if err != nil {
		return []byte{}, errors.New("failed to read image file: " + err.Error())
	}

	return binaryData, nil
}
