
<script>
import { sharedData } from '../services/sharedData.js';
import axios from "../services/axios.js"
import ProfileObject from '../components/ProfileObject.vue';

export default 
{
    props: {
        selectedConversation:{
            type: Object,
            default: null,
        },
        messageToForward:{
            type: Object,
            default: null,
        },
        myConversations:{
            type: Object,
            default: null,
        }
    },
    data() {
        return {
            conversationProfiles: [],

            forwardToConversationsLocal: [],
        }
    },
    components:{
        ProfileObject,
    },
    mounted(){
        this.GetConversationNameAndPicture();
    },
    methods: {
        clickedCheckbox(index) {
            // console.log("clicked to: ", index);

            const userIndex = this.forwardToConversationsLocal.indexOf(index);
            
            if (userIndex === -1) {
                this.forwardToConversationsLocal.push(index);
            } else {
                this.forwardToConversationsLocal.splice(userIndex, 1);
            }

            // console.log("Updated forwardToUsers: ", this.forwardToConversationsLocal);
        },
        OnForwardButtonClicked(){
            // console.log("forward to: ", this.forwardToConversationsLocal);

            for(let i = 0; i < this.forwardToConversationsLocal.length; i++){

                this.ForwardRequest(this.myConversations[this.forwardToConversationsLocal[i]].Id);
            }
        },
        async ForwardRequest(forwardToConvID){
            // console.log("forwarding msg to conv: ", forwardToConvID, " to msg: ", this.messageToForward.Id)
            try{
                let response = await this.$axios.post(
                "/conversation/"+ this.selectedConversation.Id
                + "/message/"+ this.messageToForward.Id
                + "/comment", 
                {
                    ForwardToConvID: forwardToConvID
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
                // this.$emit('refreshLocalMessage', response.data)
                this.$emit('closeOverlay');

            } catch (error) {
                console.error("Error forwarding message! ", error);
                alert("Error forwarding message!")
            }

        },
        async GetConversationNameAndPicture(){
            console.log("Getting profiles!")
            this.conversationProfiles = []

            for(let i = 0; i < this.myConversations.length; i++){
                if(this.myConversations[i].Type == "GroupType"){
                    this.conversationProfiles.push({
                        Name: this.myConversations[i].GroupName,
                        Picture: this.myConversations[i].GroupPicture
                    })
                }
                else{
                    // if I am the first participant, return the other user
                    if(this.myConversations[i].Participants[0] == sharedData.UserSession.UserID){
                        let prof = await sharedData.getUserProfile(this.myConversations[i].Participants[1])
                        this.conversationProfiles.push({
                            Name: prof.Username,
                            Picture: prof.ProfilePicture 
                        })
                    }
                    else{
                        let prof = await sharedData.getUserProfile(this.myConversations[i].Participants[0])
                        this.conversationProfiles.push({
                            Name: prof.Username,
                            Picture: prof.ProfilePicture 
                        })
                    }
                }
            }
        },
    },
    computed: {
 
    },
}
</script>

<template>
    <div class="OverlayCenterCenter">
        <div class="OverlayCenterBlock" style="display: flex; min-height: 400px; height: auto; margin-bottom: 200px; flex-direction: column; align-items: center;">
            <div id="OverlayTop" class="SimpleContainer">
                Forward to:
            </div>
            <div id="OverlayBottom" class="SimpleContainer" style="display: flex; max-height: 600px;">
                
                <div style="display:block; overflow-y: auto; width: 300px" class="custom-scrollbar">
                    <div class="custom-scrollbar">
                        <div id="mainList" v-for="(conv, index) in myConversations" :key="index">
                        
                        <ProfileObject   
                            v-if="conversationProfiles[index] != null"   
                            :username="conversationProfiles[index].Name"
                            :profilePicture="conversationProfiles[index].Picture"
                            :editable="false"
                            :hasCheckBox="true"
                            @clickedCheckbox="clickedCheckbox(index)"
                        />

                        </div>
                        
                    </div>
                </div>

                <button id="CreateButton" style="margin-bottom: 5px; margin-top:5px"
                @click="OnForwardButtonClicked">
                    Forward
                </button>
         
            </div>
        </div>
    </div>
</template>

<style>


</style>
