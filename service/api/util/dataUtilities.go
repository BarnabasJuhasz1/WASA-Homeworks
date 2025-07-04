package util

type SessionStruct struct {
	User         User
	SessionToken string
}

type User struct {
	Id             int
	Username       string
	ProfilePicture []byte
}

type MessageStatus string
type ConversationType string
type ReactionType string

const (
	UserType  ConversationType = "UserType"
	GroupType ConversationType = "GroupType"

	SingleCheckmark MessageStatus = "SingleCheckmark"
	DoubleCheckmark MessageStatus = "DoubleCheckmark"

	MessageReaction ReactionType = "MessageReaction"
	EmojiReaction   ReactionType = "EmojiReaction"
)

type Message struct {
	Id        int // Id grows incrementally. if a message gets deleted, the content gets deleted, and the message is flagged with "HasBeenDeleted" = true.
	Sender    int
	Content   string
	Timestamp string
	Status    MessageStatus
	// only stores reactions of type "EmojiReaction", as "MessageReaction"s will just create a message that has originMessageId set
	EmojiReactions []Reaction
	// stores the id of the message this is a reply to. if this is not a reply message, the value is initialized to -1.
	OriginMessageId int
	// flag that is set to true if the message has been deleted
	HasBeenDeleted bool
	// flag that is set to true if the message was forwarded
	WasForwarded bool
	// list of userIDs that have read the message
	ReadBy []int
}

type Reaction struct {
	UserWhoReacted int
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
