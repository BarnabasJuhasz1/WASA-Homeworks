<script>
import { sharedData } from '../services/sharedData.js';
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
		isProfileVisible: {
			required: true,
		},
	},
	watch: {
		message: {
			handler(newValue, oldValue) {
				//this.message = newValue;
				//console.error("updated text msgs")

				setTimeout(() => {
					this.getProfile(this.message.Sender);
					this.setCheckMarkImg();
					this.setMessageReactionList();

				}, 100);
					
			},
			deep: true,
		},
	},
	methods: {
		async getProfile(userID) {

			try{
				const profile = await sharedData.getUserProfile(userID);
				this.username = userID == sharedData.UserSession.UserID ? "You" : profile.Username;
				this.profilePic = profile.ProfilePicture;

			}catch(error){
				console.log("Error getting user profile!")
			}
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
		async setMessageReactionList(){
			if(this.message.EmojiReactions == null)
				return;

			let newReactionList = []

			for(let i = 0; i < this.message.EmojiReactions.length; i++){
				
				let prof = await sharedData.getUserProfile(this.message.EmojiReactions[i].UserWhoReacted)
				if(prof == null)
				{
					break;
				}

				newReactionList.push(
				{
					Username: prof.Username,
					Content: this.message.EmojiReactions[i].Content
				})
			}

			this.messageReactionList = newReactionList;
		}
	},
	data() {
		return {
			username: null,
			profilePic: null,

			receivedCheckMark: "https://cdn-icons-png.flaticon.com/128/447/447147.png",
			readCheckMark: "https://cdn-icons-png.flaticon.com/128/18604/18604719.png",
		
			currentCheckmarkState: this.message.Status,
			currentCheckmarkImage: null,

			// new list of message reactions to make sure it works for evaluation
			messageReactionList: [],
		}
	},
	components: {
		OriginMessage,
	},
	mounted() {
		this.getProfile(this.message.Sender);
		this.setCheckMarkImg();
		this.setMessageReactionList();
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

			const isValidUrl = this.profilePic.startsWith("http");

			if (isValidUrl) {
				return this.profilePic;
			} else {
				return `data:image/png;base64,${this.profilePic}`;
			}
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
		isBase64Image() {
			return this.message.Content.startsWith("data:image/")
			&& this.message.Content.includes(";base64,");
		}
  	}
}
</script>


<template>

	<div id="MessageParent"
		:style="{textAlign: this.msgStyle.wasSentByUser ? 'right' : 'left',
			marginLeft: this.isProfileVisible ? '0px;' : '50px',
		}"
	>
		
		<div v-if="!this.msgStyle.wasSentByUser
					&& this.profilePic != null
					&& this.convType != 'UserType'
					&& this.isProfileVisible"
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

				
				<div
				style="font-style: italic; border-radius: 10px; max-width: 500px;"
				v-if="this.message.WasForwarded && !this.message.HasBeenDeleted">
					Forwarded
				</div>


				<div  id="username"
					v-if="!this.msgStyle.wasSentByUser
						&& this.convType != 'UserType'
						&& this.isProfileVisible"
				>
						{{ this.username }}
				</div>
				
				
				<div style="display: flex; gap: 20px;">
					<div
					class="message"
					style="width:100%;"
					:style="{
						TimeStyle,
						fontStyle: this.message.HasBeenDeleted ? 'italic' : 'normal',
						fontSize: /^(\p{Emoji_Presentation}|\p{Extended_Pictographic})$/u.test(this.message.Content.trim()) ? '4rem' : '1rem'
					}">
					
						<div v-if="!this.isBase64Image" style="max-width: 100%;">
							{{ this.message.Content }}
						</div>

						<div v-if="this.isBase64Image">
							<img :src="this.message.Content" style="width: 100%; height: 100%; object-fit: contain; max-width: 250px; max-height: 250px;"/>
						</div>

					</div>

					<div class="timeAndCheckmark">
						
						{{ formattedTimestamp() }}

						<img
						v-if="isCheckMarkVisible"
						style="width:15px; height:15px; margin-bottom:3px;"
						:src="this.currentCheckmarkImage"
						>
					</div>
				</div>
				<!-- <div class="timeAndCheckmark">
					{{ formattedTimestamp() }}
					<img
					v-if="isCheckMarkVisible"
					style="width:15px; height:15px;"
					:src="this.currentCheckmarkImage"
					>
				</div> -->

			</div>

			<div id="EmojiAndCount"
			v-if="!this.message.HasBeenDeleted"
			v-show="this.hasEmojiReactions"
			@contextmenu.prevent="this.$emit('openReactionsMenu')"
			:style="{
				backgroundColor: msgStyle.wasSentByUser ? 'var(--message-own)' : 'var(--message-other)',
				marginLeft: msgStyle.wasSentByUser ? '0px' : '10px',
				marginRight: msgStyle.wasSentByUser ? '10px' : '0px'
				}"
			style="z-index: 1;"
			>
				<div class="custom-scrollbar" style="display: flex; margin-top: 10px; margin-left: 10px; margin-right: -5px; width: 100%; max-height: 100px; overflow-x: none; overflow-y: auto; flex-direction: column;">

					<div
					v-for="(reaction, i) in this.messageReactionList" :key="`${i}-${reaction.Username}-${reaction.Content}`">
					{{ reaction.Username }} : {{ reaction.Content }} &nbsp; 
					</div>

				</div>
				
				<!-- 
				<div id = "messageEmoji"
					v-for="(value, key) in reactedEmojis.emojis" :key="key">
						{{ value[0] }}
				</div>
				<div id = "emojiCount" style="font-weight: 600;"
					v-show="reactedEmojis.count > 1">
					{{ reactedEmojis.count }}
				</div> -->
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

	z-index:1;
}

#profileImageOfOtherPerson {
	visibility: visible;
}

.timeAndCheckmark {
	display: flex;
	flex-direction: row;
	align-items: flex-end;
	gap: 2px;
	flex-grow: 1;

	stroke-width: 50%;
	font-weight: 10;
	font-style: italic;
	/*margin-top: 5px;*/
	text-align: right;

	margin-top:10px;
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
	z-index:1;
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
