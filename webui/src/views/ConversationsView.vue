
<script>
import MessagesList from '../components/MessageList.vue';
import GroupList from '../components/GroupList.vue';
import CurrentGroupHeader from '../components/CurrentGroupHeader.vue';
import CurrentProfile from '../components/CurrentProfile.vue';
import AddGroup from '../components/AddGroup.vue';
import OriginMessage from '../components/OriginMessage.vue';
import axios from "../services/axios.js"

import ContextMenu from '../components/ContextMenu.vue';

import { sharedData } from '../services/sharedData.js';

import 'emoji-picker-element';

export default 
{
  props: {
    myConversations:{
      type: Object,
      default: null,
    }
  },
  setup() {
    // Log sharedData.UserSession
    console.log("Accessing UserSession in ConversationsView:", sharedData.UserSession);

    return {
      UserSession: sharedData.UserSession, // so we can use it in the template 
    };
  },
  data() {
    return {
        custom_colors : ["0, 31, 84"],

        showOverlay: false,
        currentMessage: "",
        selectedConversationIndexLocal : 0,

        contextMenuVisible: false,
        selectedMessageID: null,

        emojiPanelVisible: false,

        originMessage: null,

        };
    },
    methods: {
        sendMessage(event) {
            if(event)
              event.preventDefault();

            if((this.currentMessage).trim() == ""){
              this.currentMessage = "";
              return
            }

            const now = new Date();

            const newID = this.selectedConversation.Messages ? this.selectedConversation.Messages.length : 0
            // create a new message and push it to the messages
            const newMessage = {
              Id: newID,
              Sender: sharedData.UserSession.UserID,
              Content: this.currentMessage,
              Timestamp: now,
              SentByUser: true,
            }
            // console.log("about to send a message (USERID: ), ", sharedData.UserSession.UserID);
            // send the message content to backend
            this.sendMessageRequest();

            if (!this.selectedConversation.Messages)
              this.selectedConversation.Messages = [];

            this.selectedConversation.Messages.push(newMessage);
            // window.location.reload();

            // reset textArea input
            this.currentMessage = "";

            this.$nextTick(() => {
              this.scrollToBottom();
              this.adjustHeight();
            });

        },
        async sendMessageRequest() {
          try {
            // console.log("attempting to send messsage: ", this.currentMessage, " to id: ", this.myConversations[this.selectedConversationIndexLocal].Id)

            let response = await axios.post(
              "/conversation/"+this.myConversations[this.selectedConversationIndexLocal].Id, 
              // JSON body:
              {
                MessageContent: this.currentMessage,
              },
              // Headers:
              {
                headers: {
                  "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                  "Content-Type": "application/json",
                },
              }
            );

            // console.log("message sent with response: ", response.data);
            
          }
          catch (error) {
            console.error("Error sending message! ", error);
            alert("Error sending message!")
          }
        },
        scrollToBottom() {
            const scrollContainer = document.getElementById("MessagesList");
            // const scrollContainer = this.$refs.messagesList;
            scrollContainer.scrollTop = scrollContainer.scrollHeight;
        },
        resetCurrentText(){
            this.currentMessage = "";
        },
        SelectNewConversationInApp(index){
          
          if(this.selectedConversationIndexLocal != index)
          {
            this.originMessage = null;               
            this.selectedConversationIndexLocal = index
            this.currentMessage = "";
          }

          this.$nextTick(() => {
            this.scrollToBottom();
            this.adjustHeight();

              const groupHeader = this.$refs.groupHeaderRef;
              groupHeader.getProfile();
          });

        },
        adjustHeight() {
          console.log("adjusting height ")

          const textarea = document.getElementById("currentTextArea");
          const originMessage = document.getElementById("OriginMessage");

          if (textarea) {
            const originalHeight = parseInt(textarea.style.height, 10) || 0; // Parse previous height or default to 0
            textarea.style.height = "auto";
            const newHeight = Math.min(textarea.scrollHeight, 200); // New height
            textarea.style.height = `${newHeight}px`;

            const difference = newHeight - originalHeight; // Calculate height change

            const currentTextAreaBottom = parseInt(window.getComputedStyle(textarea).bottom, 10) || 0;
            textarea.style.position = "relative";
            textarea.style.bottom = `${currentTextAreaBottom + difference/2}px`;

            console.log("Original Height:", originalHeight, "New Height:", newHeight);

            if (originMessage) {
              const currentBottom = parseInt(window.getComputedStyle(originMessage).bottom, 10) || 0;

              originMessage.style.position = "relative";
              originMessage.style.bottom = `${currentBottom + difference}px`;
            }
          }
        },
        openOverlayInMode(mode, overlayProfileText, overlayProfilePicture) {
          this.$emit("openOverlayInMode", mode, overlayProfileText, overlayProfilePicture);
        },
        openOverlayInGroupMode(conversation){
          this.$emit("openOverlayInGroupMode", conversation, this.selectedConversationIndexLocal);
        },
        onPageRefresh() {
          
          this.$nextTick(() => {
            this.scrollToBottom();
            this.adjustHeight();

            this.originMessage = null;
          });
          
        },
        openContextMenu(messageID) {
          this.selectedMessageID = messageID;
          // console.log("selecting message: ", messageID)
          this.contextMenuVisible = true;
          this.$refs.contextMenu.show(event.clientX, event.clientY, messageID, this.selectedConversation.Messages[messageID].Sender);
          document.addEventListener('click', this.closeContextMenu);
        },
        closeContextMenu() {
          this.$refs.contextMenu.hide();
          document.removeEventListener('click', this.closeContextMenu);
          this.contextMenuVisible = false;
        },
        refreshLocalMessage(newMessage){
          this.selectedConversation.Messages[newMessage.Id] = newMessage
        },
        openEmojis(){
          this.emojiPanelVisible = !this.emojiPanelVisible
        },
        addEmojiToCurrentMessage(emoji){
          this.currentMessage += emoji.detail.emoji.unicode;
        },
        setOriginMessage(messageID){
          this.originMessage = this.selectedConversation.Messages[messageID];
        }
    },
    mounted() {
      // this.$nextTick(() => {
      //   this.adjustHeight();
      //   this.scrollToBottom();
      // });
      // setTimeout(() => {
      //   this.scrollToBottom();
      // }, 100);
    },
    components: {
        MessagesList,
        GroupList,
        CurrentGroupHeader,
        CurrentProfile,
        AddGroup,
        ContextMenu,
        OriginMessage
    },
    computed: {
      selectedConversation()
      {
        // console.log("my conversations: ", this.myConversations)
        return this.myConversations && this.myConversations[this.selectedConversationIndexLocal]
        ? this.myConversations[this.selectedConversationIndexLocal] : null;
      },
      currentMessageText()
      {
          return this.currentMessage;
      },
    }
  }
