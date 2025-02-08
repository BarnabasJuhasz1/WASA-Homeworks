
<script>
import { sharedData } from '../services/sharedData.js';
import axios from "../services/axios.js"
import ProfileObject from '../components/ProfileObject.vue';

export default 
{
    data() {
        return {
            wasaTextUserIDs: null,
            wasaTextUsers: [],
            currentUsernameToSearch: "",

            selectedUsersToText: [],
        }
    },
    components:{
        ProfileObject,
    },
    mounted(){
        this.GetWasaTextUsers();
    },
    methods: {
        clickedCheckbox(index) {
            console.log("clicked to: ", index);

            const userIndex = this.selectedUsersToText.indexOf(index);
            
            if (userIndex === -1) {
                this.selectedUsersToText.push(index);
            } else {
                this.selectedUsersToText.splice(userIndex, 1);
            }
            console.log("Updated forwardToUsers: ", this.selectedUsersToText);

        },
        async FetchAllUserProfiles(){
            this.wasaTextUsers = [];
            this.selectedUsersToText = []

            for(let i = 0; i < this.wasaTextUserIDs.length; i++){
                let userProfile = await sharedData.getUserProfile(this.wasaTextUserIDs[i])
                this.wasaTextUsers.push(userProfile)
            }
            // console.log("wasa text users: ", this.wasaTextUsers)
        },
        async GetWasaTextUsers(){
            const searchQuery = this.currentUsernameToSearch != "" ? `${encodeURIComponent(this.currentUsernameToSearch)}` : "";

            try {
                let response = await axios.get(
                "/user/all?search="+searchQuery,
                // Headers:
                {
                    headers: {
                        "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                        "Content-Type": "application/json",
                    },
                }
                );
                // console.log("all userIDs received: ", response.data)
                this.wasaTextUsers = []
                this.selectedUsersToText = []
            
                if(response.data != null){
                    this.wasaTextUserIDs = response.data
                    await this.FetchAllUserProfiles()
                    // console.log("wasatext user count: ", this.wasaTextUserIDs.length)
                }
            }
            catch (error) {
                console.error("Error getting wasaText users! ", error);
                alert("Error getting wasaText users!")
            }
        },
        async GetFormattedPicture(){

            if(this.currentProfilePicture == null)
                return "";

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
            // console.log("Getting formated prof pic for wasatext returning: ", formattedProfilePic)

            return formattedProfilePic;
        },
        async onTextButtonClicked(){

            if(this.selectedUsersToText.length == 0){
                this.$emit('closeOverlay')
            }
            else if(this.selectedUsersToText.length == 1){
                let userProfile = await sharedData.getUserProfile(this.wasaTextUserIDs[this.selectedUsersToText[0]])
                if(userProfile.Id == sharedData.UserSession.UserID){
                    this.$emit('closeOverlay')
                    alert("You cannot text yourself.")
                    return;
                }
                this.$emit('textPerson', userProfile);
            }
            else{
                let participants = []
                for(let i = 0; i < this.selectedUsersToText.length; i++){
                    if(this.wasaTextUserIDs[this.selectedUsersToText[i]] != sharedData.UserSession.UserID){
                        participants.push(this.wasaTextUserIDs[this.selectedUsersToText[i]]);
                    }
                }
                this.$emit('textGroup', participants);
            }

        },
    },
    computed: {
        textButtonString(){
            if(this.selectedUsersToText.length == 0){
                return "Close"
            }
            else if(this.selectedUsersToText.length == 1){
                return "Text User"
            }
            else{
                return "Text Group"
            }
        }
    },
}
</script>

<template>
    <div class="OverlayCenterCenter">
        <div class="OverlayCenterBlock" style="display: flex; min-height: 400px; height: auto; margin-bottom: 200px; flex-direction: column; align-items: center;">
            <div id="OverlayTop" class="SimpleContainer">
                WasaText Users
            </div>

            <div id="OverlayBottom" class="SimpleContainer" style="display: flex; max-height: 600px;">
                <input
                    type="text"
                    class="basicTextField"
                    style="margin-top: 5px;"
                    v-model="currentUsernameToSearch"
                    placeholder="Search Username"
                    @input="GetWasaTextUsers"
                 />

                <div style="display:block; overflow-y: auto; width: 300px;" class="custom-scrollbar">
                    <!-- <ParticipantsList
                        :participants="this.wasaTextUsers"
                        :buttonForEachElement="true"
                        @textPerson="textPerson"
                    /> -->
                    <div v-for="(profile, index) in this.wasaTextUsers" :key="index">
                    
                        <ProfileObject
                            v-if="index < this.wasaTextUserIDs.length"
                            :username="profile.Username"
                            :profilePicture="profile.ProfilePicture"
                            :editable="false"
                            :hasCheckBox="true"
                            @clickedCheckbox="clickedCheckbox(index)"
                        />

                    </div>
                        

                </div>

                <button id="CreateButton" style="margin-bottom: 5px; margin-top:5px; width: 100px;"
                @click="onTextButtonClicked">
                    {{ textButtonString }}
                </button>
         
            </div>
        </div>
    </div>
</template>

<style>


</style>
