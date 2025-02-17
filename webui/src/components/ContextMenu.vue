
<script>
import { sharedData } from '../services/sharedData';

export default 
{
    props: {
        conversationID: {
            type: Number,
            required: true,
        },
    },
    setup(props) {
        
        let showConversationType = false;

        return {
            showConversationType,
        };
    },
    data() {
        return {
            visible: false,
            position: { x: 0, y: 0},

            // mostUsedEmojis: ['😀', '😂', '😍', '😭', '😎', '😡', '👍', '👏', '🎉'],
            mostUsedEmojis: ['❤', '😂', '😭', '😡', '👍'],

            message: null,
            messageID: null,
            sender: null,
        }
    },
    components:{
        
    },
    methods: {
        show(x, y, message) {
            this.position = { x, y };
            this.visible = true;

            this.message = message;
            this.messageID = message.Id;
            this.sender = message.Sender
            // console.log("selected message ID: ", this.messageID, sender)
            
        },
        hide() {
            this.visible = false;
        },
        async addEmojiClicked(emoji) {

            // if there is no emoji on the message yet
            if(this.message.EmojiReactions == null || this.message.EmojiReactions.length == 0){
                this.CommentEmojiRequest(emoji);
                return;
            }

            // check if you have already the same emoji in this message
            // if you do, then remove the emoji
            for(let i=0; i < this.message.EmojiReactions.length; i++){
                
                if(this.message.EmojiReactions[i].UserWhoReacted == sharedData.UserSession.Username){
                    
                    if(this.message.EmojiReactions[i].Content == emoji){
                        // uncomment old emoji
                        // console.log("about to uncomment old emoji")
                        this.UnCommentEmojiRequest();
                        return;
                    }
                    else{
                        // update old emoji with new one
                        // console.log("about to comment new emoji")
                        this.CommentEmojiRequest(emoji);
                        return;
                    }
                }
            }

            // if you dont have any emoji on the message, make emoji comment
            this.CommentEmojiRequest(emoji);
            
        },
        async CommentEmojiRequest(emoji){
            try{
                let response = await this.$axios.put(
                "/conversation/"+ this.conversationID
                + "/message/"+ this.messageID
                + "/comment", 
                // JSON body:
                {
                    TypeOfReaction: "EmojiReaction",
                    ContentOfReaction: emoji
                },
                // Headers:
                {
                    headers: {
                    "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                    "Content-Type": "application/json",
                    },
                }
                );

                // console.log(response.data)
                // make sure conversation is reloaded
                this.$emit('refreshLocalMessage', response.data)

            } catch (error) {
                console.error("Error sending emoji! ", error);
                alert("Error sending emoji!")
            }
        },
        async UnCommentEmojiRequest(){
            try{
                let response = await this.$axios.delete(
                "/conversation/"+ this.conversationID
                + "/message/"+ this.messageID
                + "/comment", 
                // Headers:
                {
                    headers: {
                    "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                    "Content-Type": "application/json",
                    },
                }
                );

                // console.log(response.data)
                // make sure conversation is reloaded
                this.$emit('refreshLocalMessage', response.data)

            } catch (error) {
                console.error("Error unsending emoji! ", error);
                alert("Error unsending emoji!")
            }
        },
        selectMessageToReplyTo() {
            this.$emit("setOriginMessage", this.messageID)
        },
        DeleteEvent() {
            this.DeleteMessage();
        },
        async DeleteMessage() {  
            try{
                let response = await this.$axios.delete(
                "/conversation/"+ this.conversationID
                + "/message/"+ this.messageID, 
                // Headers:
                {
                    headers: {
                    "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                    "Content-Type": "application/json",
                    },
                }
                );
                // console.log("deleting message: ", this.messageID)
                // console.log(response.data)
                this.$emit('refreshLocalMessage', response.data)

            } catch (error) {
                console.error("Error deleting message! ", error);
                alert("Error deleting message!")
            }
        },

    },
    computed: {
        IamTheMessageSender(){
            return this.sender == sharedData.UserSession.UserID
        },
        ContextMenuStyle(){
            if(this.message == null)
                return "";

            if (this.message.Sender == sharedData.UserSession.UserID) {
                const viewportWidth = window.innerWidth;
                const viewportHeight = window.innerHeight;

                const isBottomHalf = this.position.y > viewportHeight / 2;

                const calculatedRight = viewportWidth - this.position.x -10; // Calculate distance from the right edge
                const verticalPosition = isBottomHalf
                    ? `bottom: ${viewportHeight - this.position.y - 50}px;`
                    : `top: ${this.position.y - 50}px;`;

                return `${verticalPosition} right: ${calculatedRight}px;`;
            } else {
                const viewportHeight = window.innerHeight;

                const isBottomHalf = this.position.y > viewportHeight / 2;

                const verticalPosition = isBottomHalf
                    ? `bottom: ${viewportHeight - this.position.y - 50}px;`
                    : `top: ${this.position.y - 50}px;`;

                return `${verticalPosition} left: ${this.position.x - 175}px;`;
            }
        }
    }
}
</script>

<template>
    <div :style="ContextMenuStyle" style="z-index: 2;">
        <div id="ContextMenuParent" :style="ContextMenuStyle">
            <div id="ContextMenu">
                
                <button class="ContextMenuButton"
                    @click="selectMessageToReplyTo">
                    Reply
                </button>

                <button class="ContextMenuButton"
                    @click="this.$emit('openForwardOverlay', this.message)">
                    Forward
                </button>

                <button class="ContextMenuButton"
                    v-if="IamTheMessageSender"
                    @click="DeleteMessage">
                    Delete
                </button>

                <div id="EmojiRow" v-if="!IamTheMessageSender">
                    <!-- <emoji-picker id="EmojiPicker"
                    v-if="visible"
                    @emoji-click="addEmoji"
                    /> -->

                    <span
                        v-for="emoji in mostUsedEmojis"
                        :key="emoji"
                        @click="addEmojiClicked(emoji)"
                        class="emoji"
                    >
                        {{ emoji }}
                    </span>

                </div>
            </div>
        </div>
    </div>
</template>

<style>

#ContextMenu {
    width: 195px;
    max-height: 210px;
}

.emoji {
  cursor: pointer;
  font-size: 25px;
}

#EmojiRow {
    width: 100%;
    height: 40px;
    display: flex;
    flex-direction: row;

    overflow: hidden;

    display: flex;
    gap: 5px;
    overflow-x: auto;
    background-color: transparent;
}

#ContextMenuParent {
    position: absolute;
    background: rgb(0, 0, 0, 0.5);
    z-index: 1000;
    border-radius: 5px;
}

.ContextMenuButton {
    width: 100%;
    height: 50px;
    background-color: var(--background-light);

    border-radius: 5px;
    color: var(--font-light);
    font-size: medium;
    font-weight: bold;
    
    outline: none;
    border: 2px solid black;
    appearance: none;
    box-shadow: none;
}

</style>