</script>

<template>
    <div id="body" class="Flexbox">
        <div id="mainLeft">
            <div id="groupListParent" class="custom-scrollbar">
              <GroupList v-if="myConversations != null"
              :conversations="myConversations"
              @SelectNewConversation="SelectNewConversationInApp"/>
            </div>

            <div id="AddNewGroup">
              <AddGroup @openOverlayInMode="openOverlayInMode"/>
            </div>
          
            <CurrentProfile
            :username="UserSession.Username"
            :profilePicture="UserSession.ProfilePicture"

            @openOverlayInMode="openOverlayInMode"
            />
        </div>

        <div id="main" v-if="myConversations != null">

            <CurrentGroupHeader
            id="CurrentGroupHeader"
            ref="groupHeaderRef"
            :selectedConversation="selectedConversation"
            @openOverlayInGroupMode="openOverlayInGroupMode"
            />

            <MessagesList
            id="MessagesList"
            ref="messagesList"
            :textMessages="this.selectedConversation.Messages"
            :convType="this.selectedConversation.Type"
            :refreshKey="this.selectedConversationIndexLocal"
            @onPageRefresh="onPageRefresh"
            @openContextMenu="openContextMenu"
            />

            <div id="BottomPartWrapper">

              <OriginMessage id="OriginMessage" ref="OriginMessage" style="display: flex;"
              v-if="this.originMessage != null"
              :convType="this.selectedConversation.Type"
              :message="this.originMessage"
              />
          

              <div id="TextAndSend" class="Flexbox" style="gap: 5px;">
                  
                <button id="sendButton" @click="openEmojis()" class="sendButtonImageContainer"> 
                  <img src="https://cdn-icons-png.flaticon.com/128/11202/11202612.png"/>
                </button>

                <emoji-picker id="EmojiPicker"
                  v-if="emojiPanelVisible"
                  @emoji-click="addEmojiToCurrentMessage"
                />

                <textarea id="currentTextArea"
                rows="1"
                v-model="currentMessage"
                placeholder="Type a message"
                @keydown.enter="sendMessage"
                @input="adjustHeight"
                class="custom-scrollbar"
                ></textarea>

                <button id="sendButton" @click="sendMessage()" class="sendButtonImageContainer"> 
                  <img src="https://cdn-icons-png.flaticon.com/128/561/561226.png"
                  style="margin-right: 2.5px;"
                  />
                </button>

              </div>
            </div>

            <ContextMenu ref="contextMenu"
            v-show="contextMenuVisible"
            :conversationID="this.selectedConversation.Id"
            @click="closeContextMenu"
            @refreshLocalMessage="refreshLocalMessage"
            @setOriginMessage="setOriginMessage"
            />

        </div>

    </div>
