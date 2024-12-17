package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

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
	if len(requestBody.Username) < 3 || len(requestBody.Username) > 16 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	_token, _tokenGenError := util.GetRandomToken(32)
	if _tokenGenError != nil {
		rt.baseLogger.Errorln("Token generation error!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, userExistsError := rt.db.GetUser(requestBody.Username)

	if userExistsError == nil {

		util.TokenMap[_token] = user.Username
		rt.baseLogger.Println("User ", user.Username, " found in Database, logged in with token: ", _token)

	} else {

		basicPic, picError := util.GetBasicUserPicture()
		if picError != nil {
			rt.baseLogger.Errorln("Error with picture of new user!")
		}

		// create a new user
		user = util.User{
			Username:       requestBody.Username,
			ProfilePicture: basicPic,
		}
		// save new user in the database
		dberr := rt.db.InsertUser(user)
		if dberr != nil {
			rt.baseLogger.Errorln("Saving new User into DB error!")
		}

		util.TokenMap[_token] = user.Username
		rt.baseLogger.Println("New User ", user.Username, " was created and logged in with token: ", _token)
	}

	//fmt.Println("-----Func doLogin Finished-----")
	loggedInUserStruct := util.SessionStruct{
		User:         user,
		SessionToken: _token,
	}

	encodeErr := json.NewEncoder(w).Encode(loggedInUserStruct)

	if encodeErr != nil {
		rt.baseLogger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
