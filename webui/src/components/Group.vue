<script>
import { sharedData } from '../views/sharedData.js';

export default {
	props: {
		conversation:{
			type: Object,
			required: true,
			default: null,
		},
		index:{
			required: true,
		},
		isSelected:{
			type: Boolean,
			required: true,
		},
	},
	emits: ['SelectNewConversationAtGroupList'],
	data() {
		return {
			headerName: null,
			profilePic: null,
			lastMessageSender: "",
		}
	},
	mounted() {
		
		this.getProfile();
		this.setLastMessageSender();
  	},
	methods: {
		SelectThisGroup()
		{
			// console.log('select group:', this.index);
			this.$emit('SelectNewConversationAtGroupList', this.index)
		},
		// GetHeaderName() {
      
		// 	if(this.conversation.Type == 'GroupType')
		// 		return conversation.GroupName

		// 	// return the name of the other person in the one-on-one conversation
		// 	for (let i = 0; i < this.conversation.Participants.length; i++) {
		// 		let participant = this.conversation.Participants[i];
		// 		if (participant !== sharedData.UserSession.Username) {
		// 			return participant;
		// 		}
		// 	}
		// },
		async getProfile()
		{
			if(this.conversation.Type == 'UserType')
			{
				for (let i = 0; i < this.conversation.Participants.length; i++)
				{
					let participant = this.conversation.Participants[i];
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
        		this.headerName = this.conversation.GroupName;
				this.profilePic = this.conversation.GroupPicture;
			}
		},
		async setLastMessageSender()
		{
			if (this.conversation.Messages == null || this.conversation.Messages.length == 0) {
				return "";
			}
			
			const sender = this.conversation.Messages[this.conversation.Messages.length-1].Sender;
			if(sender == sharedData.UserSession.UserID)
			{
				this.lastMessageSender = "You";
			}
			else
			{
				const senderProfile = await sharedData.getUserProfile(sender);
				this.lastMessageSender = senderProfile.Username;
			}
			//return senderProfile.Username;
		}
	},
	computed: {
		formattedProfilePicture() {
			return `data:image/png;base64,${this.profilePic}`;
    	},
		getLastMessageContent() {

			if (this.conversation.Messages == null || this.conversation.Messages.length == 0) {
				return "";
			}
		
			return (this.conversation.Messages[this.conversation.Messages.length-1]).Content;
		},
	
  	}
}
</script>


<template>

	<div id="Parent" @click="SelectThisGroup"
		style="border-radius: 12px;"
		:style="{outline: this.isSelected ? '2px solid orange': ''}">

		<div class="image-container">
			<img :src="formattedProfilePicture"/>
		</div>

		<div class="NameAndMessage"> 
			<div id="GroupName">
				{{ headerName }}
			</div>

			<div id="LastMessage" v-if="getLastMessageContent != ''">
				{{ this.lastMessageSender }}: {{ getLastMessageContent }}
			</div>
		</div>

	</div>
	
</template>


<style>

#Parent {
  	overflow: hidden;
  	text-overflow: ellipsis;
	display: flex;
}

.NameAndMessage {
	overflow: hidden;
  	text-overflow: ellipsis;
	display: block;
}

#GroupName {
	color: var(--font-light);
	font-weight: bold;
	padding-right: 5px;
	padding-top: 5px;

	white-space: nowrap;
	overflow: hidden;
	text-overflow: ellipsis; 
  	display: block;
}

#LastMessage {
	color: var(--font-light);
	padding-right: 5px;
	padding-top: 5px;

	white-space: nowrap;
	overflow: hidden; 
	text-overflow: ellipsis;
	display: block;
}

.image-container {
	padding: 5px;

	width: 50px;
	height: 50px;

	min-width: 50px;
	min-height: 50px;

	overflow: hidden;
	display: flex;
	justify-content: center;
	align-items: center;
}

.image-container img {
	border-radius: 100%;
	width: 100%; 
	height: 100%;
	object-fit: cover;
}

</style>
