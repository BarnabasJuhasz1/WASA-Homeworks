
<script>
import MessagesList from '../components/MessageList.vue';
import GroupList from '../components/GroupList.vue';
import CurrentGroupHeader from '../components/CurrentGroupHeader.vue';
import CurrentProfile from '../components/CurrentProfile.vue';
import AddGroup from '../components/AddGroup.vue';
import OriginMessage from '../components/OriginMessage.vue';
import axios from "../services/axios.js"
// import { useRoute } from 'vue-router';

import ContextMenu from '../components/ContextMenu.vue';

import { sharedData } from '../services/sharedData.js';
import { ref } from "vue";

import 'emoji-picker-element';
import PopUpReactions from '../components/PopUpReactions.vue';

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

    return {
      UserSession: sharedData.UserSession, // so we can use it in the template 
    };
  },
  watch: {
    // Watch for changes in route parameters
    '$route'(to) {
      // console.log("Updated ROUTER PARAMS: ", to.params.id);
      // this.getConversation(to.params.id);
      this.SelectNewConversationInApp(to.params.id)
    },
    
  },
  data() {
    return {
        custom_colors : ["0, 31, 84"],

        showOverlay: false,
        currentMessage: "",
        selectedConversationIndexLocal : 0,

        contextMenuVisible: false,
        popUpReactionsVisible: false,
        selectedMessageID: null,

        emojiPanelVisible: false,

        originMessage: null,

        intervalId: null,

        };
    },
    beforeUnmount() {
      clearInterval(this.intervalId);
    },
    methods: {
        startPolling() {
          this.intervalId = setInterval(() => {
            // console.log("auto-refresh conv:", this.myConversations)
            this.RefreshSelectedConversation();
          }, 1000);
          // console.log("Polling started with ID:", this.intervalId);
        },
        stopRefreshing(){
          clearInterval(this.intervalId);
          this.intervalId = null;
          // console.log("Refreshing stopped at convView!");
        },
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
              OriginMessageId: this.originMessage == null ? -1 : this.originMessage.Id,
            }
            // console.log("about to send a message (USERID: ), ", sharedData.UserSession.UserID);
            // send the message content to backend
            // console.log("New Message Sent: ", newMessage)

            if (!this.selectedConversation.Messages){
              this.selectedConversation.Messages = [];
            }

            this.selectedConversation.Messages.push(newMessage);

            if(this.originMessage == null){
              this.sendMessageRequest(newID);
            }
            else{
              this.sendMessageReplyRequest(newID);
            }

            // window.location.reload();

            // reset textArea input
            this.currentMessage = "";
            // stop replying to messages
            this.originMessage = null;

            this.$nextTick(() => {
              this.scrollToBottom();
              this.adjustHeight();
            });

        },
        async sendMessageRequest(newMessageID) {
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
            
            // console.log("new message has been received on server: ", response.data)
            setTimeout(() => {
              this.selectedConversation.Messages[newMessageID] = response.data
            }, 150)
            // console.log("message sent with response: ", response.data);
            
          }
          catch (error) {
            console.error("Error sending message! ", error);
            alert("Error sending message!")
          }
        },
        async sendMessageReplyRequest(newMessageID) {
          try {
            // console.log("sending to conv: ", this.myConversations[this.selectedConversationIndexLocal].Id, " with msg origin ID: ", this.originMessage.Id, " with token: ", sharedData.UserSession.SessionToken)
            let response = await axios.put(
              "/conversation/"+this.myConversations[this.selectedConversationIndexLocal].Id
              + "/message/"+ this.originMessage.Id
              + "/comment", 
              // JSON body:
              {
                TypeOfReaction: "MessageReaction",
		            ContentOfReaction: this.currentMessage,
              },
              // Headers:
              {
                headers: {
                  "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                  "Content-Type": "application/json",
                },
              }
            );

            // console.log(response.data)
            // console.log("new reply has been received on server: ", response.data)

            setTimeout(() => {
              this.selectedConversation.Messages[newMessageID] = response.data
            }, 150)
            // console.log("message sent with response: ", response.data);
            
          }
          catch (error) {
            console.error("Error sending reply message! ", error);
            alert("Error sending reply message!")
          }
        },
        async getConversation(convID) {

          //console.log("myconv length: ", this.myConversations.length, " and trying to get: ", convID)
          //if(this.myConversations.length < convID)
            //return;

          try {
            let response = await axios.get(
              "/conversation/"+ convID,
              // Headers:
              {
                headers: {
                  "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                  "Content-Type": "application/json",
                },
              }
            );

            // console.log("conv before: ", this.myConversations[this.selectedConversationIndexLocal])
            // console.log("CONV RECEIVED: ", response.data)

            const updatedConversation = response.data;
            // console.log("COMPARING: ", this.myConversations[this.selectedConversationIndexLocal].Messages.length
            //   , " with new:", updatedConversation.Messages.length)
          
            // let oldMessageCount = this.myConversations[this.selectedConversationIndexLocal].Messages.length
            this.$emit('update-conversation', this.selectedConversationIndexLocal, updatedConversation);
            
            // if(oldMessageCount != updatedConversation.Messages.length){
            //   console.log("NEW MSG! SCROLL")
            //   setTimeout(() => {
            //     this.scrollToBottom();
            //     this.adjustHeight();
            //   }, 1000);
            // }
            // this.myConversations[this.selectedConversationIndexLocal] = response.data;
          }
          catch (error) {
            console.error("Error getting conversation! ", error);
            alert("Error getting conversation!")
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
        SelectNewConversationInApp(localConvIndex){
          
          // this.$router.push('/conversation/'+this.myConversations[localConvIndex].Id);
          this.$router.push('/conversation/'+localConvIndex);

          // let oldMessageCount = this.myConversations[localConvIndex].Messages.length;
          // console.log("old message count: ", oldMessageCount)
          this.getConversation(this.myConversations[localConvIndex].Id);

          if(this.selectedConversationIndexLocal != localConvIndex)
          {
            // console.log("local index changed from ", this.selectedConversationIndexLocal, " to ", localConvIndex)
            this.originMessage = null; 
            this.selectedConversationIndexLocal = localConvIndex
            this.currentMessage = "";

            // this.$nextTick(()=>
            // {
            //   this.adjustHeight();
            //   this.scrollToBottom();
            // });
            setTimeout(() => {
              this.scrollToBottom();
              this.adjustHeight();
            }, 100);
          }
        },
        async RefreshSelectedConversation(){
          if(this.myConversations == null || this.myConversations[this.selectedConversationIndexLocal] == null){
            return;
          }

          if(this.myConversations[this.selectedConversationIndexLocal].Messages == null){
            await this.getConversation(this.myConversations[this.selectedConversationIndexLocal].Id);
            return;
          }

          // this.$refs.messagesList.refreshMessageList();

          let oldMessageCount = this.myConversations[this.selectedConversationIndexLocal].Messages.length;
          
          await this.getConversation(this.myConversations[this.selectedConversationIndexLocal].Id);

          if(oldMessageCount != this.myConversations[this.selectedConversationIndexLocal].Messages.length){
            setTimeout(() => {
              this.scrollToBottom();
              this.adjustHeight();
            }, 100);
          }
        },
        adjustHeight() {

          const textarea = document.getElementById("currentTextArea");
          const spaceForBottom = document.getElementById("spaceForBottom");
          const originMessage = document.getElementById("OriginMessage");

          if (textarea) {
            const originalHeight = parseInt(textarea.style.height, 10) || 0;
            const originalTextAreaTop = parseInt(window.getComputedStyle(textarea).top, 10) || 0;

            textarea.style.height = "auto";
            const newHeight = Math.min(textarea.scrollHeight, 100);
            textarea.style.height = `${newHeight}px`;

            const difference = newHeight - originalHeight;

            textarea.style.position = "relative";
            textarea.style.top = `${originalTextAreaTop - difference}px`;

         
            const spaceForBottomHeight = parseInt(spaceForBottom.style.height, 10) || 0;
            if(difference < 50)
              spaceForBottom.style.height = `${spaceForBottomHeight + difference}px`;

            if (originMessage) {
              const currentBottom = parseInt(window.getComputedStyle(originMessage).top, 10) || 0;
              // console.log("origin top: ", currentBottom, " diff: ", difference)
              originMessage.style.position = "relative";
              originMessage.style.top = `${currentBottom - difference}px`;
            }
       
          }
        },
        openOverlayInMode(mode, overlayProfileText, overlayProfilePicture) {
          this.$emit("openOverlayInMode", mode, overlayProfileText, overlayProfilePicture);
          //console.log("conv: ", this.myConversations)
        },
        openOverlayInGroupMode(conversation){
          this.$emit("openOverlayInGroupMode", conversation, this.selectedConversationIndexLocal);
        },
        openForwardOverlay(message){
          this.$emit("openForwardOverlay", this.selectedConversation, message);
        },
        onPageRefresh() {
          
          this.$nextTick(() => {
            this.scrollToBottom();
            this.adjustHeight();

            this.originMessage = null;

            this.SelectNewConversationInApp(this.selectedConversationIndexLocal)
          });
          
        },
        openContextMenu(messageID) {
          this.closeContextMenu();
          // console.log("selecting message: ", messageID)
          this.selectedMessageID = messageID;
          this.contextMenuVisible = true;
          this.$refs.contextMenu.show(event.clientX, event.clientY, this.selectedConversation.Messages[messageID]);
          document.addEventListener('click', this.closeContextMenu);
        },
        closeContextMenu() {
          if(this.contextMenuVisible){
            this.$refs.contextMenu.hide();
            this.contextMenuVisible = false;
          }
          if(this.popUpReactionsVisible){
            this.$refs.reactionsMenu.hide();
            this.popUpReactionsVisible = false;
          }
          document.removeEventListener('click', this.closeContextMenu);
        },
        openReactionsMenu(messageID){
          this.closeContextMenu();
          // console.log("selecting message: ", messageID)
          this.selectedMessageID = messageID;
          this.popUpReactionsVisible = true;
          // console.log("WOW-1 ", this.popUpReactionsVisible)
          this.$refs.reactionsMenu.show(event.clientX, event.clientY, this.selectedConversation.Messages[messageID]);
          this.$nextTick(() => {
            setTimeout(() => {
              document.addEventListener('click', this.closeContextMenu);
            }, 0.01);
          });
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

          this.$nextTick(() => {
            const textarea = document.getElementById("currentTextArea");
            const originMessage = document.getElementById("OriginMessage");
            const currentBottom = parseInt(window.getComputedStyle(textarea).top, 10) || 0;

            originMessage.style.position = "relative";
            originMessage.style.top = `${currentBottom + 50}px`;
          });
        },
        logout(){
          this.$router.push('/');
          clearInterval(this.intervalId);
        },
        openAttach(){
          this.$refs.fileInput.click();
        },
        uploadPicture(event) {
          const file = event.target.files[0];
          if (file) {
            const reader = new FileReader();

            // Read file as binary string
            reader.readAsArrayBuffer(file);

            // When the file is loaded
            reader.onload = (e) => {
              const arrayBuffer = e.target.result;
              const base64String = btoa(
                  new Uint8Array(arrayBuffer)
                      .reduce((data, byte) => data + String.fromCharCode(byte), "")
              );
              // this.currentMessage = base64String;
              this.currentMessage = `data:image/png;base64,${base64String}`;
              // console.log("File uploaded as Base64:", this.currentMessage);
              this.sendMessage()
            };

            reader.onerror = (e) => {
                console.error("Error reading file:", e);
                alert("Error reading file!")
            };
          }
        },
        cancelReply(){
          this.originMessage=null;
        },
        
    },
    mounted() {
      this.startPolling();
    },
    components: {
        MessagesList,
        GroupList,
        CurrentGroupHeader,
        CurrentProfile,
        AddGroup,
        ContextMenu,
        OriginMessage,
        PopUpReactions,
    },
    computed: {
      selectedConversation()
      {
        return this.myConversations && this.myConversations[this.selectedConversationIndexLocal]
        ? this.myConversations[this.selectedConversationIndexLocal] : null;
      },
      currentMessageText()
      {
          return this.currentMessage;
      },
      originMessageBGcolor(){
        if(this.originMessage.Sender == sharedData.UserSession.UserID)
          return "backgroundColor: var(--message-own)"
        else  
          return "backgroundColor: var(--message-other)"
      }
    }
  }
