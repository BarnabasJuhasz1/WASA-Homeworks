
<script>
import { sharedData } from '../services/sharedData';

// import 'emoji-picker-element';

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


            messageID: null,
            sender: null,
        }
    },
    components:{
        
    },
    methods: {
        show(x, y, messageID, sender) {
            this.position = { x, y };
            this.visible = true;

            this.messageID = messageID;
            this.sender = sender
            // console.log("selected message ID: ", this.messageID, sender)
            
        },
        hide() {
            this.visible = false;
        },
        async addEmoji(emoji) {
            console.log("emoji added: ", emoji)
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
        selectMessageToReplyTo() {
            this.$emit("setOriginMessage", this.messageID)
        },
        async ReplyToMessage() {  
            try{
                let response = await this.$axios.put(
                "/user", 
                // JSON body:
                { NewUsername: newName },
                // Headers:
                {
                    headers: {
                    "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                    "Content-Type": "application/json",
                    },
                }
                );

                console.log(response.data)

            } catch (error) {
                console.error("Error sending message! ", error);
                alert("Error sending message!")
            }
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
                console.log("deleting message: ", this.messageID)
                console.log(response.data)
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
        }
    }
}
</script>

<template>
    <div :style="{ top: `${position.y}px`, left: `${position.x}px` }">
        <div id="ContextMenuParent" :style="{ top: `${position.y}px`, left: `${position.x}px` }">
            <div id="ContextMenu" :style="{ top: `${position.y}px`, left: `${position.x}px` }">
                
                <button class="ContextMenuButton"
                    @click="selectMessageToReplyTo">
                    Reply
                </button>

                <button class="ContextMenuButton"
                    @click="universalButtonClicked">
                    Forward
                </button>

                <button class="ContextMenuButton"
                    v-if="IamTheMessageSender"
                    @click="DeleteMessage">
                    Delete
                </button>

                <div id="EmojiRow">
                    <!-- <emoji-picker id="EmojiPicker"
                    v-if="visible"
                    @emoji-click="addEmoji"
                    /> -->

                    <span
                        v-for="emoji in mostUsedEmojis"
                        :key="emoji"
                        @click="addEmoji(emoji)"
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
