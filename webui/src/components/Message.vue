<script>
import { sharedData } from '../views/sharedData.js';
import OriginMessage from './OriginMessage.vue';

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
		originMessage:{
			type: Object,
		},
		msgStyle: {
			required: true,
		},

	},
	methods: {
		async getProfile(userID) {

			const profile = await sharedData.getUserProfile(userID);
			// console.log("for message content: ", this.content, " the found sender: ", profile.Username)
			this.username = userID == sharedData.UserSession.UserID ? "You" : profile.Username;
			this.profilePic = profile.ProfilePicture;
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
		setCheckMarkImg() {
			if (this.currentCheckmarkState == "DoubleCheckmark") {
				this.currentCheckmarkImage = this.readCheckMark;
			} else if (this.currentCheckmarkState == "SingleCheckmark") {
				this.currentCheckmarkImage =  this.receivedCheckMark;
			} else {
				this.currentCheckmarkImage =  null;
			}
		},
	},
	data() {
		return {
			username: null,
			profilePic: null,

			receivedCheckMark: "https://cdn-icons-png.flaticon.com/128/447/447147.png",
			readCheckMark: "https://cdn-icons-png.flaticon.com/128/18604/18604719.png",
		
			currentCheckmarkState: this.message.Status,
			currentCheckmarkImage: null,
		}
	},
	components: {
		OriginMessage,
	},
	mounted() {
		
		this.getProfile(this.message.Sender);
		this.setCheckMarkImg();
  	},
	computed: {
		TimeStyle(){
			return this.msgStyle.wasSentByUser ? "margin-right: 15px" : "margin-right: 15px"
		},
		ExtraSpace(){
				return this.msgStyle.wasSentByUser ? "display: block; margin-right: 0; margin-left: auto;" : ""
		},
    	MessageStyle(){
				//'rgb(255, 221, 70)'
				//'rgb(255, 209, 0)'
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

			// const isValidUrl = this.profilePic.startsWith("http");

			// if (isValidUrl) {
			//	return this.profilePic; // Return URL directly if it's valid
			// } else {
				return `data:image/png;base64,${this.profilePic}`; // Return formatted Base64 string
			//}
    	},
		hasEmojiReactions(){
			return this.message.EmojiReactions != null && this.message.EmojiReactions.length > 0;
		},
		mostReactedEmojiOnMessage(){

			if(this.hasEmojiReactions == false){
				return{
					mostCommonReaction: "",
					count: 0
				}
			}

			let reactionCounts = new Map();

			for (let i = 0; i < this.message.EmojiReactions.length; i++) {

				let reaction = this.message.EmojiReactions[i].Content;

				if (reactionCounts.has(reaction)) {
					reactionCounts.set(reaction, reactionCounts.get(reaction) + 1);
				} else {
					reactionCounts.set(reaction, 1);
				}
			}

			let mostCommonReaction = null;
			let maxCount = 0;

			reactionCounts.forEach((count, reaction) => {
				if (count > maxCount) {
					maxCount = count;
					mostCommonReaction = reaction;
				}
			});

			return {
				mostCommonReaction: mostCommonReaction,
				count: maxCount
			}
		},
		reactedEmojis(){

			if(this.hasEmojiReactions == false){
				return{
					emojis: "",
					count: 0
				}
			}

			let reactionCounts = new Map();

			for (let i = 0; i < this.message.EmojiReactions.length; i++) {

				let reaction = this.message.EmojiReactions[i].Content;

				if (reactionCounts.has(reaction)) {
					reactionCounts.set(reaction, reactionCounts.get(reaction) + 1);
				} else {
					reactionCounts.set(reaction, 1);
				}
			}
			// console.log(reactionCounts,  this.message.EmojiReactions.length)
			return {
				emojis: reactionCounts,
				count: this.message.EmojiReactions.length
			}
		},
		isCheckMarkVisible() {
			return this.msgStyle.wasSentByUser
				&& this.currentCheckmarkImage != null
				&& !this.message.HasBeenDeleted
				&& !this.currentCheckmarkState == 0;
		},
  	}
}
</script>


