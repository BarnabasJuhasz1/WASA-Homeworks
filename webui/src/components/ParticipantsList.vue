
<script>
import ProfileObject from './ProfileObject.vue';
import { sharedData } from '../services/sharedData.js';

export default {
  props: {
    participants: {
      type: Object,
      required: true,
    },
    buttonForEachElement: {
      type: Boolean,
      default: false,
    },
  },
  components: {
    ProfileObject,
  },
  data() {
    return {
      sharedData
    };
  },
  methods: {
    
  }
};
</script>


<template>
  <div class="custom-scrollbar">
    <div id="mainList" style="display: flex;"
    v-for="(profile, index) in participants" :key="index">
      <ProfileObject
      
        :username="profile.Username"
        :profilePicture="profile.ProfilePicture"
        :editable="false"
      />
      
      <button class="ButtonBase" style="margin-left: -55px;"
      v-if="buttonForEachElement && profile.Id != this.sharedData.UserSession.UserID"
      @click="this.$emit('textPerson', profile)">
          Text
      </button>

    </div>
  </div>
</template>



<style>
.ButtonBase {
    height: 30px;
    min-height: 30px;
    background-color: var(--buttonBase);
    border: 2px solid var(--button);
      
    border-radius: 10px;
    font-size: medium;
    font-weight: bold;
    
    outline: none;

    margin-top: auto;
    margin-bottom: auto;
    margin-left: auto;
    margin-right: auto;
}
</style>


