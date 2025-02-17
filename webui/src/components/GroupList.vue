
<script>
import Group from './Group.vue';

export default {
  props:{
    conversations: {
      type: Object,
      required: true,
    },
    selectedConversationIndexLocal: {
      type: Number,
      required: true,
    },
  },
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
      // console.log('select conversation at group list: ', index);
      this.$emit("SelectNewConversation", index);
    },
    GetLastMessage(conversation){

      if(conversation.messages == null || conversation.messages.length == 0) {
        return null;
      }

      return conversation.messages[conversation.messages.length-1];
    }
  }

};
</script>


<template>
  <div>
    <div>
      <div id="MainGroupList" v-for="(conversation, index) in conversations" :key="`${conversation.GroupName}-${conversation.GroupPicture}-${conversation.Messages}-${index}`">
        
        <Group
          :conversation="conversation"
          :index="index"
          :isSelected="this.selectedConversationIndexLocal == index"
          @SelectNewConversationAtGroupList="SelectNewConversation"
        />

      </div>
    </div>

    <div id="emptySpace">
      
    </div>
  </div>
  
</template>


<style>

#MainGroupList {
  display: block;
  background-color: rgba(0, 0, 0, .25);
  
  border-radius: 5px;
  
  margin-bottom: 5px;
  /*margin-left: 5px;
  margin-right: 5px;*/
}

</style>


