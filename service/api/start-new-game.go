package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	//"encoding/json"
	//"math/rand"
)

// Start a new game generating the secret number
// and return the created game id
func (rt *_router) startNewGame(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//w.Header().Set("content-type", "application/json")

	println("start new game func called")

	/*
		id := len(Games)

		Games = append(Games,Game{
			Id: id,
			secret: rand.Intn(100),
			Outcome: "",
			Guesses: 0,
		})

		gamesGuesses = append(gamesGuesses,[]Guess{})

		json.NewEncoder(w).Encode(id)
	*/
}
