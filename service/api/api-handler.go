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

	rt.router.DELETE("/conversation/:ConversationID/message/:MessageID", rt.deleteMessage)

	rt.router.POST("/conversation/:ConversationID/message/:MessageID/comment", rt.forwardMessage)
	rt.router.PUT("/conversation/:ConversationID/message/:MessageID/comment", rt.commentMessage)
	rt.router.DELETE("/conversation/:ConversationID/message/:MessageID/comment", rt.uncommentMessage)

	rt.router.PUT("/conversation/:ConversationID/groupPicture", rt.setGroupPhoto)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)
	// rt.router.GET("/context", rt.wrap(rt.getContextReply))

	return rt.router
}
