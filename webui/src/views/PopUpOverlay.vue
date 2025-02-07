
<script>
import PopUpProfileHeader from '../components/PopUpProfileHeader.vue';
import { sharedData } from '../services/sharedData';

export default 
{
    props: {
        overlayMode: {
            type: String,
            required: true,
            default: "USER",
        },
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
        let showConversationType = false;

        if(props.overlayMode == "USER") {
            titleText = "Edit User";
            buttonText = "Save";
            showConversationType = false;
        }
        else if (props.overlayMode == "GROUP") {
            titleText = "Edit Group";
            buttonText = "Save";
            showConversationType = false;
        }
        // else if (props.overlayMode === "CREATE_CONVERSATION") {
        //     titleText = "Create New Conversation";
        //     buttonText = "Create";
        //     showConversationType = true;
        // }
        else {
            alert("Unsupported operation!");
        }

        // console.log("Title Text:", titleText);
        // console.log("Button Text:", buttonText);

        return {
            titleText,
            buttonText,
            showConversationType,
        };
    },
    mounted(){
        this.$emit('stopRefreshing');
    },
    data() {
        return {
            // always populated for every PopUp
            currentProfileText: null,
            currentProfilePicture: null,

        }
    },
    components:{
        PopUpProfileHeader,
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
             
            if(this.overlayMode == "USER"){
                
                try {
                    if(this.currentProfileText != null)
                    {
                        this.SetNameOperation(this.currentProfileText)
                    }
                    if(this.currentProfilePicture != null)
                    {
                        this.SetPictureOperation(this.currentProfilePicture)
                    }
                    
                } catch (error) {
                    console.error("Error trying to edit user profile: ", error)
                    // alert("There was an issue with editing the user profile.")
                }
            } 
        },
        async SetNameOperation(newName) {

            try {
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

                // console.log("REQUESTED SET-USERNAME, RESPONSE: ", response.data);
                sharedData.UserSession.Username = response.data;

                this.$nextTick(()=>{
                    this.$emit('closeOverlay');
                });
                alert("User profile edited successfully.")

            } catch (error) {
                alert("Username is occupied by someone else!")
            }
        },
        async SetPictureOperation(newPicture) {
            try {
                let response = await this.$axios.put(
                "/user/profilePicture", 
                // JSON body:
                {   ProfilePicture: newPicture },
                // Headers:
                {
                    headers: {
                    "Authorization": "Bearer "+sharedData.UserSession.SessionToken,
                    "Content-Type": "application/json",
                    },
                }
                );  

                // console.log(response.data);
                sharedData.UserSession.ProfilePicture = response.data;

                this.$emit('closeOverlay');

                alert("User profile edited successfully.")

            }
            catch (error) {
                console.error("There was an error editing profile picture: ", error)
                alert("Error editing profile picture!")
            }
        }
            
    },
    computed: {

    }
}
</script>

<template>
    <div class="OverlayCenterCenter">
        <div class="OverlayCenterBlock">
            <div id="OverlayTop" class="SimpleContainer">
                {{ titleText }}
            </div>
            <div id="OverlayBottom" class="SimpleContainer">
                <div v-if="showConversationType" id="ConversationTypeSelector">
                    <div class="container" style="width: 100%; line-height: 25px; max-height:50px">
                        <div class="simpleCenteredText" style="font-size: large;">
                            Single
                        </div>
                    </div>
                    <div class="container" style="width: 100%; line-height: 25px; max-height:50px">
                        <div class="simpleCenteredText" style="font-size: large;">
                            Group
                        </div>
                    </div>
                </div>
                <div v-if="showConversationType" id="GroupSpecific">
                    <PopUpProfileHeader
                    :username="profileText"
                    :profilePicture="profilePicture"
                    @updateUsername="handleUsernameUpdate" 
                    @updateProfilePicture="handleProfilePictureUpdate"
                    />
                </div>

                <input v-if="showConversationType" type="text" class="basicTextField" v-model="currentLoginUsernameText" placeholder="Enter Username" @keydown.enter="LoginEvent"/>
                
                <div v-if="!showConversationType" id="GroupSpecific">
                    <PopUpProfileHeader
                    :username="profileText"
                    :profilePicture="profilePicture"
                    @updateUsername="handleUsernameUpdate" 
                    @updateProfilePicture="handleProfilePictureUpdate"
                    />
                </div>

                <button id="CreateButton" @click="universalButtonClicked"> {{ buttonText }} </button>
            </div>
        </div>
    </div>
</template>

<style>

.OverlayCenterCenter {
    display: block;

    margin: auto;
    gap: 5px;
}

.OverlayCenterBlock {
    display: block;
    margin-bottom: 25%;

    min-width: 300px;
    width: 40vw;
    max-width: 500px;
    
    height: 25vh;
}

.SimpleContainer {
    display: block;
    background-color: rgba(0, 0, 0, .25);
    border-radius: 15px;
}

#OverlayTop {
    margin-bottom: 5px;
    background-color: rgba(0, 0, 0, .5);
    border: 2px solid rgba(0, 0, 0, .25);
    color: var(--font-light);
    font-size: x-large;
    font-weight: bold;
    text-align: center;
    margin-left: auto;
    margin-right: auto;
}

#OverlayBottom {
    height: 100%;
    min-height: 250px;
    border: 2px solid rgba(0, 0, 0, .25);
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}


#ConversationTypeSelector {
    display: flex;
    width: calc(100% - 10px);
    margin-top: 5px;
    height: 50px;
    gap: 5px;
}

.ConversationType {
    
    display: flex;
    margin: auto;
    margin-top: 25px;
}

.basicTextField {
    width: calc(100% - 10px);
    /*height: 15px;*/
    height: 50px;


    margin-left: 5px;
    margin-right: 5px;
    margin-top: auto;

    resize: none;
    
    overflow: auto;

    padding-left: 15px;
    padding-right: 15px;
    padding-top: 15px;
    padding-bottom: 15px;
    font-size: larger;
    outline: 0;


    border-radius: 5px;
    border: 3px solid rgba(0, 0, 0, .25);

    /*color: var(--font-light);*/
    /*background-color: rgba(0, 0, 0, .25);*/
    color: black;
    background-color: var(--message-box);

    align-items: center;
}

.usernameTextField::placeholder {
    color: var(--font-light);
}



#CreateButton {
    width: 25%;
    height: 30px;
    min-width: 25%;
    min-height: 30px;
    background-color: rgb(255, 209, 0);

    border-radius: 10px;
    font-size: medium;
    font-weight: bold;
    
    border: 0;
    outline: none;

    margin-top: auto;
    margin-bottom: auto;
    margin-left: auto;
    margin-right: auto;
}

</style>