</script>

<template>
    <div id="body" class="Flexbox" style="width: calc(100vw - 6vh);">
        <div id="mainLeft">
            <div id="groupListParent" class="custom-scrollbar">
              <GroupList v-if="myConversations != null"
              :selectedConversationIndexLocal="selectedConversationIndexLocal"
              :conversations="myConversations"
              @SelectNewConversation="SelectNewConversationInApp"/>
            </div>

            <div id="AddNewGroup">
              <AddGroup
              @openOverlayInMode="openOverlayInMode"
              @logout="logout"
              />
            </div>
          
            <CurrentProfile
            :username="UserSession.Username"
            :profilePicture="UserSession.ProfilePicture"

            @openOverlayInMode="openOverlayInMode"
            />
        </div>

        <div id="main" style="position: relative;">

            <div id="messageListBackgroundImage">
            </div>

            <CurrentGroupHeader
            id="CurrentGroupHeader"
            ref="groupHeaderRef"
            :selectedConversation="this.selectedConversation"
            @openOverlayInGroupMode="openOverlayInGroupMode"
            />

            <div style="display:flex; margin-bottom:0px; margin-top:auto;">
            </div>
            
            <MessagesList
            id="MessagesList"
            ref="messagesList"
            v-if="myConversations != null && this.selectedConversation != null"
            :textMessages="this.selectedConversation.Messages"
            :convType="this.selectedConversation.Type"
            :refreshKey="this.selectedConversationIndexLocal"
            @refreshMessageList="refreshMessageList"
            @onPageRefresh="onPageRefresh"
            @openContextMenu="openContextMenu"
            @openReactionsMenu="openReactionsMenu"
            />

            <div id="BottomPartWrapper">

              <div id="spaceForBottom"
              v-if="myConversations != null"
              style="display: block; width: 100%; height: 0px; border-top: 2px solid rgb(206, 215, 240); margin-bottom: 5px;">
              </div>

              <div id="OriginMessage" ref="OriginMessage"
              :style=originMessageBGcolor
              style="display: flex; flex-direction: row;"
              v-if="this.originMessage != null"
              >

                <OriginMessage
                id="OriginMessage"
                style="flex-grow: 1; margin-left: 0px; margin-right: 0px; border: 0px; background-color: unset;"
                :convType="this.selectedConversation.Type"
                :message="this.originMessage"
                />

                <div style="display: flex; min-width: 55px; justify-content: center; align-items: center;"
                  @click="cancelReply"
                >
                  <img
                  src="https://cdn-icons-png.flaticon.com/128/190/190406.png"
                  style="max-width: 25px; max-height: 25px;"
                  />
                </div>
               

              </div>

              <div id="TextAndSend" v-if="myConversations != null">
                  
                <button id="sendButton" @click="openEmojis()" class="sendButtonImageContainer"> 
                  <img src="https://cdn-icons-png.flaticon.com/128/11202/11202612.png"/>
                </button>

                <input type="file" ref="fileInput" @change="uploadPicture" style="display: none;" />

                <emoji-picker id="EmojiPicker"
                  v-if="emojiPanelVisible"
                  @emoji-click="addEmojiToCurrentMessage"
                />

                <button id="sendButton" @click="openAttach()" class="sendButtonImageContainer"> 
                  <img src="https://cdn-icons-png.flaticon.com/128/4859/4859820.png"/>
                </button>

                <div style="display: block; position: relative; width: 100%; max-height: 100px;">
                  <textarea id="currentTextArea"
                  rows="1"
                  v-model="currentMessage"
                  placeholder="Type a message"
                  @keydown.enter="sendMessage"
                  @input="adjustHeight"
                  class="custom-scrollbar"
                  ></textarea>
                </div>

                <button id="sendButton" @click="sendMessage()" class="sendButtonImageContainer"> 
                  <img src="https://cdn-icons-png.flaticon.com/128/561/561226.png"
                  style="margin-right: 2.5px;"
                  />
                </button>

              </div>
            </div>

            <ContextMenu ref="contextMenu"
            v-show="contextMenuVisible"
            v-if="myConversations != null && this.selectedConversation != null"
            :conversationID="this.selectedConversation.Id"
            @click="closeContextMenu"
            @refreshLocalMessage="refreshLocalMessage"
            @setOriginMessage="setOriginMessage"
            @openOverlayInMode="openOverlayInMode"
            @openForwardOverlay="openForwardOverlay"
            />

            <PopUpReactions ref="reactionsMenu"
            v-show="popUpReactionsVisible"
            @click="closeContextMenu"
            />

        </div>

    </div>