<template>

	<div id="MessageParent"
		:style="{textAlign: this.msgStyle.wasSentByUser ? 'right' : 'left'}">

		<div v-if="!this.msgStyle.wasSentByUser
					&& this.profilePic != null
					&& this.convType != 'UserType'"
			id="profileImageOfOtherPerson {{content}}" class="image-container">

				<img :src="formattedProfilePicture"/>
		</div>
		
		<div id="ComplexMessageAndEmoji"
			:style="{alignItems: this.msgStyle.wasSentByUser ? 'flex-end' : 'flex-start'}">

			<div id="ComplexMessage"
				:style="MessageStyle"
				@contextmenu.prevent="this.message.HasBeenDeleted ? null : this.$emit('openContextMenu')"
				>

				<OriginMessage style="border-radius: 10px; max-width: 500px;"
					v-if="this.originMessage != null && !this.message.HasBeenDeleted"
					:convType="this.convType"
					:message="this.originMessage"
					>
				</OriginMessage>


				<div v-if="!this.msgStyle.wasSentByUser && this.convType != 'UserType'" id="username">
						{{ this.username }}
				</div>
				
				<div
				class="message"
				:style="{
					TimeStyle,
					fontStyle: this.message.HasBeenDeleted ? 'italic' : 'normal',
					fontSize: /^(\p{Emoji_Presentation}|\p{Extended_Pictographic})$/u.test(this.message.Content.trim()) ? '4rem' : '1rem'
				}">
				{{ this.message.Content }}
				</div>

				<div class="timeAndCheckmark">
					{{ formattedTimestamp() }}
					<img
					v-if="isCheckMarkVisible"
					style="width:15px; height:15px;"
					:src="this.currentCheckmarkImage"
					>
				</div>

			</div>

			<div id="EmojiAndCount"
			v-if="!this.message.HasBeenDeleted"
			v-show="this.hasEmojiReactions"
			:style="{
				backgroundColor: msgStyle.wasSentByUser ? 'var(--message-own)' : 'var(--message-other)',
				marginLeft: msgStyle.wasSentByUser ? '0px' : '10px',
				marginRight: msgStyle.wasSentByUser ? '10px' : '0px'
				}"
			>
		
				<!-- <div id = "messageEmoji">
					{{ mostReactedEmojiOnMessage.mostCommonReaction }} 
				</div>

				<div id = "emojiCount"
				v-show="this.mostReactedEmojiOnMessage.count > 1"
				>
					{{ mostReactedEmojiOnMessage.count }}
				</div> -->

				<div style="display: flex; margin-left: 10px; margin-right: -5px;">
					<div id = "messageEmoji"
					v-for="(value, key) in reactedEmojis.emojis" :key="key">
						{{ value[0] }}
					</div>
				</div>

				<div id = "emojiCount" v-show="reactedEmojis.count > 1">
					{{ reactedEmojis.count }}
				</div>
			</div>

		</div>

	</div>
	

</template>


<style>

#MessageParent {
	display: flex;

	/*
	padding-left: 20px;
	padding-right: 20px;
	*/

	padding-left: 5vw;
	padding-right: 5vw;
}

#profileImageOfOtherPerson {
	visibility: visible;
}

.timeAndCheckmark {
	stroke-width: 50%;
	font-weight: 10;
	font-style: italic;
	/*margin-top: 5px;*/
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
	padding-bottom: 5px;
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

#EmojiAndCount {
	display: flex;

	margin-top: -15px;

	border-radius: 15px;

	justify-content: center; /* Centers content horizontally */
	align-items: center;     /* Centers content vertically */

	padding-right: 5px;

}

#messageEmoji {
	display: flex;
	
	width:25px;
	height:25px;

	gap: -10px;
	margin-left: -10px;

	justify-content: center; /* Centers content horizontally */
	align-items: center;     /* Centers content vertically */
}

#emojiCount {
	display: flex;
	
	width:15px;
	height:15px;
	justify-content: center; /* Centers content horizontally */
	align-items: center;     /* Centers content vertically */

}

</style>
