
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
        inspectingConversation: 
        {
            type: Object,
            required: true,
        },
    },
    setup(props) {

    },
    data() {
        return {
            // always populated for every PopUp
            currentProfileText: this.profileText,
            currentProfilePicture: this.profilePicture,
            // only populated for createConversation:
        
            currentParticipantIDs: this.inspectingConversation.Participants,
            currentParticipants: [],

            newParticipantIDs: [],

            currentUsernameToAddText: "",
        }
    },
    mounted(){
        this.GetUsersOfConversation()
    },
    components:{
        PopUpProfileHeader,
        ParticipantsList,
    },
    methods: {
        async GetUsersOfConversation() {
            for (let i = 0; i < this.inspectingConversation.Participants.length; i++)
            {
                let userProfile = await sharedData.getUserProfile(this.inspectingConversation.Participants[i])
                this.currentParticipants.push(userProfile)
            }
        },
        handleUsernameUpdate(value) {
            this.currentProfileText = value;
        },
        handleProfilePictureUpdate(value) {
            this.currentProfilePicture = value;
        },
        editButtonClicked() {
            // compare old group name and new group name
            if(this.currentProfileText != this.inspectingConversation.GroupName)
            {
                this.SetConversationName()
            }

            // compare old group pic and new group pic
            if(this.currentProfilePicture != this.inspectingConversation.GroupPicture)
            {
                this.SetConversationPicture()
            }

            // check if there are any new participants added
            for (let i = 0; i < this.newParticipantIDs.length; i++)
            {
                this.AddUserRequest(this.newParticipantIDs[i])
            }
            
        },
        //attemps to add the user to the conversation
        async AddUserToConversationLocal() {

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

                // add the user to the current group editing UI
                this.currentParticipantIDs.push(response.data.Id)
                this.newParticipantIDs.push(response.data.Id)

                this.currentParticipants.push(response.data)
            
                // actually add the user to the conversation
                // this.inspectingConversation.Participants.push(response.data.Id)

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
        async SetConversationName() {
            try {
                let response = await this.$axios.put(
                "/conversation/"+this.inspectingConversation.Id, 
                // JSON body:
                {
                    GroupName: this.currentProfileText,
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
                // this.inspectingConversation.GroupName = this.currentProfileText;

                // this.$emit('updateGroup', this.inspectingConversation);
                this.$emit('closeOverlay');

            } catch (e) {
                console.error(e.toString());         
                alert("Updating conversation name failed!")
            }
        },
        async SetConversationPicture() {

            let formattedProfilePic = await this.GetFormattedPicture();

            try {
                let response = await this.$axios.put(
                "/conversation/"
                +this.inspectingConversation.Id
                +"/groupPicture", 
                // JSON body:
                {
                    GroupPicture: formattedProfilePic,
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
                // this.inspectingConversation.GroupPicture = formattedProfilePic;

                // this.$emit('updateGroup', this.inspectingConversation)
                this.$emit('closeOverlay')

            } catch (e) {
                console.error(e.toString());
                
                alert("Adding user to conversation failed!")
            }
        },
        async AddUserRequest(userID) {
            try {
                let response = await this.$axios.put(
                "/conversation/"
                +this.inspectingConversation.Id
                +"/add", 
                // JSON body:
                {
                    UserIDToAdd: userID,
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
                // this.inspectingConversation.GroupName = this.currentProfileText;

                // this.$emit('updateGroup', this.inspectingConversation);
                this.$emit('closeOverlay');

            } catch (e) {
                console.error(e.toString());         
                alert("Updating conversation name failed!")
            }
        },
        async LeaveFromConversation() {
            try {
                let response = await this.$axios.delete(
                "/conversation/"
                +this.inspectingConversation.Id, 
                // Headers:
                {
                    headers: {
                        "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                        "Content-Type": "application/json",
                    },
                }
                );

                console.log(response.data);
                // this.inspectingConversation.GroupName = this.currentProfileText;

                // this.$emit('updateGroup', this.inspectingConversation);
                this.$emit('closeOverlay');

            } catch (e) {
                console.error(e.toString());         
                alert("Leaving from conversation failed!")
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
                        @keydown.enter="AddUserToConversationLocal"
                    />
                </div>

                <div style="display: flex; gap: 10px;">
                    <button id="CreateButton" style="margin-bottom: 5px; margin-top:5px; width: 100px;"
                    @click="editButtonClicked">
                        Edit
                    </button>

                    <button id="CreateButton" style="margin-bottom: 5px; margin-top:5px; width: 100px; background-color: red;"
                    @click="LeaveFromConversation">
                        Leave
                    </button>

                </div>
     

                
         
            </div>
        </div>
    </div>
</template>

<style>


</style>
