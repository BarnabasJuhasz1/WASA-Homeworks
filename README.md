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

-----------Examples-----------
To Log In:
curl -X POST http://localhost:3000/session -H "Content-Type: application/json" -d '{"Username": "Juhasz"}' -v

To Change Username:
curl -X PUT http://localhost:3000/user/Juhasz -H "Content-Type: application/json" -d '{"Username": "NewJuhasz"}'

To Change Profile Picture:
curl -X PUT http://localhost:3000/user/Juhasz -H "Content-Type: application/json" -d '{"ProfilePicture": "000.."}'

To Get My Conversations List:
curl http://localhost:3000/user/Juhasz/myConversations

To Get My Conversation:
curl http://localhost:3000/conversation/0

To Send Message to group:
curl -X POST http://localhost:3000/conversation/0 -H "Content-Type: application/json" -d '{"SenderUsername": Juhasz, "MessageContent": "Hello, World!"}'

To Set Group Name:
curl -X POST http://localhost:3000/conversation/:0 -H "Content-Type: application/json" -d '{"GroupName": "New Group Name"}'