</template>

<style>


#mainLeft {
  display: block; 
  margin-left: auto;
  margin-top: -105px;
  width: 200px;
  height: calc(85vh - 105px);
}

#groupListParent {
  display: block; 
  margin-left: auto;
  padding: 5px;
  /*height: calc(70vh - 65px - 46px);*/
  height: 100%;
  border-radius: 15px;
  border-bottom-left-radius: 0;
  border-bottom-right-radius: 0;
  background-color: rgba(0, 0, 0, 0.25);

  border: 3px solid rgba(0, 0, 0, .25);
  border-bottom: 0px;

  overflow-y: auto;
}

#main{
  display: flex; 
  /*margin-top: auto;*/
  margin-right: auto;
  padding: 5px;
  flex-direction: column;
  background-color: rgba(0, 0, 0, .25);
  /*height: calc(70vh - 6px);*/

  
  width: 70vw;
  border-radius: 15px;

  border: 3px solid rgba(0, 0, 0, .25);
  height: calc(85vh);

  /*flex-grow: 1;*/

}

#CurrentGroupHeader {
  display: flex;
  height: 75px;
  min-height: 75px;
  max-height: 75px;

  background-color: rgba(0, 0, 0, .5);
  margin-bottom: 5px;
  border-radius: 15px;
}

#MessagesList {
  display: flex;
  flex-direction: column;
  
  margin-bottom: auto;
  overflow-y: auto;

  /*height: 50vh;*/

  /*height: calc(85vh - 150px);*/
  max-height: calc(85vh - 150px);

  flex-grow: 1;
  flex-shrink: 1;
  flex-basis: 0;

  margin-bottom: 10px;
  /*
  background-color: rgba(255, 0, 0, .25);
  flex-grow: 1;*/

}

#TextAndSend {
  margin-top: auto;
  margin-bottom: 0;

  height: 75px;
}

.Flexbox {
  display: flex;
  align-items: center;
  gap: 5px;
}


#currentTextArea {
  
  display: block;
  position: fixed;

  width: 100%;
  /*height: 15px;*/
  max-height: 100px;

  margin-top: 25px;
  margin-bottom: -45px;
  
  resize: none;
  border-radius: 15px;
  border: 0;
  
  padding-left: 15px;
	padding-right: 20px;
	padding-top: 15px;
	padding-bottom: 15px;

  box-sizing: border-box; 

  background-color: var( --message-box);
  outline: none;

  /*color: var(--font-light);*/
  color: rgb(0, 0, 0)

}

#currentTextArea::placeholder {
    color: rgb(0, 0, 0);
}

.sendButtonImageContainer {
  padding: 5px;

  width: 30px;
  height: 30px;

	min-width: 30px;
	min-height: 30px;

  border-radius: 50%;
  overflow: hidden;
  display: flex;
  justify-content: center;
  align-items: center;

}

.sendButtonImageContainer img {
  width: 75%; 
  height: 75%;
  object-fit: cover;
}

#sendButton {
  width: 50px;
  height: 50px;
  min-width: 50px;
  min-height: 50px;
  /* background-color: rgb(255, 209, 0);
  border: 0; */
  background-color: var(--background-light);
  border: 2px solid  rgb(255, 209, 0);

  border-radius: 15px;
  font-weight: bold;
  
  outline: none;

  margin-top: auto;
  margin-bottom: 0;
}




.custom-scrollbar::-webkit-scrollbar {
  width:2px;
  padding: 15px;
}


.custom-scrollbar::-webkit-scrollbar-track {
  background:rgba(28, 28, 28, 0);
  border-radius:15px;
  margin-bottom: 5px; /* Optional: Add spacing at the bottom */
}


.custom-scrollbar::-webkit-scrollbar-thumb {
  background:rgb(161, 161, 161);
  border-radius:15px;
}




.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.blurred-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.4);
  backdrop-filter: blur(8px); /* Apply the blur effect */
  z-index: 1;
}

.modal {
  position: relative;
  z-index: 2;
  padding: 20px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  text-align: center;
}

#EmojiPicker {
  display: absolute;
  position: fixed;
  background: white;
  border: 2px solid white;
  border-radius: 5px;
  margin-bottom: 450px;
}


#OriginMessage {
  display: flex;
  position: relative;

  height: auto;
  max-height: 100px;

  background-color: var(--origin-message-bg);
  border: 2px solid var(--origin-message-border);

  border-radius: 15px;

  margin-right: 55px;
  margin-left: 55px;

  margin-bottom: -20px;
}


#BottomPartWrapper {

  display: flex;
  flex-direction: column;
  position: relative;

}


</style>