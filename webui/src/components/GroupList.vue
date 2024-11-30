
<script>
import Group from './Group.vue';

export default {
  props: ['conversations'],
  emits: ['SelectNewConversation'],
  components: {
    Group,
  },
  data() {
    return {
    };
  },
  methods: {
    SelectNewConversation(index){
      console.log('Received notification at group list: ', index);
      /*this.selectedConversationIndex = index*/
      this.$emit("SelectNewConversation", index);
    },
  }

};
</script>


<template>
  <div>
    <div id="MainGroupList" v-for="(conversation, index) in conversations" :key="index">
      <Group
      :index="index"
      :picture="conversation.picture"
      :name="conversation.name"
      :last-message="conversation.messages[conversation.messages.length-1]"
      @SelectNewConversationAtGroupList="SelectNewConversation"/>
    </div>
  </div>

  <div id="emptySpace">
    
  </div>
  
  <div>
    test
  </div>

</template>


<style>

#MainGroupList {
  display: block; /* Enable Flexbox */
  /*justify-content: center; /* Center horizontally */
  border: 2px solid rgba(0, 0, 0, .25);
  background-color: rgba(0, 0, 0, .25);
  border-radius: 15px;
  margin-bottom: 5px;
  /*margin-left: 5px;
  margin-right: 5px;*/
}

</style>


