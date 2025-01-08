
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
  data() {
		return {
      headerName: null,
			profilePic: null
		}
	},
	mounted() {
		this.getProfile();
  },
  methods: {
    // GetHeaderName() {
      
    //   if(this.selectedConversation.Type == 'GroupType')
    //     return this.selectedConversation.GroupName

    //   // return the name of the other person in the one-on-one conversation
    //   for (let i = 0; i < this.selectedConversation.Participants.length; i++) {
    //     let participant = this.selectedConversation.Participants[i];
    //     if (participant !== sharedData.UserSession.Username) {
    //         return participant;
    //     }
    //   }
    // },
    check(){
      console.log("header touched")
    },
    async getProfile()
		{
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
			// const isValidUrl = this.selectedConversation.GroupPicture.startsWith("http");

			// if (isValidUrl) {
			// 	return this.selectedConversation.GroupPicture; // Return URL directly if it's valid
			// } else {
			// 	return `data:image/png;base64,${this.selectedConversation.GroupPicture}`; // Return formatted Base64 string
			// }
      return `data:image/png;base64,${this.profilePic}`; // Return formatted Base64 string
    },
  }	
};
</script>


<template>
  
  <div id="HeaderParent">

		<div id=HeaderPicture class="image-container">
				<img :src="formattedProfilePicture"/>
		</div>

    <div class="NameAndMessage">
    
      <div id="HeaderGroupName">
          {{ headerName }}
      </div>

      <div id="HeaderMemberCount" v-if="this.selectedConversation.Type != 'UserType'">
          {{ this.selectedConversation.Participants.length }} members
      </div>

    </div>
    
    <div id="settingsDots" class="image-container"
      v-if="this.selectedConversation.Type != 'UserType'"
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
  height: 65px;
  width: 65px;

  min-width: 65px;
	min-height: 65px;
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
}

#HeaderMemberCount {
  color: var(--font-light);
  font-style: italic;
  font-size: larger;

  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis; 
  display: block;
}


</style>


