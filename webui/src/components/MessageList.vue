
<script>
import Message from './Message.vue';

import { sharedData } from '../services/sharedData.js';

export default {
  props: ['refreshKey', 'textMessages', 'convType'],
  components: {
    Message,
  },
  mounted() {
    this.$.emit("onPageRefresh");
  },
  data() {
    return {
      
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
  }
};
</script>


<template>
  <div class="custom-scrollbar">
    <div id="mainList" v-for="(message) in textMessages" :key="message.Sender && message.Content && message.Timestamp">
      <Message
      
      :userID="message.Sender"
      :convType="this.convType"
      :content="message.Content"
      :timestamp="message.Timestamp"
      :msgStyle="computedStyle(message)" 
      
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


