
<script>
import { sharedData } from '../services/sharedData.js';

export default 
{
    data() {
        return{ 
            currentLoginUsernameText: "",
            loggedInUsername: "",
            UserSession: sharedData.UserSession,
        }
    },
    methods: {
        async LoginEvent() {
            
            if((this.currentLoginUsernameText).trim() == ""){
                this.currentLoginUsernameText = "";
                return
            }
            
            // wait for getting the username confirmation,
            // profile picture, and the IDs of the conversations
            // UserData = await this.GetUserData();
            try {
                let response = await this.$axios.post(
                "http://localhost:3000/session", 
                // JSON body:
                {   Username: this.currentLoginUsernameText },
                // Headers:
                {
                    headers: {
                    "Content-Type": "application/json",
                    },
                }
                );

                // response.data contains JSON
                console.log(response.data);

                sharedData.UserSession.UserID = response.data.User.Id;
                sharedData.UserSession.Username = response.data.User.Username;
                sharedData.UserSession.ProfilePicture = response.data.User.ProfilePicture;
                sharedData.UserSession.SessionToken = response.data.SessionToken;

                console.log("User session updated:", sharedData.UserSession.UserID , ":", sharedData.UserSession.Username);

                this.$router.push('/conversations');

            } catch (e) {
                console.error(e.toString());
                
                alert("Login attempt failed!")

                // reset textArea input
                this.currentLoginUsernameText = "";
            }

        },
    },
}
</script>

<template>
    <div class="CenterCenter">
        <div id="LoginCenterBlock" class="CenterBlock">
            <div id="LoginTop" class="SimpleContainer">
                Login
            </div>
            <div id="LoginBottom" class="SimpleContainer">
                <input type="text" class="usernameTextField" v-model="currentLoginUsernameText" placeholder="Enter Username" @keydown.enter="LoginEvent"/>
                <button id="LoginButton" @click="LoginEvent()"> Login </button>
            </div>
        </div>
    </div>
</template>

<style>

.CenterCenter {
    display: block;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 0;
	gap: 5px;
    margin-top: 15vh;
}

.CenterBlock {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 300px;
    height: 200px;
}

.SimpleContainer {
    display: block;
    background-color: rgba(0, 0, 0, .25);
    border-radius: 15px;
}

#LoginTop {
    height: 25%;
    margin-bottom: 5px;
    background-color: rgba(0, 0, 0, .5);
    border: 2px solid rgba(0, 0, 0, .25);
    color: rgb(146, 146, 146);
    font-size: xx-large;
    font-weight: bold;
    text-align: center;
    margin-left: auto;
    margin-right: auto;
}

#LoginBottom {
    height: 75%;
    border: 2px solid rgba(0, 0, 0, .25);
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}

.usernameTextField {
    width: calc(100% - 10px);
    /*height: 15px;*/
    height: calc(50px);


    margin-left: 5px;
    margin-right: 5px;
    margin-top: auto;

    resize: none;
    border-radius: 10px;
    border: 0;
    
    overflow: auto;

    padding-left: 15px;
    padding-right: 15px;
    padding-top: 15px;
    padding-bottom: 15px;
    font-size: larger;

    background-color: rgb(75, 75, 75);
    outline: none;

    color: rgb(255, 255, 255);
}

.usernameTextField::placeholder {
    color: rgb(200, 200, 200);
}



#LoginButton {
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
