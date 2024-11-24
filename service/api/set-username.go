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

	//oldUsername := ps.ByName("Username")

	if UserLoggedIn == nil {
		fmt.Println("User is not logged in!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Read the request body
	var requestBody struct {
		NewUsername string `json:"NewUsername"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	fmt.Println("User: ", UserLoggedIn.Username, " is trying to change name to: ", requestBody.NewUsername)

	// oldUser, oldExists := AllUsers[UserLoggedIn.Username]
	// if !oldExists {
	// 	fmt.Println("User ", UserLoggedIn.Username, " is not in the database!")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	if UserLoggedIn.Username == requestBody.NewUsername {
		fmt.Println("Your username is already ", UserLoggedIn.Username)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(requestBody.NewUsername) < 3 || len(requestBody.NewUsername) > 16 {
		fmt.Println("Username too short or too long ", requestBody.NewUsername)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	_, newExists := AllUsers[requestBody.NewUsername]

	if newExists {
		fmt.Println("Username ", requestBody.NewUsername, " occupied by someone else!")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	} else {
		// //delete old entry in users
		// delete(Users, oldUsername)
		// //rename old user's name to new value
		// oldUser.Username = requestBody.NewUsername
		// //add same user with the new name
		// Users[requestBody.NewUsername] = oldUser

		//update the username of the user logged in
		//Remark: since UserLoggedIn is a pointer, the "AllUsers" map is updated as well
		UserLoggedIn.Username = requestBody.NewUsername

		fmt.Println("User ", UserLoggedIn.Username, " renamed sucessfully to ", requestBody.NewUsername, "!")
	}

	json.NewEncoder(w).Encode(AllUsers[requestBody.NewUsername])

}
