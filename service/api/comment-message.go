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

func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func commentMessage Called-----")

	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

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

	// get messageID from path
	messageIDString := ps.ByName("MessageID")
	// make sure the messageID is correct
	messageID, messageErr := strconv.Atoi(messageIDString)
	if messageErr != nil || messageID < 0 || messageID >= len(Conversation.Messages) {
		ctx.Logger.Debugln("Invalid MessageID in path! ", messageErr)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read the request body
	var requestBody struct {
		TypeOfReaction    util.ReactionType `json:"TypeOfReaction"`
		ContentOfReaction string            `json:"ContentOfReaction"`
	}
	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// create the instance of the reaction struct
	ReactionToMake := util.Reaction{
		UserWhoReacted: LoggedInUser.Username,
		Type:           requestBody.TypeOfReaction,
		Content:        requestBody.ContentOfReaction,
	}

	// if the reaction is an emoji reaction, just append it to the list of reactions on the message
	if ReactionToMake.Type == util.EmojiReaction {

		// if there are no emoji reactions on this message
		if len(Conversation.Messages[messageID].EmojiReactions) == 0 {

			Conversation.Messages[messageID].EmojiReactions = append(Conversation.Messages[messageID].EmojiReactions, ReactionToMake)

		} else {

			// loop through all the emoji reactions to this message
			for i, ReactionAti := range Conversation.Messages[messageID].EmojiReactions {

				// if the user has a reaction to this message already, replace that reaction with the new one
				if ReactionAti.UserWhoReacted == LoggedInUser.Username {

					messageToReplace := Conversation.Messages[messageID]
					// if you are placing a new emoji on the message
					if ReactionAti.Content != ReactionToMake.Content {
						// replace old reaction content with new one
						messageToReplace.EmojiReactions[i] = ReactionToMake

					} else {
						// remove my reaction
						messageToReplace.EmojiReactions = append(
							messageToReplace.EmojiReactions[:i],
							messageToReplace.EmojiReactions[i+1:]...,
						)
					}
					// update reaction in the message
					Conversation.Messages[messageID] = messageToReplace
					break

					// if we checked all the reactions and none of them were created by the user
				} else if i == len(Conversation.Messages[messageID].EmojiReactions)-1 {

					Conversation.Messages[messageID].EmojiReactions = append(Conversation.Messages[messageID].EmojiReactions, ReactionToMake)
				}
			}
		}

		// if the reaction is a reply (=comment), then create a new message
	} else if ReactionToMake.Type == util.MessageReaction {

		var emptyReactions []util.Reaction
		// create a new message with the content of the reaction,
		// and add it to the conversation
		Conversation.Messages = append(Conversation.Messages, util.Message{
			Id:              len(Conversation.Messages),
			Sender:          LoggedInUser.Id,
			Content:         ReactionToMake.Content,
			Timestamp:       time.Now().Format("2006-01-02 15:04:05"),
			Status:          util.UserName,
			EmojiReactions:  emptyReactions,
			OriginMessageId: messageID,
		})
	}

	// update conversations map by reassigning the struct
	// util.AllConversations[Conversation.Id] = Conversation
	dberr := rt.db.UpdateConversation(Conversation.Id, Conversation)
	if dberr != nil {
		ctx.Logger.Errorln("Failed to update conversation:", dberr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func commentMessage Finished-----")

	var encodeErr error
	if ReactionToMake.Type == util.EmojiReaction {
		encodeErr = json.NewEncoder(w).Encode(Conversation.Messages[messageID])
	} else {
		encodeErr = json.NewEncoder(w).Encode(Conversation.Messages[messageID+1])
	}

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