</template>

<style>


#mainLeft {
  display: block; 
  margin-left: auto;
  margin-top: -105px;
  width: 20%;
  min-width: 125px;
  height: calc(90vh - 105px);
  z-index: 1;

}

#groupListParent {
  display: block; 
  margin-left: auto;
  padding: 5px;
  /*height: calc(70vh - 65px - 46px);*/
  height: 100%;
  /*
  border-radius: 15px;
  border-bottom-left-radius: 0;
  border-bottom-right-radius: 0;
  */
  background-color: var(--panel-bg);

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
  background-color: var(--panel-bg);
  /*height: calc(70vh - 6px);*/

  
  width: 80%;
  /*border-radius: 15px;*/

  border: 3px solid rgba(0, 0, 0, .25);
  height: calc(90vh);

  /*flex-grow: 1;*/
}

#CurrentGroupHeader {
  display: flex;
  height: 75px;
  min-height: 75px;
  max-height: 75px;

  margin-bottom: 5px;
  border-radius: 15px;
  
  background-color: var(--main-outline);
  z-index: 1;
}

#MessagesList {
  display: flex;
  flex-direction: column;
  
  margin-bottom: auto;
  overflow-y: auto;

  /*height: 50vh;*/

  /*height: calc(85vh - 150px);*/
  height: auto;
  max-height: calc(90vh - 150px);

  flex-grow: 1;
  flex-shrink: 1;
  flex-basis: 0;

  /*
  background-color: rgba(255, 0, 0, .25);
  flex-grow: 1;*/

}

#TextAndSend {
  display: flex;
  flex-direction: row;

  margin-top: 0px;
  margin-bottom: 0px;

  height: 60px;
  padding-top: 10px;

  gap:5px;
}

.Flexbox {
  display: flex;
  align-items: center;
  gap: 5px;
}


#currentTextArea {
  
  display: block;

  width: 100%;
  max-height: 100px;

  margin-top: 50px;

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
  margin-top: -410px;
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

  margin-bottom: -2px;
}


#BottomPartWrapper {

  display: flex;
  flex-direction: column;

  background-color: var(--panel-bg);
  z-index: 1;
}


</style>