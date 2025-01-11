<script>
import { sharedData } from '../services/sharedData.js';

export default {
	props: {
		convType: {
			type: String,
			required: true,
		},
		message:{
			type: Object,
			required: true,
		},
	},
	methods: {
		async getProfile(userID) {

			const profile = await sharedData.getUserProfile(userID);
			this.username = profile.Username;
		},
		formattedTimestamp() {
			// Parse the input string
			const date = new Date(this.message.Timestamp);

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
		}
	},

	mounted() {
		
		this.getProfile(this.message.Sender);
  	},
	computed: {
    	MessageStyle(){

        let bgColor = this.msgStyle.wasSentByUser ? 'var(--message-own)' : 'var(--message-other)'
		let alignment = this.msgStyle.wasSentByUser ? 'right' : 'left'
		let bottomPadding = this.hasEmojiReactions ? "10px;" :"3px;"

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

			return `data:image/png;base64,${this.profilePic}`;
    	},
  	}
}
</script>


<template>

	<div id="MessageParent" style="text-align: left">
		
		<div id="ComplexMessageAndEmoji" style="align-items: flex-start; white-space: nowrap; overflow-x: hidden; overflow-y: hidden; text-overflow: ellipsis;">

			<div id="ComplexMessage" style="max-width: 100%;color: black; size: 0.9;">
				
				<div id="username">
					{{ this.username }}
				</div>

				<div class="message" style="white-space: nowrap; overflow-x: hidden; overflow-y: hidden; text-overflow: ellipsis;">
					{{ this.message.Content }}
				</div>

				<div class="time">
					{{ formattedTimestamp() }}
				</div>

			</div>

		</div>

	</div>
	

</template>


<style>




</style>
