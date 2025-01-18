
<script>
import ParticipantsList from '../components/ParticipantsList.vue';
import { sharedData } from '../services/sharedData.js';
import axios from "../services/axios.js"

export default 
{
    data() {
        return {
            wasaTextUserIDs: null,
            wasaTextUsers: [],
        }
    },
    components:{
        ParticipantsList,
    },
    mounted(){
        this.GetWasaTextUsers();
    },
    methods: {
        async FetchAllUserProfiles(){
            for(let i = 0; i < this.wasaTextUserIDs.length; i++){
                let userProfile = await sharedData.getUserProfile(this.wasaTextUserIDs[i])
                this.wasaTextUsers.push(userProfile)
            }
            // console.log("wasa text users: ", this.wasaTextUsers)
        },
        async GetWasaTextUsers(){
            try {
                let response = await axios.get(
                "/user/all",
                // Headers:
                {
                    headers: {
                        "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                        "Content-Type": "application/json",
                    },
                }
                );
                // console.log("all userIDs received: ", response.data)
                this.wasaTextUserIDs = response.data
                this.FetchAllUserProfiles()
            }
            catch (error) {
                console.error("Error getting conversation! ", error);
                alert("Error getting conversation!")
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
    },
    computed: {

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
                
                <div style="display:block; overflow-y: auto; width: 300px" class="custom-scrollbar">
                    <ParticipantsList
                        :participants="this.wasaTextUsers"
                    />
                </div>

                <button id="CreateButton" style="margin-bottom: 5px; margin-top:5px;"
                @click="this.$emit('closeOverlay')">
                    Close
                </button>
         
            </div>
        </div>
    </div>
</template>

<style>


</style>
