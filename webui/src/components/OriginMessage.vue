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
			
			this.username = userID == sharedData.UserSession.UserID ? "You" : profile.Username;

		},
		BackgroundColor() {
			return this.message.Sender == sharedData.UserSession.UserID ? 'var(--message-own)' : 'var(--message-other)';
		}
	},
	data() {
		return {
			username: null,
		}
	},

	mounted() {
		
		this.getProfile(this.message.Sender);
  	},
	watch: {
		message: {
			handler(newValue) {
				if (newValue.Sender !== this.username) {
					this.getProfile(newValue.Sender);
				}
			},
			deep: true,
		},
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
		isBase64Image() {
			return this.message.Content.startsWith("data:image/")
			&& this.message.Content.includes(";base64,");
		},
  	}
}
</script>


<template >

	<div id="MessageParent" style="text-align: left; border: 2px solid rgb(0, 0, 0, 0.5); white-space: nowrap; overflow-x: hidden; text-overflow: ellipsis; padding:0px;"
		:style="{ backgroundColor: BackgroundColor() }">
		
		<div id="ComplexMessageAndEmoji" style="align-items: flex-start; white-space: nowrap; overflow-x: hidden; overflow-y: hidden; text-overflow: ellipsis;">

			<div id="ComplexMessage" style="max-width: 100%;color: black; size: 0.9;">
				
				<div id="username">
					{{ this.username }}
				</div>

				<div class="message" style="white-space: nowrap; overflow-x: hidden; overflow-y: hidden; text-overflow: ellipsis;">

					<div v-if="!this.isBase64Image" style="max-width: 100px;">
							{{ this.message.Content }}
					</div>

					<div v-if="this.isBase64Image">
						<img :src="this.message.Content" style="width: 100%; height: 100%; object-fit: contain; max-width: 50px;"/>
					</div>

				</div>

			</div>

		</div>

	</div>
	

</template>


<style>




</style>
