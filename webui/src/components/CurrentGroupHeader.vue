
<script>
import { sharedData } from '../services/sharedData.js';

export default {
  props: {
    selectedConversation: {
      type: Object,
      required: true,
      default: null,
    }
  },
  watch:{
    selectedConversation: {
      handler(newValue, oldValue) {
        if(newValue != oldValue){
          this.getProfile();
        }
      },
    },
    
  },
  data() {
		return {
      headerName: null,
			profilePic: null,

      groupParticipantNames: [],
		}
	},
	mounted() {

  },
  methods: {
    async getProfile()
		{
      // console.log("getting header for conv: ", this.selectedConversation)
      if(this.selectedConversation == null)
        return;

			if(this.selectedConversation.Type == 'UserType')
			{
				for (let i = 0; i < this.selectedConversation.Participants.length; i++)
				{
					let participant = this.selectedConversation.Participants[i];
					if (participant != sharedData.UserSession.UserID)
					{
						const profile = await sharedData.getUserProfile(participant);
						this.headerName = profile.Username;
            this.profilePic = profile.ProfilePicture;
					}
				}
			}
			else
			{
        this.headerName = this.selectedConversation.GroupName;
				this.profilePic = this.selectedConversation.GroupPicture;

        this.groupParticipantNames = await this.getParticipantNames();
			}
		},
    async getParticipantNames(){
    
      let participantNames = [];
      this.headerName = this.selectedConversation.GroupName;
      this.profilePic = this.selectedConversation.GroupPicture;
      for (let i = 0; i < this.selectedConversation.Participants.length; i++)
      {
        let participant = await sharedData.getUserProfile(this.selectedConversation.Participants[i]);
        participantNames.push(participant.Username)
      }
      this.groupParticipantNames = [];
      return participantNames;
    }
  },
	computed: {
		formattedProfilePicture() {
      return `data:image/png;base64,${this.profilePic}`;
    },
  }	
};
</script>


<template>
  
  <div id="HeaderParent">

		<div id=HeaderPicture class="image-container">
				<img
        v-if="this.selectedConversation != null"
        :src="formattedProfilePicture"/>
		</div>

    <div class="NameAndMessage">
    
      <div id="HeaderGroupName"
      v-if="this.selectedConversation != null"
      >
          {{ headerName }}
      </div>

      <div id="HeaderMemberCount"
      v-if="this.selectedConversation != null && 
      this.selectedConversation.Type != 'UserType'"
      >
          {{ this.selectedConversation.Participants.length }}
          members: 
          <span v-if="this.groupParticipantNames.length > 0">
            {{ this.groupParticipantNames.join(', ') }}
          </span>
      </div>

    </div>
    
    <div id="settingsDots" class="image-container"
      v-if="this.selectedConversation != null && 
      this.selectedConversation.Type != 'UserType'"
      @click="this.$emit('openOverlayInGroupMode', this.selectedConversation)">
        <img src="https://icon-library.com/images/android-three-dots-icon/android-three-dots-icon-0.jpg"/>
    </div>

	</div>

</template>


<style>

#HeaderParent {
  display: flex;
  flex-direction: row;
  
}


#HeaderPicture {
  height: 75px;
  width: 75px;

  min-width: 75px;
	min-height: 75px;
}

#HeaderGroupName {
  /*color: rgb(200, 200, 200);*/
  color: var(--font-light);
	font-weight: bold;
  font-size: xx-large;

  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  display: block;

  margin-top: -5px;

}

#HeaderMemberCount {
  color: var(--font-light);
  font-style: italic;
  font-size: larger;

  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis; 
  display: block;

  margin-top: -5px;

}


</style>


