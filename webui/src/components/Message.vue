<script>
import { ref } from 'vue'
import { sharedData } from '../services/sharedData.js';

export default {
	props: ['message', 'userID', 'convType', 'content', 'timestamp', 'msgStyle'],
	methods: {
		async getProfile(userID) {

			const profile = await sharedData.getUserProfile(userID);
			// console.log("for message content: ", this.content, " the found sender: ", profile.Username)
			this.username = profile.Username;
			this.profilePic = profile.ProfilePicture;
		},
		formattedTimestamp() {
			// Parse the input string
			const date = new Date(this.timestamp);

			// Extract hours and minutes
			const hours = String(date.getHours()).padStart(2, '0');
			const minutes = String(date.getMinutes()).padStart(2, '0');

			// Return the formatted time
			return `${hours}:${minutes}`;
		},
	},
	data() {
		return {
			username: null,
			profilePic: null
		}
	},

	mounted() {
		
		this.getProfile(this.userID);
  	},
	computed: {
		TimeStyle(){
			return this.msgStyle.wasSentByUser ? "margin-right: 15px;" : "margin-right: 15px;"
		},
		ExtraSpace(){
				return this.msgStyle.wasSentByUser ? "display: block; margin-right: 0; margin-left: auto;" : ""
		},
    	MessageStyle(){
				//'rgb(255, 221, 70)'
				//'rgb(255, 209, 0)'
        let bgColor = this.msgStyle.wasSentByUser ? 'var(--message-own)' : 'var(--message-other)'
		let alignment = this.msgStyle.wasSentByUser ? 'right' : 'left'
		let bottomPadding = this.message.EmojiReactions && this.message.EmojiReactions.length > 0 ? "padding-bottom: 25px;" :"padding-bottom: 5px;"

        return 'color:' + this.msgStyle.color +';' +
				' font-size:'+ this.msgStyle.fontSize + 'rem;' +
				" background-color:" + bgColor + ";" +
				" text-align:"+ alignment+";"+
				" padding-bottom:"+bottomPadding+";"+
				" font-weight: 400"
    	},
		formattedProfilePicture() {

			if(this.profilePic == null)
				return "";

			// const isValidUrl = this.profilePic.startsWith("http");

			// if (isValidUrl) {
			//	return this.profilePic; // Return URL directly if it's valid
			// } else {
				return `data:image/png;base64,${this.profilePic}`; // Return formatted Base64 string
			//}
    	},
		hasEmojiReactions(){
			return message.EmojiReactions.length > 0;
		}
  	}
}
</script>


<template>

	<div id="MessageParent" :style="{textAlign: this.msgStyle.wasSentByUser ? 'right' : 'left'}">

		<div v-if="!this.msgStyle.wasSentByUser
					&& this.profilePic != null
					&& this.convType != 'UserType'"
			id="profileImageOfOtherPerson {{content}}" class="image-container">

				<img :src="formattedProfilePicture"/>
		</div>
		
		<div id="ComplexMessageAndEmoji" :style="{alignItems: this.msgStyle.wasSentByUser ? 'flex-end' : 'flex-start'}">

			<div id="ComplexMessage" :style="MessageStyle">
				
				<div v-if="!this.msgStyle.wasSentByUser && this.convType != 'UserType'" id="username">
						{{ username }}
				</div>

				<div class="message" :style="TimeStyle">
						{{ content }}
				</div>

				<div class="time">
						{{ formattedTimestamp() }}
				</div>

			</div>

			<div id = "messageEmoji" :style="{backgroundColor: msgStyle.wasSentByUser ? 'var(--message-own)' : 'var(--message-other)'}"> 
				❤
			</div>
		</div>





	</div>
	

</template>


<style>

#MessageParent {
	display: flex;
}

#profileImageOfOtherPerson {
	visibility: visible;
}

.time {
	stroke-width: 50%;
	font-weight: 10;
	font-style: italic;
	margin-top: 5px;
	text-align: right;
}

#username {
	font-weight: bold;
	text-align: left;
	/*margin-left: 10px;*/
}

#ComplexMessageAndEmoji {
	display: flex;
	flex-grow: 1;

  	flex-direction: column;
}

#ComplexMessage{
	display: inline-block;
	
	/*border: 2px solid rgba(0, 0, 0, .25);
	background-color: rgba(0, 0, 0, .25);*/
	border-radius: 15px;
	padding-left: 8px;
	padding-right: 8px;
	padding-top: 5px;
	
	max-width: 75%;

}

.message {
	display: block; 

  	word-break: break-word; 
  	overflow-wrap: break-word;
}

#messageEmoji {
	display: block;
	
	width:25px;
	height:25px;

	margin-top: -15px;
	margin-left: 8px;
	margin-bottom: 5px;

	border-radius: 15px;
	
	align-content: center;
	text-align: center;
	
}

</style>
