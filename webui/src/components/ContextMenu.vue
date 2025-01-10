
<script>
import { sharedData } from '../services/sharedData';

import 'emoji-picker-element';

export default 
{
    props: {
        // messageID: {
        //     type: Number,
        //     required: true,
        // },
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
            mostUsedEmojis: ['❤', '😂', '😭', '😎', '😡', '👍', '👏', '🎉'],


            messageID: null,
        }
    },
    components:{
        
    },
    methods: {
        show(x, y, messageID) {
            this.position = { x, y };
            this.visible = true;

            this.messageID = messageID;
            console.log("sleected message ID: ", this.messageID)

        },
        hide() {
            this.visible = false;
        },
        addEmoji(emoji) {
            console.log("emoji added: ", emoji)
            
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

            } catch (error) {
                console.error("Error sending message! ", error);
                alert("Error sending message!")
            }
        },

    },
    computed: {

    }
}
</script>

<template>
    <div :style="{ top: `${position.y}px`, left: `${position.x}px` }">
        <div id="ContextMenuParent" :style="{ top: `${position.y}px`, left: `${position.x}px` }">
            <div id="ContextMenu" :style="{ top: `${position.y}px`, left: `${position.x}px` }">
                
                <button class="ContextMenuButton"
                    @click="universalButtonClicked">
                    Reply
                </button>

                <button class="ContextMenuButton"
                    @click="universalButtonClicked">
                    Forward
                </button>

                <button class="ContextMenuButton"
                    @click="universalButtonClicked">
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
    height: 210px;
}

.emoji {
  cursor: pointer;
  font-size: 25px;
}

#EmojiRow {
    width: 100%;
    height: 60px;
    display: flex;
    flex-direction: row;

    overflow: hidden;

    display: flex;
    gap: 5px;
    overflow-x: auto; /* Allow scrolling for longer lists */
}

#EmojiPicker {
    background: white;
    border: .5px solid white;
    border-radius: 5px;
}

#ContextMenuParent {
    position: absolute;
    background: white;
    border: .5px solid white;
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
    
    border: .5px solid white;
    outline: none;

}

</style>
