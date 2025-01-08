package util

// we should not store info of which conversation a user is part of inside the user struct
// instead, in a separate map that maps from username (string) to conversationIds ([]int)

type SessionStruct struct {
	User         User
	SessionToken string
}

type User struct {
	Id             int
	Username       string
	ProfilePicture []byte
	// MyConversations []int
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
	Id        int // Id grows incrementally. if a message gets deleted, the content gets deleted, and the message is flagged.
	Sender    int
	Content   string
	Timestamp string
	Status    MessageStatus
	// only stores reactions of type "EmojiReaction", as "MessageReaction"s will just create a message that has originMessageId set
	EmojiReactions []Reaction
	// stores the id of the message this is a reply to. if this is not a reply message, the value is initialized to -1.
	OriginMessageId int
}

type Reaction struct {
	UserWhoReacted string
	Type           ReactionType
	// the content of the reaction. if the reaction is a message, this is a reply and so the content is the message content.
	// if the reaction is an emoji, this is the encoding of the emoji
	Content string
}

type Conversation struct {
	Id int // Id corresponding to the key in AllConversations
	// ConversationGroup Group
	Type         ConversationType
	GroupName    string
	GroupPicture []byte
	// Participants []User
	Participants []int
	Messages     []Message
}

// type Group struct {
// 	//Id int
// 	Participants []User
// 	GroupName    string
// 	GroupPicture []byte
// }

// stores all the conversations of all users. This should only exist in the database...
// var AllConversations = make(map[int]Conversation) // keys grow incrementally

// stores the IDs of the conversations the user is part of
// the field maps from username to list of conversation IDs
// the conversation can be retrieved from the AllConversations field
var UserToConversationIDs = make(map[string]int)
