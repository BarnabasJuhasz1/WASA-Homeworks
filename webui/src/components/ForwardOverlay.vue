
<script>
import { sharedData } from '../services/sharedData.js';
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
            conversationsProfilesSuperset: [],

            conversationProfiles: [],

            forwardToConversationsLocal: [],
        }
    },
    components:{
        ProfileObject,
    },
    mounted(){
        this.SetConversationsProfilesSuperset();
    },
    methods: {
        async GetAllUsers(){
            try {
                let response = await this.$axios.get(
                "/user/all?search=",
                // Headers:
                {
                    headers: {
                        "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                        "Content-Type": "application/json",
                    },
                }
                );

                return response.data;
            }
            catch (error) {
                console.error("Error getting wasaText users for forwarding! ", error);
                alert("Error getting wasaText users for forwarding!")
            }
        },
        async SetConversationsProfilesSuperset(){
            this.conversationsProfilesSuperset = [];

            // get the "profiles" of the conversations we already have
            let myConversationProfiles = await this.GetConversationNameAndPicture();
            this.conversationsProfilesSuperset = myConversationProfiles;
            
            let allUsers = await this.GetAllUsers();
            
            if(allUsers == null)
                return;

            // loop through all the users
            for(let i = 0; i < allUsers.length; i++){
                
                // make sure you cannot forward to yourself
                if(allUsers[i] != sharedData.UserSession.UserID){
                    // get user profile
                    let userProfile = await sharedData.getUserProfile(allUsers[i])
                    // loop through all the conversations we already have in the superset
                    for(let j = 0; j < this.conversationsProfilesSuperset.length; j++){

                        // if user is already in my conversations, dont add it to the list again
                        if(userProfile.Username == this.conversationsProfilesSuperset[j].Name)
                        {
                            break;
                        }
                        else if (j == this.conversationsProfilesSuperset.length-1){
                            //otherwise, if we checked all conversations, add the user
                            this.conversationsProfilesSuperset.push({
                                Name: userProfile.Username,
                                Picture: userProfile.ProfilePicture,
                                Id: userProfile.Id,
                            })
                        }
                    }
                }
            }
        },
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
        async OnForwardButtonClicked(){
            // console.log("forward to: ", this.forwardToConversationsLocal);

            for(let i = 0; i < this.forwardToConversationsLocal.length; i++){

                let clickedIndex = this.forwardToConversationsLocal[i];
                // we have to check if the conversation already exists or not
                
                // if yes, forward message
                if(clickedIndex < this.myConversations.length)
                {
                    this.ForwardRequest(this.myConversations[clickedIndex].Id);
                }
                // if not, first create a new one-on-one conversation with the recipient
                // and then forward message
                else
                {
                    let conv = await this.CreateOneOnOneConversation([this.conversationsProfilesSuperset[clickedIndex].Id])
                    this.ForwardRequest(conv.Id);
                }
            }

            this.$emit('closeOverlay');
            
        },
        async CreateOneOnOneConversation(otherUser) {

            try {
                let response = await this.$axios.post(
                "/create/conversation", 
                // JSON body:
                {
                    ConversationType: "UserType",
                    Participants: otherUser,
                    ConversationName: "",
                    ConversationPicture: "",
                },
                // Headers:
                {
                    headers: {
                        "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                        "Content-Type": "application/json",
                    },
                }
                );

                return response.data;

            } catch (e) {
                console.error(e.toString());
                
                alert("Creating conversation attempt failed!")
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

            } catch (error) {
                console.error("Error forwarding message! ", error);
                alert("Error forwarding message!")
            }

        },
        async GetConversationNameAndPicture(){
            let conversationProfiles = []
            for(let i = 0; i < this.myConversations.length; i++){
                if(this.myConversations[i].Type == "GroupType"){
                    conversationProfiles.push({
                        Name: this.myConversations[i].GroupName,
                        Picture: this.myConversations[i].GroupPicture
                    })
                }
                else{
                    // if I am the first participant, return the other user
                    if(this.myConversations[i].Participants[0] == sharedData.UserSession.UserID){
                        let prof = await sharedData.getUserProfile(this.myConversations[i].Participants[1])
                        conversationProfiles.push({
                            Name: prof.Username,
                            Picture: prof.ProfilePicture 
                        })
                    }
                    else{
                        let prof = await sharedData.getUserProfile(this.myConversations[i].Participants[0])
                        conversationProfiles.push({
                            Name: prof.Username,
                            Picture: prof.ProfilePicture 
                        })
                    }
                }
            }
            return conversationProfiles
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
                        <div id="mainList" v-for="(conv, index) in conversationsProfilesSuperset" :key="index">
                        
                        <ProfileObject   
                            v-if="conversationsProfilesSuperset[index] != null"   
                            :username="conversationsProfilesSuperset[index].Name"
                            :profilePicture="conversationsProfilesSuperset[index].Picture"
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
