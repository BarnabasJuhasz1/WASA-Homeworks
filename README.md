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
curl.exe -X PUT http://localhost:3000/user -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d '{\"NewUsername\": \"NewJuhasz\"}'

setMyPhoto:
curl.exe -X PUT http://localhost:3000/user/profilePicture -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d '{\"ProfilePicture\": \"0000\"}'

getMyConversations:
curl.exe http://localhost:3000/user/myConversations -H "Authorization: Bearer TOKEN"

getConversation:
curl.exe http://localhost:3000/conversation/0 -H "Authorization: Bearer TOKEN"

sendMessage:
curl.exe -X POST http://localhost:3000/conversation/0 -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d '{\"MessageContent\": \"Hello, World!\"}'

setGroupName:
curl.exe -X POST http://localhost:3000/conversation/0 -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d '{\"GroupName\": \"New Group Name\"}'

leaveGroup:
curl.exe -X DELETE http://localhost:3000/conversation/0 -H "Authorization: Bearer TOKEN"

addToGroup:
curl.exe -X PUT http://localhost:3000/conversation/0/add -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d '{\"UserNameToAdd\": \"Natali\"}'

deleteMessage:
curl.exe -X DELETE http://localhost:3000/conversation/0/message/0-H "Authorization: Bearer TOKEN"

forwardMessage:
curl.exe -X POST http://localhost:3000/conversation/0/message/0/comment -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d '{\"RecipientUsername\": \"Natali\"}'

commentMessage:
curl.exe -X PUT http://localhost:3000/conversation/0/message/0/comment -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d '{\"ReactionType\": \"EmojiReaction\", \"Content\": \":)\"}'
curl.exe -X PUT http://localhost:3000/conversation/0/message/0/comment -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d '{\"ReactionType\": \"MessageReaction\", \"Content\": \"This is a reply message :) \"}'

uncommentMessage:
curl.exe -X DELETE http://localhost:3000/conversation/0/message/0/comment -H "Authorization: Bearer TOKEN"

setGroupPhoto:
curl.exe -X PUT http://localhost:3000/conversation/:ConversationID/groupPicture -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d '{\"GroupPicture\": \"0000\"}'

createConversation:
curl.exe -X POST http://localhost:3000/create/conversation -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d '{\"ConversationType\": \"UserType\", \"Participants\": [\"Natali\"]}'


