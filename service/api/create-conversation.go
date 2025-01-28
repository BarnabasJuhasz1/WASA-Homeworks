package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func createConversation Called-----")
	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	// Read the request body
	var requestBody struct {
		ConvType            util.ConversationType `json:"ConversationType"`
		Participants        []int                 `json:"Participants"`
		ConversationName    string                `json:"ConversationName"`
		ConversationPicture []byte                `json:"ConversationPicture"`
	}

	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// if this is a one-on-one conversation
	if requestBody.ConvType == util.UserType {
		// fetch my conversations
		myConversations, fetchMyConvErr := rt.db.GetMyConversations(LoggedInUser.Id)
		if fetchMyConvErr != nil {
			ctx.Logger.Errorln("Failed to fetch user conversations: ", fetchMyConvErr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// check if I already have a conversation with the same person
		for _, conv := range myConversations {
			if conv.Type == util.UserType {
				if conv.Participants[0] == requestBody.Participants[0] ||
					conv.Participants[1] == requestBody.Participants[0] {

					ctx.Logger.Errorln("You already have a conversation with that person!")
					w.WriteHeader(http.StatusForbidden)
					return
				}
			}
		}
	}

	var ConversationUsers []int
	// add the user logged in to the conversation automatically
	ConversationUsers = append(ConversationUsers, LoggedInUser.Id)

	for _, userIDAtI := range requestBody.Participants {

		// check if all the users exist, if so, add them to a list
		userToAdd, userExistsError := rt.db.GetUser(userIDAtI)
		if userExistsError != nil {
			ctx.Logger.Debugln("User", userIDAtI, "is not in the database!")

			w.WriteHeader(http.StatusNotFound)
			return
		}
		ConversationUsers = append(ConversationUsers, userToAdd.Id)
	}

	var emptyMessages []util.Message
	// basicPic, _ := util.GetBasicGroupPicture()
	// create the Conversation with all the users
	Conversation := util.Conversation{
		Type:         requestBody.ConvType,
		GroupName:    requestBody.ConversationName,
		GroupPicture: requestBody.ConversationPicture,
		Participants: ConversationUsers,
		Messages:     emptyMessages,
	}

	// insert new conversation to database
	id, insertingErr := rt.db.InsertConversation(Conversation)
	if insertingErr != nil {
		ctx.Logger.Errorln("Error inserting new conversation into database! ", insertingErr)
	}

	// map myself to the conversation
	dbAddErr := rt.db.AddConversationIDToUser(LoggedInUser.Id, id)
	if dbAddErr != nil {
		ctx.Logger.Errorln("Failed to add user to conversation:", dbAddErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// map the other users to the conversation
	for _, userIDAtI := range requestBody.Participants {

		// ctx.Logger.Debugln("trying to add user: ", userNameAtI, " to conversation ", id)
		dbAddErr = rt.db.AddConversationIDToUser(userIDAtI, id)
		if dbAddErr != nil {
			ctx.Logger.Errorln("Failed to add user to conversation:", dbAddErr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	Conversation.Id = id
	// util.AllConversations[conversation.Id] = conversation
	dberr := rt.db.UpdateConversation(Conversation.Id, Conversation)
	if dberr != nil {
		ctx.Logger.Errorln("Failed to update conversation:", dberr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func createConversation Finished-----")
	encodeErr := json.NewEncoder(w).Encode(Conversation)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
