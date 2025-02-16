<script>
import { sharedData } from '../services/sharedData.js';

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
	watch:{
		conversation: {
			handler(newValue, oldValue) {
				if(newValue != oldValue){
					this.getProfile();
					this.setLastMessageSender();
				}
			},
		},
		
	},
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
						if(profile == null)
						{
							break;
						}
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
				if(senderProfile != null)
					this.lastMessageSender = senderProfile.Username;
			}
			//return senderProfile.Username;
		}
	},
	computed: {
		formattedProfilePicture() {
			return `data:image/png;base64,${this.profilePic}`;
    	},
		getLastMessage() {
			if (this.conversation.Messages == null || this.conversation.Messages.length == 0) {
				return null;
			}
			// console.log(this.conversation.Messages[this.conversation.Messages.length-1])
			return (this.conversation.Messages[this.conversation.Messages.length-1]);
		},
		getFormattedLastMessageTimestamp() {
			if (!this.getLastMessage)
				return null;

			const lastMessageTimestamp = new Date(this.getLastMessage.Timestamp);
			const now = new Date();

			// check if the dates are the same
			const isSameDay = 
				lastMessageTimestamp.getFullYear() === now.getFullYear() &&
				lastMessageTimestamp.getMonth() === now.getMonth() &&
				lastMessageTimestamp.getDate() === now.getDate();

			if (isSameDay) {
				// return hour and minute
				return lastMessageTimestamp.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
			} else {
				// return year month day
				return lastMessageTimestamp.toLocaleDateString([], { year: 'numeric', month: '2-digit', day: '2-digit' });
			}
		},
		isBase64Image() {

			return this.getLastMessage.Content.startsWith("data:image/")
			&& this.getLastMessage.Content.includes(";base64,");
		},
	
  	}
}
</script>


<template>

	<div id="Parent" @click="SelectThisGroup"
		style="border-radius: 5px; padding-bottom: 5px; padding-top: 5px;"
		:style="{outline: this.isSelected ? '2px solid rgb(242, 157, 0, 0.75)': ''}">

		<div class="image-container">
			<img :src="formattedProfilePicture"/>
		</div>

		<div class="NameAndMessage"> 
			<div id="GroupName">
				<div style="flex:1;">
					{{ headerName }}
				</div>
				<div
				style="text-align: right; width: 100px; font-weight: normal; padding-left:10px; padding-right:5px;"
				v-if="getLastMessage != null">

				{{ getFormattedLastMessageTimestamp }}
			
				</div>
			</div>

			<div id="LastMessage" v-if="getLastMessage != null">
				{{ this.lastMessageSender }}: &nbsp;
				
				<div v-if="!this.isBase64Image">
					{{ getLastMessage.Content }}
				</div>

				<div v-if="this.isBase64Image">
					<img :src="getLastMessage.Content" style="width: 100%; height: 100%; object-fit: contain; max-width: 24px; max-height: 24px;"/>
				</div>
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
	display: flex;
	flex-direction: column;

	justify-content: center;
	flex: 1;
}

#GroupName {
	color: var(--font-light);
	font-weight: bold;
	padding-right: 5px;
	white-space: nowrap;
	overflow: hidden;
	text-overflow: ellipsis; 
  	display: flex;
}

#LastMessage {
	color: var(--font-light);
	padding-right: 5px;

	white-space: nowrap;
	overflow: hidden; 
	text-overflow: ellipsis;
	display: flex;

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
