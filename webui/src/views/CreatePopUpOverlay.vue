
<script>
import PopUpProfileHeader from '../components/PopUpProfileHeader.vue';
import ProfileObject from '../components/ProfileObject.vue';
import { sharedData } from '../services/sharedData';

export default 
{
    props: {
        profileText: {
            type: String,
            required: true,
        },
        profilePicture: {
            type: String,
            required: true,
            default: "https://cdn-icons-png.flaticon.com/128/14721/14721998.png",
        },
        conversationID: {
        }
    },
    setup(props) {
        
        // console.log("Overlay Mode:", props.overlayMode);

        let titleText = "";
        let buttonText = "";

        titleText = "Create New Conversation";
        // console.log("Title Text:", titleText);
        // console.log("Button Text:", buttonText);

        return {
            titleText,
        };
    },
    data() {
        return {
            // always populated for every PopUp
            currentProfileText: null,
            currentProfilePicture: null,
            // only populated for createConversation:
            currentParticipants: [],
            currentConversationType: "UserType",

            currentUsernameToAddText: "",
            //this should rather be the first participant in participants maybe ?
            invitedFriendText: null,
        }
    },
    components:{
        ProfileObject,
        PopUpProfileHeader,
    },
    methods: {
        handleUsernameUpdate(value) {
            this.currentProfileText = value;
        },
        handleProfilePictureUpdate(value) {
            this.currentProfilePicture = value;
        },
        selectNewConversationType(value) {
            this.currentConversationType = value;
        },
        universalButtonClicked() {
            // compare old profile pic and new profile pic
            // compare old username and new username

            this.CreateConversation()
        
        },
        //attemps to add the user to the conversation
        async AddUserToConversation() {

            // make http request to make sure user to add exist in database
            try {
                const response = await this.$axios.get('http://localhost:3000/user/fromName', {
                    params: { Username: this.currentUsernameToAddText },
                    headers: {
                    Authorization: 'Bearer ' + sharedData.UserSession.SessionToken,
                    'Content-Type': 'application/json',
                    },
                });

                // return response.data.ProfilePicture;
                console.log("added:", response.data)
                console.log("User added to the conversation: ", response.data.Id, ":", this.currentUsernameToAddText)
                this.currentParticipants.push(response.data.Id)

            } catch (e) {
                console.error('User not found!', e);
                alert('User not found!');
            }

        },
        async CreateConversation() {
            
            if((this.currentParticipants).length == 0){
                return
            }
            
            try {
                let response = await this.$axios.post(
                "http://localhost:3000/create/conversation", 
                // JSON body:
                {
                    ConversationType: this.currentConversationType, // either UserType or 
                    Participants: this.currentParticipants, // list of usernames
                },
                // Headers:
                {
                    headers: {
                        "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                        "Content-Type": "application/json",
                    },
                }
                );

                console.log(response.data);
                
                // notify parent view about the new conversation

                // sharedData.UserSession.Username = response.data.User.Username;
                // sharedData.UserSession.ProfilePicture = response.data.User.ProfilePicture;
                // sharedData.UserSession.SessionToken = response.data.SessionToken;
   
                // sharedData.UserSession = {
                //     Username: response.data.User.Username,
                //     ProfilePicture: response.data.User.ProfilePicture,
                //     SessionToken: response.data.SessionToken,
                // }

            } catch (e) {
                console.error(e.toString());
                
                alert("Creating conversation attempt failed!")

                // reset textArea input
                this.currentGroupNameText = "";
            }
        },
        selectedConversationFontStyle(myType) {
            const style = {
                fontSize: this.currentConversationType === myType ? 'large' : 'larger',
                fontWeight: this.currentConversationType === myType ? 'bold' : 'normal',
                color: this.currentConversationType === myType ? 'rgb(255, 209, 0)' : 'white',
            };
            console.log('Style:', style); // Debugging output
            return style;
        },
    },
    computed: {
        
    }
}
</script>

<template>
    <div class="OverlayCenterCenter">
        <div class="OverlayCenterBlock" style="margin-bottom: 200px;">
            <div id="OverlayTop" class="SimpleContainer">
                Create New Conversation
            </div>
            <div id="OverlayBottom" class="SimpleContainer" style="height: 400px">
                <div id="ConversationTypeSelector">
                    <div class="container" @click="currentConversationType='UserType'" style="width: 100%; line-height: 25px; max-height:50px">
                        <div class="simpleCenteredText" :style="selectedConversationFontStyle('UserType')">
                            Single
                        </div>
                    </div>
                    <div class="container" @click="currentConversationType='GroupType'" style="width: 100%; line-height: 25px; max-height:50px">
                        <div class="simpleCenteredText" :style="selectedConversationFontStyle('GroupType')">
                            Group
                        </div>
                    </div>
                </div>
                <div id="GroupSpecific">
                    <PopUpProfileHeader
                    :username="profileText"
                    :profilePicture="profilePicture"
                    @updateUsername="handleUsernameUpdate" 
                    @updateProfilePicture="handleProfilePictureUpdate"
                    />
                </div>
                
                
                <input
                type="text"
                class="basicTextField"
                v-model="currentUsernameToAddText"
                placeholder="Invite Username"
                @keydown.enter="AddUserToConversation"/>

                <button id="CreateButton" @click="universalButtonClicked"> Create </button>
            </div>
        </div>
    </div>
</template>

<style>


</style>
