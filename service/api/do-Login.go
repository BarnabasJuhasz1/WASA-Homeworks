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

	user, exists := util.AllUsers[requestBody.Username]

	if exists {

		_token, _tokenGenError := util.GetRandomToken(32)
		if _tokenGenError != nil {
			rt.baseLogger.Debug("Token generation error!")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		util.TokenMap[_token] = user.Username
		rt.baseLogger.Println("User ", user.Username, " found in Database, logged in with token: ", _token)

	} else {
		// create a new user
		util.AllUsers[requestBody.Username] = util.User{
			Username: requestBody.Username,
		}
		// save new user in the database
		user = util.AllUsers[requestBody.Username]

		_token, _tokenGenError := util.GetRandomToken(32)
		if _tokenGenError != nil {
			rt.baseLogger.Errorln("Token generation error!")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		util.TokenMap[_token] = user.Username
		rt.baseLogger.Println("New User ", user.Username, " was created and logged in with token: ", _token)
	}

	//fmt.Println("-----Func doLogin Finished-----")

	encodeErr := json.NewEncoder(w).Encode(requestBody.Username)

	if encodeErr != nil {
		rt.baseLogger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
