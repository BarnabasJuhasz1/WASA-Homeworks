
<script>
import Message from './Message.vue';
import ContextMenu from './ContextMenu.vue';

import { sharedData } from '../services/sharedData.js';

export default {
  props: ['refreshKey', 'textMessages', 'convType'],
  components: {
    Message,
    ContextMenu,
  },
  mounted() {
    this.$.emit("onPageRefresh");
  },
  data() {
    return {
      // contextMenuVisible: false,
    };
  },
  methods: {
    computedStyle(message) {
      let color = 'black'
      let size = 0.9;

      return {
        color: color,
        fontSize: size,
        wasSentByUser: message.Sender == sharedData.UserSession.UserID,
      };
    },
    getOriginMessage(message) {
      
      if(message.OriginMessageId == -1){
        return null
      }

      return this.textMessages[message.OriginMessageId]
    }
    // openContextMenu() {
    //   this.contextMenuVisible = true;
    //   console.log("contextmenu found? ", this.$refs.contextMenu)
    //   this.$refs.contextMenu.show(event.clientX, event.clientY);
    //   document.addEventListener('click', this.closeContextMenu);
    // },
    // closeContextMenu() {
    //   this.contextMenuVisible = false;
    //   this.$refs.contextMenu.hide();
    //   document.removeEventListener('click', this.closeContextMenu);
    // },
  }
};
</script>


<template>
  <div class="custom-scrollbar">
    <div id="mainList"
      v-for="(message) in textMessages"
      :key="`${message.Sender}-${message.Content}-${message.Timestamp}-${message.EmojiReactions}`"
      >

      <Message

      :convType="this.convType"
      :message="message"
      :originMessage="getOriginMessage(message)"
      :msgStyle="computedStyle(message)" 
      
      @openContextMenu="this.$emit('openContextMenu', message.Id)"
      />      
    </div>


  </div>
</template>


<style>

#mainList {
  display: block; 
  padding-right: 5px; /* Adds a margin-like effect to the scrollbar */
}

</style>


