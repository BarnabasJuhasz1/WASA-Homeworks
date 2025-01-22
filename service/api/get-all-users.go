package api

import (
	"encoding/json"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	ctx.Logger.Debugln("-----Func getAllUsers Called-----")

	searchQuery := r.URL.Query().Get("search")

	var userIDs []int
	var dbErr error
	if searchQuery == "" {
		userIDs, dbErr = rt.db.GetAllUsers()
	} else {
		userIDs, dbErr = rt.db.GetUsersFromQuery(searchQuery)
	}

	if dbErr != nil {
		ctx.Logger.Errorln("Error getting userIDs from database: ", dbErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("getUsers query results: ", userIDs)

	encodeErr := json.NewEncoder(w).Encode(userIDs)

	if encodeErr != nil {
		ctx.Logger.Errorln("Failed to encode to JSON:", encodeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Debugln("-----Func getAllUsers Finished-----")

}
