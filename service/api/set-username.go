package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	//"math/rand"
)

// Start a new game generating the secret number
// and return the created game id
func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("Func setUsername Called")

	oldUsername := ps.ByName("Username")

	// Read the request body
	var requestBody struct {
		NewUsername string `json:"NewUsername"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	fmt.Println("User: ", oldUsername, " is trying to change name to: ", requestBody.NewUsername)

	oldUser, oldExists := Users[oldUsername]
	if !oldExists {
		fmt.Println("User ", oldUsername, " is not in the database!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if oldUsername == requestBody.NewUsername {
		fmt.Println("Your username is already ", oldUsername)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(requestBody.NewUsername) < 3 || len(requestBody.NewUsername) > 16 {
		fmt.Println("Username too short or too long ", requestBody.NewUsername)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	_, newExists := Users[requestBody.NewUsername]

	if newExists {
		fmt.Println("Username ", requestBody.NewUsername, " occupied by someone else!")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	} else {
		//delete old entry in users
		delete(Users, oldUsername)
		//rename old user's name to new value
		oldUser.Username = requestBody.NewUsername
		//add same user with the new name
		Users[requestBody.NewUsername] = oldUser

		fmt.Println("User ", oldUsername, " renamed sucessfully to ", requestBody.NewUsername, "!")
	}

	json.NewEncoder(w).Encode(Users[requestBody.NewUsername])

}
