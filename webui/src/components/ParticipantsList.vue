
<script>
import ProfileObject from './ProfileObject.vue';

export default {
  props: ['participants'],
  components: {
    ProfileObject,
  },
  data() {
    return {
      
    };
  },
  methods: {

  }
};
</script>


<template>
  <div class="custom-scrollbar">
    <div id="mainList" v-for="(profile, index) in participants" :key="index">
      <ProfileObject
      
        :username="profile.Username"
        :profilePicture="profile.ProfilePicture"
        :editable="false"
      />
      
    </div>
    
  </div>
</template>


<style>

</style>


