<script>
import PopUpOverlay from "./PopUpOverlay.vue";
import CreatePopUpOverlay from "./CreatePopUpOverlay.vue";
import axios from "../services/axios.js"
import {ref} from "vue";
import { sharedData } from '../services/sharedData.js';
import EditGroupPopUp from "./EditGroupPopUp.vue";
import NewConversationsView from "./NewConversationsView.vue";
import WasaTextUsersOverlay from "./WasaTextUsersOverlay.vue";
import ForwardOverlay from "./ForwardOverlay.vue";

export default {
  components: {
    NewConversationsView,
    PopUpOverlay,
    CreatePopUpOverlay,
    EditGroupPopUp,
    WasaTextUsersOverlay,
    ForwardOverlay,
  },
  setup() {
    const myFetchedConversations = ref(null);

    // console.log("TOKEN: ", sharedData.UserSession.SessionToken)
    // Define a method
    const GetMyConversations = async () => {
      try {
        let response = await axios.get(
          "/user/myConversations", 
          // Headers:
          {
            headers: {
            "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
            },
          }
        );

        console.log("myconversations: ", response.data);
        myFetchedConversations.value = response.data;
        
        //this.myFetchedConversations2 = response.data;
      }
      catch (error) {
        console.error("Error fetching conversations! ", error);
        alert("Error fetching conversations!")
      }
    };

    // Call the method
    GetMyConversations();

    return {
      myFetchedConversations,
      GetMyConversations,
    };
  },
  data() {
    return {
      showOverlay: false,
      overlayMode: "",
      overlayProfileText: "",
      overlayProfilePicture: "",
      inspectingConversation: null,

      messageToForward: null,
      // selectedConversationIndexLocal: null,
    };
  },
  methods: {
    openOverlayInMode(mode, profileText, profilePicture) {
      // console.log("grandparent user: ", profileText)
      // set the overlay mode enum ["USER", "GROUP", "CREATE_CONVERSATION"]
      this.overlayMode = mode;
      this.overlayProfileText = profileText;
      this.overlayProfilePicture = profilePicture;
      // enable the overlay
      this.showOverlay = true;
    },
    openForwardOverlay(conversation, message){
      this.openOverlayInMode("FORWARD", null, null);
      this.messageToForward = message
      this.inspectingConversation = conversation
    },
    openOverlayInGroupMode(conversation, selectedConvIndexLocal)
    {
      this.openOverlayInMode("GROUP", conversation.GroupName, conversation.GroupPicture);
      this.inspectingConversation = conversation
      // this.selectedConversationIndexLocal = selectedConvIndexLocal
    },
    async closeOverlay() {
      // disable the overlay
      this.showOverlay = false;

      const convCount = this.myFetchedConversations ? this.myFetchedConversations.length : 0;

      await this.GetMyConversations();

      if(this.myFetchedConversations == null)
        return;
      // if the overlay is closed because a new conversation was added,
      // we switch to that conversation
      if(convCount != this.myFetchedConversations.length)
      {
        const conversationsView = this.$refs.ConversationsViewRef;
        conversationsView.SelectNewConversationInApp(this.myFetchedConversations.length-1);
      }
      else // otherwise we update the currently selected one
      {
        const conversationsView = this.$refs.ConversationsViewRef;
        //conversationsView.SelectNewConversationInApp(this.selectedConversationIndexLocal);
        conversationsView.SelectNewConversationInApp(this.$route.params.id);
      }
    },
    handleUpdateConversation(index, newConversation) {
      this.myFetchedConversations[index] = newConversation;
    },
    async textPerson(profile){
      
      // console.log("Want to text person: ", profile)

      await this.GetMyConversations();

      // if we already have a conversation with that person,
      // then we just switch to that conversation
      if(this.myFetchedConversations != null)
      {
        for(let i = 0; i < this.myFetchedConversations.length; i++)
        {
          if(this.myFetchedConversations[i].Type == "UserType"){
            if((this.myFetchedConversations[i].Participants[0] == sharedData.UserSession.UserID &&
            this.myFetchedConversations[i].Participants[1] == profile.Id) || 
            (this.myFetchedConversations[i].Participants[1] == sharedData.UserSession.UserID &&
            this.myFetchedConversations[i].Participants[0] == profile.Id) )
            {
              this.showOverlay = false;
              const conversationsView = this.$refs.ConversationsViewRef;
              conversationsView.SelectNewConversationInApp(i);
              return;
            }
          }
        }
      }
      
      this.CreateConversation(profile.Id)

    },
    async CreateConversation(profileID) {

      try {
          let response = await this.$axios.post(
          "/create/conversation", 
          // JSON body:
          {
              ConversationType: "UserType",
              Participants: [profileID],
              ConversationName: "",
              ConversationPicture: "",
          },
          // Headers:
          {
              headers: {
                  "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                  "Content-Type": "application/json",
              },
          }
          );
          this.closeOverlay();

      } catch (e) {
          console.error(e.toString());
          
          alert("Texting user attempt failed!")
      }
    }
  },
};
</script>

<template >
  <div style="height: 90vh">

    <div class="background" style="height: 100%;">
      <NewConversationsView ref="ConversationsViewRef"
      @openOverlayInMode="openOverlayInMode"
      @openOverlayInGroupMode="openOverlayInGroupMode"
      @openForwardOverlay="openForwardOverlay"
      @update-conversation="handleUpdateConversation"
      :myConversations="myFetchedConversations"/>
    </div>

    <div v-if="showOverlay" class="overlay">
    
      <div class="blurEffect" @click="closeOverlay()"></div>

      <PopUpOverlay v-if="overlayMode=='USER'"
        :overlayMode="overlayMode"
        :profileText="overlayProfileText"
        :profilePicture="overlayProfilePicture"
        conversationID="0"
        @closeOverlay="closeOverlay"
        style="z-index: 1001;"
      />

      <EditGroupPopUp v-if="overlayMode=='GROUP'"
        :overlayMode="overlayMode"
        :profileText="overlayProfileText"
        :profilePicture="overlayProfilePicture"
        :inspectingConversation="inspectingConversation"
        @closeOverlay="closeOverlay"
        style="z-index: 1001;"
      />

      <CreatePopUpOverlay v-if="overlayMode=='CREATE_CONVERSATION'"
        :profileText="overlayProfileText"
        :profilePicture="overlayProfilePicture"
        conversationID="0"
        @closeOverlay="closeOverlay"
        style="z-index: 1001;"
      />

      <WasaTextUsersOverlay v-if="overlayMode=='WASA_TEXT_USERS'"
        @closeOverlay="closeOverlay"
        @textPerson="textPerson"
        style="z-index: 1001;"
      />

      <ForwardOverlay v-if="overlayMode=='FORWARD'"
        :selectedConversation="inspectingConversation"
        :myConversations="myFetchedConversations"
        :messageToForward="this.messageToForward"
        @closeOverlay="closeOverlay"
        style="z-index: 1001;"
      />

    </div>
  </div>
</template>

<style>

.background {
  display: block;
  flex-direction: column;
  justify-content: center;
  align-items: center;
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

.blurEffect {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(8px);
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
</style>
