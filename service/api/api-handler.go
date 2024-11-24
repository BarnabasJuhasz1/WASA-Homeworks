package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.doLogin)

	rt.router.PUT("/user", rt.setUsername)

	rt.router.PUT("/user/profilePicture", rt.setMyPhoto)

	rt.router.GET("/user/myConversations", rt.getMyConversations)

	rt.router.GET("/conversation/:ConversationID", rt.getConversation)
	rt.router.POST("/conversation/:ConversationID", rt.sendMessage)
	rt.router.PUT("/conversation/:ConversationID", rt.setGroupName)
	rt.router.DELETE("/conversation/:ConversationID", rt.leaveGroup)

	rt.router.PUT("/conversation/:ConversationID/add", rt.addToGroup)

	rt.router.POST("/create/conversation", rt.createConversation)

	//rt.router.GET("/", rt.getHelloWorld)
	//rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
