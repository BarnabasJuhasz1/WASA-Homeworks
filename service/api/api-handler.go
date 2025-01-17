package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.doLogin)

	rt.router.PUT("/user", rt.wrap(rt.setUsername))
	// add to API!
	rt.router.GET("/user", rt.wrap(rt.getUser))
	rt.router.GET("/user/fromName", rt.wrap(rt.getUserFromName))
	rt.router.GET("/user/all", rt.wrap(rt.getAllUsers))

	rt.router.PUT("/user/profilePicture", rt.wrap(rt.setMyPhoto))

	rt.router.GET("/user/myConversations", rt.wrap(rt.getMyConversations))

	rt.router.GET("/conversation/:ConversationID", rt.wrap(rt.getConversation))
	rt.router.POST("/conversation/:ConversationID", rt.wrap(rt.sendMessage))
	rt.router.PUT("/conversation/:ConversationID", rt.wrap(rt.setGroupName))
	rt.router.DELETE("/conversation/:ConversationID", rt.wrap(rt.leaveGroup))

	rt.router.PUT("/conversation/:ConversationID/groupPicture", rt.wrap(rt.setGroupPhoto))

	rt.router.PUT("/conversation/:ConversationID/add", rt.wrap(rt.addToGroup))

	rt.router.POST("/create/conversation", rt.wrap(rt.createConversation))

	rt.router.DELETE("/conversation/:ConversationID/message/:MessageID", rt.wrap(rt.deleteMessage))

	rt.router.POST("/conversation/:ConversationID/message/:MessageID/comment", rt.wrap(rt.forwardMessage))
	rt.router.PUT("/conversation/:ConversationID/message/:MessageID/comment", rt.wrap(rt.commentMessage))
	rt.router.DELETE("/conversation/:ConversationID/message/:MessageID/comment", rt.wrap(rt.uncommentMessage))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
