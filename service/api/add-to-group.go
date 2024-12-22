package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func addToGroup Called-----")

	LoggedInUser := rt.db.GetLoggedInUser(w, ctx)

	// conversationIDString := ps.ByName("ConversationID")

	// conversationID, convErr := strconv.Atoi(conversationIDString)
	// if convErr != nil || conversationID < 0 {
	// 	fmt.Println("Invalid conversationID in path! ", convErr)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// Conversation, existsConv := AllConversations[conversationID]
	// //if the conversation does not exist
	// if !existsConv {
	// 	fmt.Println("Invalid conversationID in path! ", existsConv)
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

	// Read the request body
	var requestBody struct {
		UserIDToAdd int `json:"UserIDToAdd"`
	}

	requestErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if requestErr != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// userToAdd, userExists := util.AllUsers[requestBody.UserNameToAdd]
	userToAdd, userExistsError := rt.db.GetUser(requestBody.UserIDToAdd)
	if userExistsError != nil {
		ctx.Logger.Debugln("User ", requestBody.UserIDToAdd, " is not in the database!")

		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if the person already belongs to the group or not
	// if isUserFoundInList(Conversation.ConversationGroup.Participants, requestBody.UserNameToAdd) {

	// 	fmt.Println("User ", requestBody.UserNameToAdd, " is already in the group!")
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return

	// }

	// make sure user to add does not already belong to the conversation
	if util.UserBelongsToConversation(Conversation, userToAdd) {
		ctx.Logger.Debugln("User is already in the conversation!")

		w.WriteHeader(http.StatusForbidden)
		return
	}

	// add user to the group (group's perspective)
	Conversation.Participants = append(Conversation.Participants, userToAdd.Id)

	// add user to the group (user's perspective)
	// userToAdd.MyConversations = append(userToAdd.MyConversations, Conversation.Id)

	// update users map by reassigning the struct
	// util.AllUsers[requestBody.UserNameToAdd] = userToAdd

	// update conversations map by reassigning the struct
	// util.AllConversations[Conversation.Id] = Conversation
	rt.db.UpdateConversation(Conversation.Id, Conversation)

	ctx.Logger.Debugln("-----Func addToGroup Finished-----")
	encodeErr := json.NewEncoder(w).Encode(Conversation)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
