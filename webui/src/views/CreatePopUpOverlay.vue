
<script>
import ParticipantsList from '../components/ParticipantsList.vue';
import PopUpProfileHeader from '../components/PopUpProfileHeader.vue';
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
            currentProfileText: "New Friend",
            currentProfilePicture: "https://cdn-icons-png.flaticon.com/128/668/668709.png",
            // only populated for createConversation:
            currentParticipantIDs: [],
            currentParticipants: [],
            currentConversationType: "UserType",

            currentUsernameToAddText: "",
            //this should rather be the first participant in participants maybe ?
            invitedFriendText: null,
        }
    },
    components:{
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
        selectNewConversationType(value) {
            this.currentParticipantIDs = [];
            this.currentParticipants = [];
            this.currentConversationType = value;
            this.currentUsernameToAddText = ""

            if(value == "UserType")
            {
                this.currentProfileText="New Friend";
                this.currentProfilePicture = "https://cdn-icons-png.flaticon.com/128/668/668709.png";
            }
            else{
                this.currentProfileText="New Group";
                this.currentProfilePicture = "https://cdn-icons-png.flaticon.com/128/14721/14721998.png";
            }
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
                // console.log("response:", response.data)
                // console.log("User added to the conversation: ", response.data.Id, ":", this.currentUsernameToAddText)
                // console.log("me:", sharedData.UserSession.UserID)

                this.currentUsernameToAddText = ""

                // check if you were trying to add yourself
                if(response.data.Id == sharedData.UserSession.UserID)
                {
                    console.error('You cannot add yourself!');
                    alert('You cannot add yourself!');
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

                // if profile pic is unset, grab default picture instead of null
                if(response.data.ProfilePicture == null){
                    const profile = await sharedData.getUserProfile(response.data.Id)
                    response.data.ProfilePicture = profile.ProfilePicture
                }

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
                console.error("profile picture conversion attempt failed: ", e.toString());
                // alert("profile pic conversion attempt failed!")
            }
            return formattedProfilePic;
        },
        async CreateConversation() {
            
            if((this.currentParticipantIDs).length == 0){
                return
            }

            let formattedProfilePic = await this.GetFormattedPicture();
            // console.log(this.currentProfileText)
            // console.log(formattedProfilePic)
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

                // console.log(response.data);
                this.$emit('closeOverlay')

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
                color: this.currentConversationType === myType ? 'rgb(255, 209, 0)' : 'var(--font-light)',
            };
            // console.log('Style:', style); // Debugging output
            return style;
        },
    },
    computed: {
        IsParticipantListVisible() {

            if(this.currentConversationType == "UserType")
                return false;

            return true;
        },
        IsInviteFieldVisible() {
            
            if(this.currentConversationType == "UserType"
            && this.currentParticipantIDs.length > 0)
                return false;

            return true;
        },
    },
}
</script>

<template>
    <div class="OverlayCenterCenter">
        <div class="OverlayCenterBlock" style="display: flex; min-height: 400px; height: auto; margin-bottom: 200px; flex-direction: column; align-items: center;">
            <div id="OverlayTop" class="SimpleContainer">
                Create New Conversation
            </div>
            <div id="OverlayBottom" class="SimpleContainer" style="display: flex; max-height: 600px;">
                <div id="ConversationTypeSelector" style="margin-top: 0px;">
                    <div class="container" @click="selectNewConversationType('UserType')" style="width: 100%; line-height: 25px; max-height:50px">
                        <div class="simpleCenteredText" :style="selectedConversationFontStyle('UserType')">
                            Single
                        </div>
                    </div>
                    <div class="container" @click="selectNewConversationType('GroupType')" style="width: 100%; line-height: 25px; max-height:50px">
                        <div class="simpleCenteredText" :style="selectedConversationFontStyle('GroupType')">
                            Group
                        </div>
                    </div>
                </div>
                <div id="GroupSpecific">
                    <PopUpProfileHeader
                    :username="currentProfileText"
                    :profilePicture="currentProfilePicture"
                    @updateUsername="handleUsernameUpdate" 
                    @updateProfilePicture="handleProfilePictureUpdate"
                    />
                </div>
                
                <div style="display:block; overflow-y: auto; width: 300px" class="custom-scrollbar">
                    <ParticipantsList v-if="IsParticipantListVisible"
                        :participants="this.currentParticipants"
                    />
                </div>

                <div class="Flexbox">
                    <input v-if="IsInviteFieldVisible"
                        type="text"
                        class="basicTextField"
                        v-model="currentUsernameToAddText"
                        placeholder="Invite Username"
                        @keydown.enter="AddUserToConversation"
                    />
                </div>

                <button id="CreateButton" style="margin-bottom: 5px; margin-top:5px"
                @click="universalButtonClicked">
                    Create
                </button>
         
            </div>
        </div>
    </div>
</template>

<style>


</style>
