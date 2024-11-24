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

	//username := ps.ByName("Username")

	if UserLoggedIn == nil {
		fmt.Println("User is not logged in!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Read the request body
	var requestBody struct {
		ProfilePicture string `json:"ProfilePicture"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	fmt.Println("User: ", UserLoggedIn.Username, " is trying to change profile picture to: ", requestBody.ProfilePicture)

	// oldUser, oldExists := AllUsers[username]
	// if !oldExists {
	// 	fmt.Println("User ", username, " is not in the database!")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	if UserLoggedIn.ProfilePicture == requestBody.ProfilePicture {
		fmt.Println("Your profile picture is already ", UserLoggedIn.ProfilePicture)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	UserLoggedIn.ProfilePicture = requestBody.ProfilePicture

	fmt.Println("User ", UserLoggedIn.Username, " set its profile picture sucessfully to ", requestBody.ProfilePicture, "!")

	json.NewEncoder(w).Encode(UserLoggedIn)

}
