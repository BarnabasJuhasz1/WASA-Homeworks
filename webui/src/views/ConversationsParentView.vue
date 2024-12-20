<script>
import ConversationsView from "./ConversationsView.vue";
import PopUpOverlay from "./PopUpOverlay.vue";

export default {
  components: {
    ConversationsView,
    PopUpOverlay,
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
      console.log("grandpa user: ", profileText)
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
        <ConversationsView @openOverlayInMode="openOverlayInMode"/>
      </div>
  
      <div v-if="showOverlay" class="overlay">
       
        <div class="blurEffect" @click="closeOverlay()"></div>

        <PopUpOverlay
          :overlayMode="overlayMode"
          :profileText="overlayProfileText"
          :profilePicture="overlayProfilePicture"
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
  