
<script>
import { sharedData } from '../services/sharedData';
import ProfileObject from './ProfileObject.vue';

export default 
{
    props: {

    },
    setup() {

    },
    data() {
        return {
            visible: false,
            position: { x: 0, y: 0},

            message: null,
            messageID: null,
            sender: null,

            emojiReactions: [],
            reactingProfiles: [],
        }
    },
    components:{
        ProfileObject,
    },
    methods: {
        show(x, y, message) {

            this.position = { x, y };
            this.visible = true;

            this.message = message;
            this.messageID = message.Id;
            this.sender = message.Sender;
            // console.log("selected message ID: ", this.messageID, sender)
            this.FetchAllReactingUserProfiles()
        },
        hide() {
            this.visible = false;
        },
        async FetchAllReactingUserProfiles(){
            this.reactingProfiles = []
            for(let i = 0; i < this.message.EmojiReactions.length; i++){
                let userProfile = await sharedData.getUserProfile(this.message.EmojiReactions[i].UserWhoReacted)
                this.reactingProfiles.push(userProfile)
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
    <div :style="ContextMenuStyle">
        <div id="ContextMenuParent" :style="ContextMenuStyle">
            <div id="ContextMenu" class="custom-scrollbar" style="overflow-y: auto; max-height: 250px;">
                <div class="custom-scrollbar" v-if="this.message != null && reactingProfiles != null">
                    <div id="mainList" style="display: flex; align-items: center; justify-content: center; padding-right: 25px;"
                    v-for="(reaction, index) in message.EmojiReactions" :key="`${reaction.Content}-${reaction.UserWhoReacted}-${index}`">
                        
                        <ProfileObject

                            v-if="reactingProfiles[index] != null"
                            :username="reactingProfiles[index].Username"
                            :profilePicture="reactingProfiles[index].ProfilePicture"
                            :editable="false"
                        />

                        <div style="color: var(--font-light); font-size: x-large;" v-if="reactingProfiles[index] != null">
                            {{ reaction.Content }}
                        </div>

                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>

</style>
