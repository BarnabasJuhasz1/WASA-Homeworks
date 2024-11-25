package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	fmt.Println("-----Func setMyPhoto Called-----")

	//username := ps.ByName("Username")

	//make sure user is logged in
	if !isUserLoggedIn(w) {
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

	if UserLoggedIn.ProfilePicture == requestBody.ProfilePicture {
		fmt.Println("Your profile picture is already ", UserLoggedIn.ProfilePicture)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	UserLoggedIn.ProfilePicture = requestBody.ProfilePicture

	fmt.Println("User ", UserLoggedIn.Username, " set its profile picture sucessfully to ", requestBody.ProfilePicture, "!")

	fmt.Println("-----Func setMyPhoto Finished-----")
	json.NewEncoder(w).Encode(UserLoggedIn)

}
