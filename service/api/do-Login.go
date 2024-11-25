package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//we should not store info of which conversation a user is part of inside the user struct
//instead, in a separate map that maps from username (string) to conversationIds ([]int)

type User struct {
	//Id int
	Username        string
	ProfilePicture  string
	MyConversations []int
}

// var AllUsers = []User{}
var AllUsers = make(map[string]User)

type Group struct {
	//Id int
	Participants []User
	GroupName    string
	GroupPicture string
}

type MessageStatus string
type ConversationType string
type ReactionType string

const (
	UserType  ConversationType = "UserType"
	GroupType ConversationType = "GroupType"

	SingleCheckmark MessageStatus = "SingleCheckmark"
	DoubleCheckmark MessageStatus = "DoubleCheckmark"
	UserName        MessageStatus = "Username"

	MessageReaction ReactionType = "MessageReaction"
	EmojiReaction   ReactionType = "EmojiReaction"
)

type Message struct {
	Id        int //Id grows incrementally. if a message gets deleted, the content gets deleted, and the message is flagged.
	Sender    User
	Content   string
	Timestamp string
	Status    MessageStatus
	//only stores reactions of type "EmojiReaction", as "MessageReaction"s will just create a message that has originMessageId set
	EmojiReactions []Reaction
	//stores the id of the message this is a reply to. if this is not a reply message, the value is initialized to -1.
	OriginMessageId int
}

type Reaction struct {
	UserWhoReacted User
	Type           ReactionType
	//the content of the reaction. if the reaction is a message, this is a reply and so the content is the message content.
	//if the reaction is an emoji, this is the encoding of the emoji
	Content string
}

type Conversation struct {
	Id                int //Id corresponding to the key in AllConversations
	ConversationGroup Group
	Type              ConversationType
	Messages          []Message
}

var AllConversations = make(map[int]Conversation) //keys grow incrementally

var UserLoggedIn *User

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func doLogin Called-----")

	// Read the request body
	var requestBody struct {
		Username string `json:"Username"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	if len(requestBody.Username) < 3 || len(requestBody.Username) > 16 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// id, err := strconv.Atoi(ps.ByName("Username"))
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	//id := len(Users)

	// Users = append(Users, User{
	// 	//Id: id,
	// 	Username: user,
	// })

	user, exists := AllUsers[requestBody.Username]

	if exists {
		fmt.Println("User ", user, " logged in sucessfully!")

	} else {
		//create a new user
		AllUsers[requestBody.Username] = User{
			Username: requestBody.Username,
		}
		user = AllUsers[requestBody.Username]

		fmt.Println("User ", requestBody.Username, " created sucessfully!")
	}

	UserLoggedIn = &user

	fmt.Println("-----Func doLogin Finished-----")
	json.NewEncoder(w).Encode(requestBody.Username)
}

func isUserLoggedIn(w http.ResponseWriter) bool {

	if UserLoggedIn == nil {
		fmt.Println("User is not logged in!")
		w.WriteHeader(http.StatusForbidden)
		return false
	}
	return true
}

// gets the conversation from the path.
// if there is an error, the second value of the tuple will be set to true
func getConversationFromPath(w http.ResponseWriter, ps httprouter.Params) (Conversation, bool) {

	conversationIDString := ps.ByName("ConversationID")

	//make sure the conversationID is correct
	conversationID, convErr := strconv.Atoi(conversationIDString)
	if convErr != nil || conversationID < 0 { //|| messageID > len() {
		fmt.Println("Invalid conversationID in path! ", convErr)
		w.WriteHeader(http.StatusBadRequest)
		return Conversation{}, true
	}

	//make sure the conversation exists
	ConversationStruct, existsConv := AllConversations[conversationID]
	if !existsConv {
		fmt.Println("Invalid conversationID in path! ", existsConv)
		w.WriteHeader(http.StatusBadRequest)
		return Conversation{}, true
	}

	return ConversationStruct, false
}

// Function to check if a user is found in a list of users
func isUserFoundInList(users []User, userNameToCheck string) bool {
	for _, user := range users {
		if user.Username == userNameToCheck {
			return true
		}
	}
	return false
}

func userBelongsToConversation(w http.ResponseWriter, conv Conversation, user User) bool {
	//check if the person already belongs to the group or not
	if isUserFoundInList(conv.ConversationGroup.Participants, user.Username) {
		return true
	}
	return false
}

// gets the conversation from the path.
// if the conversation is found, the second value of the tuple will be set to true
func getOneOnOneConversationWithUser(userA User, userB User) (Conversation, bool) {

	//loop through all the conversations of userA
	for _, userAconvIDi := range userA.MyConversations {

		ConversationAti, existsConv := AllConversations[userAconvIDi]

		//make sure conversation exists
		if existsConv {

			//if it is a one on one conversation
			if ConversationAti.Type == UserType {

				//if the other user matches with userB
				if ConversationAti.ConversationGroup.Participants[0].Username == userB.Username {
					return ConversationAti, true
				}
			}
		}
	}
	return Conversation{}, false
}
