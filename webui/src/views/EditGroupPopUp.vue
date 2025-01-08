
<script>
import ParticipantsList from '../components/ParticipantsList.vue';
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
        participants: 
        {
            type: Object,
            required: true,
            default: [],
        }
    },
    setup(props) {

    },
    data() {
        return {
            // always populated for every PopUp
            currentProfileText: this.profileText,
            currentProfilePicture: this.profilePicture,
            // only populated for createConversation:
            currentParticipantIDs: [],
            currentParticipants: [],

            currentUsernameToAddText: "",
        }
    },
    components:{
        ProfileObject,
        PopUpProfileHeader,
        ParticipantsList,
    },
    methods: {
        handleUsernameUpdate(value) {
            this.currentProfileText = value;
        },
        handleProfilePictureUpdate(value) {
            this.currentProfilePicture = value;
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
                const response = await this.$axios.get('/user/fromName', {
                    params: { Username: this.currentUsernameToAddText },
                    headers: {
                    Authorization: 'Bearer ' + sharedData.UserSession.SessionToken,
                    'Content-Type': 'application/json',
                    },
                });

                // return response.data.ProfilePicture;
                console.log("response:", response.data)
                console.log("User added to the conversation: ", response.data.Id, ":", this.currentUsernameToAddText)
                console.log("me:", sharedData.UserSession.UserID)

                this.currentUsernameToAddText = ""

                // check if you were trying to add yourself
                if(response.data.Id == sharedData.UserSession.UserID)
                {
                    console.error('User is already part of conversation!');
                    alert('User is already part of conversation!');
                    return;
                }
                for(let i = 0; i < this.currentParticipantIDs.length; i++)
                {
                    // check if you were trying to add someone already in group
                    if(this.currentParticipantIDs[i] == response.data.Id)
                    {
                        console.error('User is already part of conversation!');
                        alert('User is already part of conversation!');
                        return;
                    }
                }

                this.currentParticipantIDs.push(response.data.Id)
                this.currentParticipants.push(response.data)
             
                if(this.currentConversationType == "UserType"){
                    this.currentProfileText = response.data.Username
                    this.currentProfilePicture = response.data.ProfilePicture
                }

            } catch (e) {
                console.error('User not found!', e);
                alert('User not found!');
            }

        },
        async GetFormattedPicture(){
            let formattedProfilePic;
            try{
                // make sure profile picture is in base 64
                if(typeof this.currentProfilePicture === "string"
                    && this.currentProfilePicture.startsWith("https")) 
                {
                    const response = await fetch(this.currentProfilePicture);
                    const blob = await response.blob();

                    const reader = new FileReader();
                    formattedProfilePic = await new Promise((resolve, reject) => {
                        reader.onloadend = () => {
                            const base64 = reader.result;
                            resolve(base64.split(",")[1]);
                        };                        
                        reader.onerror = reject;
                        reader.readAsDataURL(blob);
                    });
                }
                else
                {
                    formattedProfilePic = this.currentProfilePicture;
                }
            } catch (e) {
                console.error(e.toString());
                alert("profile pic conversion attempt failed!")
            }
            return formattedProfilePic;
        },
        async CreateConversation() {
            
            if((this.currentParticipantIDs).length == 0){
                return
            }

            let formattedProfilePic = await this.GetFormattedPicture();
            console.log(this.currentProfileText)
            console.log(formattedProfilePic)
            try {
                let response = await this.$axios.post(
                "/create/conversation", 
                // JSON body:
                {
                    ConversationType: this.currentConversationType, // either UserType or GroupType
                    Participants: this.currentParticipantIDs, // list of usernames
                    ConversationName: this.currentProfileText,
                    ConversationPicture: formattedProfilePic,
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
                this.$emit('closeOverlay')

            } catch (e) {
                console.error(e.toString());
                
                alert("Creating conversation attempt failed!")

                // reset textArea input
                this.currentGroupNameText = "";
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
                Edit Group
            </div>
            <div id="OverlayBottom" class="SimpleContainer" style="display: flex; max-height: 600px;">
                <div id="GroupSpecific">
                    <PopUpProfileHeader
                    :username="currentProfileText"
                    :profilePicture="currentProfilePicture"
                    @updateUsername="handleUsernameUpdate" 
                    @updateProfilePicture="handleProfilePictureUpdate"
                    />
                </div>
                
                <div style="display:block; overflow-y: auto; width: 300px" class="custom-scrollbar">
                    <ParticipantsList
                        :participants="this.currentParticipants"
                    />
                </div>

                <div class="Flexbox">
                    <input
                        type="text"
                        class="basicTextField"
                        v-model="currentUsernameToAddText"
                        placeholder="Invite Username"
                        @keydown.enter="AddUserToConversation"
                    />
                </div>

                <button id="CreateButton" style="margin-bottom: 5px; margin-top:5px"
                @click="universalButtonClicked">
                    Edit
                </button>
         
            </div>
        </div>
    </div>
</template>

<style>


</style>
