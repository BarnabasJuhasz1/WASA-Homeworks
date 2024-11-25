-----------WASAText-----------
Connect with your friends effortlessly using WASAText! Send and receive messages, whether one-on-one
or in groups, all from the convenience of your PC. Enjoy seamless conversations with text or GIFs and
easily stay in touch through your private chats or group discussions.

-----------OpenAPI-----------
For the project requirements, the following API operations were defined:
• doLogin
• setMyUserName
• getMyConversations
• getConversation
• sendMessage
• forwardMessage
• commentMessage
• uncommentMessage
• deleteMessage
• addToGroup
• leaveGroup
• setGroupName
• setMyPhoto
• setGroupPhoto

• createConversation

-----------Examples-----------
doLogin:
curl.exe -X POST http://localhost:3000/session -H "Content-Type: application/json" -d '{\"Username\": \"Juhasz\"}'

setMyUserName:
curl.exe -X PUT http://localhost:3000/user -H "Content-Type: application/json" -d '{\"NewUsername\": \"NewJuhasz\"}'

setMyPhoto:
curl.exe -X PUT http://localhost:3000/user -H "Content-Type: application/json" -d '{\"ProfilePicture\": \"0000\"}'

getMyConversations:
curl.exe http://localhost:3000/user/myConversations

getConversation:
curl.exe http://localhost:3000/conversation/0

sendMessage:
curl.exe -X POST http://localhost:3000/conversation/0 -H "Content-Type: application/json" -d '{\"SenderUsername\": \"Juhasz\", \"MessageContent\": \"Hello, World!\"}'

setGroupName:
curl.exe -X POST http://localhost:3000/conversation/0 -H "Content-Type: application/json" -d '{\"GroupName\": \"New Group Name\"}'

leaveGroup:
curl.exe -X DELETE http://localhost:3000/conversation/0

addToGroup:
curl.exe -X PUT http://localhost:3000/conversation/0/add -H "Content-Type: application/json" -d '{\"UserNameToAdd\": \"Natali\"}'

deleteMessage:
curl.exe -X DELETE http://localhost:3000/conversation/0/message/0

forwardMessage:
curl.exe -X POST http://localhost:3000/conversation/0/message/0/comment -H "Content-Type: application/json" -d '{\"RecipientUsername\": \"Natali\"}'

commentMessage:
curl.exe -X PUT http://localhost:3000/conversation/0/message/0/comment -H "Content-Type: application/json" -d '{\"ReactionType\": \"EmojiReaction\", \"Content\": \":)\"}'
curl.exe -X PUT http://localhost:3000/conversation/0/message/0/comment -H "Content-Type: application/json" -d '{\"ReactionType\": \"MessageReaction\", \"Content\": \"This is a reply message :) \"}'

uncommentMessage:
curl.exe -X DELETE http://localhost:3000/conversation/0/message/0/comment

setGroupPhoto:
curl.exe -X PUT http://localhost:3000/conversation/:ConversationID/groupPicture -H "Content-Type: application/json" -d '{\"GroupPicture\": \"0000\"}'

createConversation:
curl.exe -X POST http://localhost:3000/create/conversation -H "Content-Type: application/json" -d '{\"ConversationType\": \"UserType\", \"Participants\": [\"Juhasz\", \"Natali\"]}'





