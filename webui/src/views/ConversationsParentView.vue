<script>
import ConversationsView from "./ConversationsView.vue";
import PopUpOverlay from "./PopUpOverlay.vue";
import CreatePopUpOverlay from "./CreatePopUpOverlay.vue";
import axios from "../services/axios.js"
import {ref} from "vue";
import { sharedData } from '../services/sharedData.js';

export default {
  components: {
    ConversationsView,
    PopUpOverlay,
    CreatePopUpOverlay,
  },
  setup() {
    const myFetchedConversations = ref(null);

    // console.log("TOKEN: ", sharedData.UserSession.SessionToken)
    // Define a method
    const GetMyConversations = async () => {
      try {
        let response = await axios.get(
          "http://localhost:3000/user/myConversations", 
          // Headers:
          {
            headers: {
            "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
            },
          }
        );

        // console.log("myconversations ", response.data);
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
      myFetchedConversations
    };
  },
  data() {
    return {
      showOverlay: false,
      overlayMode: "",
      overlayProfileText: "",
      overlayProfilePicture: "",
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
    closeOverlay() {
      // disable the overlay
      this.showOverlay = false;
    }
  },
};
</script>

<template>
    <div>

      <div class="background">
        <ConversationsView
        @openOverlayInMode="openOverlayInMode"
        :myConversations="myFetchedConversations"/>
      </div>
  
      <div v-if="showOverlay" class="overlay">
      
        <div class="blurEffect" @click="closeOverlay()"></div>

        <PopUpOverlay v-if="overlayMode!='CREATE_CONVERSATION'"
          :overlayMode="overlayMode"
          :profileText="overlayProfileText"
          :profilePicture="overlayProfilePicture"
          conversationID="0"
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


      </div>
    </div>
  </template>
  
  <style>

  .background {
    height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-color: #f0f8ff;
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
  