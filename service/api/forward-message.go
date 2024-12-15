package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func forwardMessage Called-----")

	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	// get the conversation from path
	OriginalConversation, convErr := util.GetConversationFromPath(w, ps, ctx)
	if convErr {

		return
	}

	// make sure the logged in user belongs to the conversation
	if !util.UserBelongsToConversation(OriginalConversation, LoggedInUser) {
		ctx.Logger.Debugln("User is not in the conversation!")

		w.WriteHeader(http.StatusForbidden)
		return
	}

	// get messageID from path
	messageIDString := ps.ByName("MessageID")
	// make sure the messageID is correct
	messageID, messageErr := strconv.Atoi(messageIDString)
	if messageErr != nil || messageID < 0 || messageID >= len(OriginalConversation.Messages) {
		ctx.Logger.Debugln("Invalid MessageID in path! ", messageErr)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read the request body
	var requestBody struct {
		RecipientUsername string `json:"RecipientUsername"`
	}
	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// make sure the recipient exists in the database as a user
	Recipient, userExistsError := rt.db.GetUser(requestBody.RecipientUsername)
	if userExistsError != nil {
		ctx.Logger.Debugln("User ", requestBody.RecipientUsername, " is not in the database!")

		w.WriteHeader(http.StatusNotFound)
		return
	}

	// find the one-on-one conversation you have with the recipient
	ConvWithRecipient, exists := util.GetOneOnOneConversationWithUser(LoggedInUser, Recipient)
	if !exists {

		ctx.Logger.Debugln("New conversation created between: ", LoggedInUser.Username, " and ", Recipient.Username)

		// if no such conversation exist, then create a one-on-one conversation with the person (or you cannot forward the message?)
		var emptyMessages []util.Message
		// create the conversation with the recipient
		ConvWithRecipient = util.Conversation{
			Id: len(util.AllConversations),
			// ConversationGroup: util.Group{
			// 	Participants: []util.User{Recipient},
			// 	GroupName:    "New Conversation",
			// },
			Type:      util.UserType,
			GroupName: "New Conversation",

			// Participants: []util.User{Recipient},
			Participants: []string{Recipient.Username},
			Messages:     emptyMessages,
		}

		// add the new conversation to the users
		LoggedInUser.MyConversations = append(LoggedInUser.MyConversations, ConvWithRecipient.Id)
		// util.AllUsers[LoggedInUser.Username] = LoggedInUser

		Recipient.MyConversations = append(Recipient.MyConversations, ConvWithRecipient.Id)
	}

	var emptyReactions []util.Reaction
	// send the message to the recipient by creating a new message with the same content of the original message
	// and modify one-on-one conversation by adding the new message
	ConvWithRecipient.Messages = append(OriginalConversation.Messages, util.Message{
		Id:              len(ConvWithRecipient.Messages),
		Sender:          LoggedInUser,
		Content:         OriginalConversation.Messages[messageID].Content,
		Timestamp:       time.Now().Format("2006-01-02 15:04:05"),
		Status:          util.UserName,
		EmojiReactions:  emptyReactions,
		OriginMessageId: -1,
	})

	// update conversations map by reassigning the struct
	util.AllConversations[ConvWithRecipient.Id] = ConvWithRecipient

	ctx.Logger.Debugln("-----Func forwardMessage Finished-----")

	encodeErr := json.NewEncoder(w).Encode(OriginalConversation.Messages[messageID])

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
