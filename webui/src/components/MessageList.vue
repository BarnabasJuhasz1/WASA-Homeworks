
<script>
import Message from './Message.vue';

import { sharedData } from '../services/sharedData.js';

export default {
  props: ['refreshKey', 'textMessages', 'convType'],
  components: {
    Message,
  },
  mounted() {
    this.$emit("onPageRefresh");
    this.messages = this.textMessages;
  },
  watch: {
    textMessages: {
      handler(newValue, oldValue) {
        this.messages = newValue;
      },
      deep: true,
    }
  },
  data() {
    return {
      messages: []
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

      return this.messages[message.OriginMessageId]
    },

  }
};
</script>


<template>
  
  <div class="custom-scrollbar">

    <div id="mainList"
      v-for="(message, index) in messages"
      :key="`${messages}-${message.Sender}-${message.Content}-${message.Timestamp}-${message.EmojiReactions}-${message.ReadBy}-${convType == 'GroupType' && (index == 0 || messages[index-1].Sender != message.Sender)}`"
      >

      <Message

      :convType="this.convType"
      :message="message"
      :originMessage="getOriginMessage(message)"
      :msgStyle="computedStyle(message)" 
      :isProfileVisible="convType == 'GroupType' && (index == 0 || messages[index-1].Sender != message.Sender)"      

      @openContextMenu="this.$emit('openContextMenu', message.Id)"
      @openReactionsMenu="this.$emit('openReactionsMenu', message.Id)"
      />
    </div>

  </div>

</template>


<style>

#mainList {
  display: block; 
  padding-right: 5px;
}

#messageListBackgroundImage {
  position: absolute;
  inset: 0;

  width: 100%;
  height: 100%;
  
  mix-blend-mode: multiply;
  background-image: url('https://img.freepik.com/premium-photo/white-minimal-geometry-background_231311-1693.jpg?w=1380');
  background-repeat: repeat;
  background-position: center;
  background-size: auto;
  
  opacity: 0.75;
  z-index: 0;
  pointer-events: none;

  mask-image: linear-gradient(black, black);
  -webkit-mask-image: linear-gradient(black, black);
}


</style>


