package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	rt.baseLogger.Println("User Login Requested")

	// Read the request body
	var requestBody struct {
		Username string `json:"Username"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// make sure the username respects the length boundaries
	if len(requestBody.Username) < 3 || len(requestBody.Username) > 16 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// generate the random token to give to user
	_token, _tokenGenError := util.GetRandomToken(32)
	if _tokenGenError != nil {
		rt.baseLogger.Errorln("Token generation error!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// make db query to retrieve user
	user, userExistsError := rt.db.GetUserFromName(requestBody.Username)

	// if the user does exist
	if userExistsError == nil {

		util.TokenMap[_token] = user.Username
		rt.baseLogger.Println("User ", user.Username, " found in Database, logged in with token: ", _token)

	} else {

		// basicPic, picError := util.GetBasicUserPicture()
		// if picError != nil {
		// 	rt.baseLogger.Errorln("Error with picture of new user! ", picError)
		// }

		var emptyProfilePic []byte
		// create a new user
		user = util.User{
			Username:       requestBody.Username,
			ProfilePicture: emptyProfilePic,
		}
		// save new user in the database
		userID, dberr := rt.db.InsertUser(user)
		if dberr != nil {
			rt.baseLogger.Errorln("Saving new User into DB error! ", dberr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// set the ID of the user struct
		user.Id = userID
		// add the user to the token map
		util.TokenMap[_token] = user.Username

		rt.baseLogger.Println("New User ", user.Username, " was created and logged in with token: ", _token)
	}

	// create the response struct
	loggedInUserStruct := util.SessionStruct{
		User:         user,
		SessionToken: _token,
	}

	// encode response
	encodeErr := json.NewEncoder(w).Encode(loggedInUserStruct)

	if encodeErr != nil {
		rt.baseLogger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// gets the conversation from the path.
// if there is an error, the second value of the tuple will be set to true
func GetConversationFromPath(rt *_router, w http.ResponseWriter, ps httprouter.Params, ctx reqcontext.RequestContext) (util.Conversation, bool) {

	conversationIDString := ps.ByName("ConversationID")

	// make sure the conversationID is correct
	conversationID, convErr := strconv.Atoi(conversationIDString)
	if convErr != nil || conversationID < 0 { //|| messageID > len() {
		ctx.Logger.Debugln("Invalid conversationID in path! ", convErr)

		w.WriteHeader(http.StatusBadRequest)
		return util.Conversation{}, true
	}

	// make sure the conversation exists
	ConversationStruct, existsConv := rt.db.GetConversation(conversationID)

	if existsConv != nil {
		ctx.Logger.Debugln("Invalid conversationID in path! ", existsConv)

		w.WriteHeader(http.StatusBadRequest)
		return util.Conversation{}, true
	}

	return ConversationStruct, false
}
