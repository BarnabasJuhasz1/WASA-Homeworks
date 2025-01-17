
<script>
import { sharedData } from '../views/sharedData.js';

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
        this.getProfile();
      },
    },
  },
  data() {
		return {
      headerName: null,
			profilePic: null
		}
	},
	mounted() {
    // this.$nextTick(() => {
    //   this.getProfile();
    // }); 
  },
  methods: {
    async getProfile()
		{
      // console.log("getting header: ", this.selectedConversation)
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
			}
		},
  },
	computed: {
		formattedProfilePicture() {
      return `data:image/png;base64,${this.profilePic}`; // Return formatted Base64 string
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
          {{ this.selectedConversation.Participants.length }} members
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


