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
func (rt *_router) setMyPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("Func setMyPhoto Called")

	username := ps.ByName("Username")

	// Read the request body
	var requestBody struct {
		ProfilePicture string `json:"ProfilePicture"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	fmt.Println("User: ", username, " is trying to change profile picture to: ", requestBody.ProfilePicture)

	oldUser, oldExists := Users[username]
	if !oldExists {
		fmt.Println("User ", username, " is not in the database!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if oldUser.ProfilePicture == requestBody.ProfilePicture {
		fmt.Println("Your profile picture is already ", oldUser.ProfilePicture)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oldUser.ProfilePicture = requestBody.ProfilePicture

	fmt.Println("User ", username, " set its profile picture sucessfully to ", requestBody.ProfilePicture, "!")

	json.NewEncoder(w).Encode(Users[username])

}
