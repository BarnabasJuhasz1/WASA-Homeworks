<script>
import { ref } from 'vue'
import { sharedData } from '../services/sharedData.js';

export default {
	props: ['userID', 'convType', 'content', 'timestamp', 'msgStyle'],
	methods: {
		async getProfile(userID) {

			const profile = await sharedData.getUserProfile(userID);
			console.log("for message content: ", this.content, " the found sender: ", profile.Username)
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
		}
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

        return 'color:' + this.msgStyle.color +';' +
				' font-size:'+ this.msgStyle.fontSize + 'rem;' +
				" background-color:" + bgColor + ";" +
				" text-align:"+ alignment+";"+
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
  	}
}
</script>


<template>

	<div id="Parent">

		<div v-if="!this.msgStyle.wasSentByUser
					&& this.profilePic != null
					&& this.convType != 'UserType'"
			id="profileImageOfOtherPerson {{content}}" class="image-container">

				<img :src="formattedProfilePicture"/>
		</div>

		<div id= "extraSpace" :style="ExtraSpace"> </div>		

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

	</div>
	

</template>


<style>

#Parent {
	display: flex;
	overflow-wrap: break-word; /* Modern equivalent for word-wrap */

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

#ComplexMessage{
	display: inline-block; /* Makes the div size fit its content */
	/*border: 2px solid rgba(0, 0, 0, .25);
	background-color: rgba(0, 0, 0, .25);*/
	border-radius: 15px;
	padding-left: 8px;
	padding-right: 8px;
	padding-top: 5px;
	padding-bottom: 5px;
	margin-bottom: 5px;

	max-width: 75%;
}

.message {
	display: inline-block; /* Makes the div size fit its content */

  	word-break: break-word; /* Breaks long words */
  	overflow-wrap: break-word; /* Modern equivalent for word-wrap */

}
</style>
