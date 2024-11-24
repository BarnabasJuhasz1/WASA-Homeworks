package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	//"math/rand"
)

type User struct {
	//Id int
	Username        string
	ProfilePicture  string
	MyConversations []int
}

// var Users = []User{}
var Users = make(map[string]User)

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
	Id        int //Id grows incrementally
	Sender    User
	Content   string
	Timestamp string
	Status    MessageStatus
	Reactions []Reaction
}

type Reaction struct {
	UserWhoReacted User
	Type           ReactionType
}

type Conversation struct {
	Id                int //Id corresponding to the key in AllConversations
	ConversationGroup Group
	Type              ConversationType
	Messages          []Message
}

var AllConversations = make(map[int]Conversation) //keys grow incrementally

var UserLoggedIn *User

// Start a new game generating the secret number
// and return the created game id
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("Func doLogin Called")

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

	user, exists := Users[requestBody.Username]

	if exists {
		fmt.Println("User ", user, " logged in sucessfully!")

	} else {
		//create a new user
		Users[requestBody.Username] = User{
			Username: requestBody.Username,
		}
		user = Users[requestBody.Username]

		fmt.Println("User ", requestBody.Username, " created sucessfully!")
	}

	UserLoggedIn = &user

	json.NewEncoder(w).Encode(requestBody.Username)

}
