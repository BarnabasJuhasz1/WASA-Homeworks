package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"

	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// httpRouterHandler is the signature for functions that accepts a reqcontext.RequestContext in addition to those
// required by the httprouter package.
type httpRouterHandler func(http.ResponseWriter, *http.Request, httprouter.Params, reqcontext.RequestContext)

// wrap parses the request and adds a reqcontext.RequestContext instance related to the request.
func (rt *_router) wrap(fn httpRouterHandler) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		reqUUID, err := uuid.NewV4()
		if err != nil {
			rt.baseLogger.WithError(err).Error("can't generate a request UUID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var ctx = reqcontext.RequestContext{
			ReqUUID: reqUUID,
		}

		// Create a request-specific logger
		ctx.Logger = rt.baseLogger.WithFields(logrus.Fields{
			"reqid":     ctx.ReqUUID.String(),
			"remote-ip": r.RemoteAddr,
		})

		if logger, ok := rt.baseLogger.(*logrus.Logger); ok {
			fmt.Println("Log level:", logger.GetLevel())
		} else {
			fmt.Println("baseLogger is not a *logrus.Logger")
		}

		// Extract and validate the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			ctx.Logger.Warn("missing Authorization header")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check if the header starts with "Bearer "
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			ctx.Logger.Warn("invalid Authorization header format")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract the token
		token := strings.TrimPrefix(authHeader, bearerPrefix)

		// Validate the token
		userID, err := validateToken(token)
		if err != nil {
			ctx.Logger.WithError(err).Warn("invalid token")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Set the RequestContext fields
		ctx.UserID = userID
		ctx.Authenticated = true

		// Call the next handler in chain (usually, the handler function for the path)
		fn(w, r, ps, ctx)
	}
}

// validateToken looks up the token in tokenMap
func validateToken(token string) (string, error) {
	if userID, exists := util.TokenMap[token]; exists {
		return userID, nil
	}

	for key, val := range util.TokenMap {
		fmt.Println("token: ", key, " Username: ", val)
	}

	return "", errors.New("invalid token")
}

