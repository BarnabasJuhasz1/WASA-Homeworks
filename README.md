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
curl -X POST http://localhost:3000/session -H "Content-Type: application/json" -d '{"Username": "Juhasz"}'

setMyUserName:
curl -X PUT http://localhost:3000/user/Juhasz -H "Content-Type: application/json" -d '{"Username": "NewJuhasz"}'

setMyPhoto:
curl -X PUT http://localhost:3000/user/Juhasz -H "Content-Type: application/json" -d '{"ProfilePicture": "000.."}'

getMyConversations:
curl http://localhost:3000/user/Juhasz/myConversations

getConversation:
curl http://localhost:3000/conversation/0

sendMessage:
curl -X POST http://localhost:3000/conversation/0 -H "Content-Type: application/json" -d '{"SenderUsername": "Juhasz", "MessageContent": "Hello, World!"}'

setGroupName:
curl -X POST http://localhost:3000/conversation/:0 -H "Content-Type: application/json" -d '{"GroupName": "New Group Name"}'

leaveGroup
curl -X DELETE http://localhost:3000/conversation/:0

addToGroup:
curl -X PUT http://localhost:3000/conversation/:0/add -H "Content-Type: application/json" -d '{"UserNameToAdd": "Natali"}'

createConversation:



