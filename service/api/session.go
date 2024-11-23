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
	Username       string
	ProfilePicture byte
}

// var Users = []User{}
var Users = make(map[string]User)

// Start a new game generating the secret number
// and return the created game id
func (rt *_router) session(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	fmt.Println("start new session called")

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

	value, exists := Users[requestBody.Username]

	if exists {
		fmt.Println("User ", value, " logged in sucessfully!")
	} else {
		//create a new user
		Users[requestBody.Username] = User{
			Username: requestBody.Username,
		}
		fmt.Println("User ", requestBody.Username, " created sucessfully!")
	}

	json.NewEncoder(w).Encode(requestBody.Username)
}
