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
curl -X POST http://localhost:3000/session -H "Content-Type: application/json" -d '{"Username": "MyName"}' -v

