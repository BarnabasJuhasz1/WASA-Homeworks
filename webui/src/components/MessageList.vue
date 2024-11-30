
<script>
import Message from './Message.vue';

export default {
  props: ['textMessages'],
  components: {
    Message,
  },
  data() {
    return {
    };
  },
  methods: {
    computedStyle(sentByUser) {
      let color = 'black'
      let size = 0.9;

      return {
        color: color,
        fontSize: size,
        wasSentByUser: sentByUser,
      };
    },
  }

};
</script>


<template>
  <div>
    <div id="mainList" v-for="(message, index) in textMessages" :key="index">
      <Message
      :username="message.username"
      :profile-pic="message.profilePic"
      :content="message.content"
      :timestamp="message.timestamp"
      :msgStyle="computedStyle(message.sentByUser)" />
    </div>
    
  </div>
</template>


<style>

#mainList {
  display: block;
  /*justify-content: center;  */
}

</style>


