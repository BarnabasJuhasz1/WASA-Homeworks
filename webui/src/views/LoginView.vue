
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

            if((this.currentLoginUsernameText).trim().length < 3 || 
                (this.currentLoginUsernameText).trim().length > 16)
            {
                this.currentLoginUsernameText = "";
                alert("Username is too short or too long!")
                return
            }

            try {
                let response = await this.$axios.post(
                "/session", 
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
                // console.log(response.data);

                // update user sesson
                sharedData.UserSession.SessionToken = response.data.SessionToken;
                sharedData.UserSession.UserID = response.data.User.Id;
                sharedData.UserSession.Username = response.data.User.Username;
                const prof = await sharedData.getUserProfile(response.data.User.Id)
                sharedData.UserSession.ProfilePicture = prof.ProfilePicture;

                // console.log("User session updated:", sharedData.UserSession.UserID , ":", sharedData.UserSession.Username);

                this.$router.push('/conversations');

            } catch (e) {
                console.error(e.toString());
                
                alert("Login attempt failed!")
                // reset textArea input
                this.currentLoginUsernameText = "";

                return e;
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
    background-color: var(--background-light);
    border-radius: 15px;
}

#LoginTop {
    height: 25%;
    margin-bottom: 5px;
    background-color: rgba(0, 0, 0, .25);
    border: 2px solid rgba(0, 0, 0, .25);
    color: var(--font-light);
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

    background-color: var(--background-light);
    outline: none;

    color: var(--font-light);
}

.usernameTextField::placeholder {
    color: var(--font-light);
}



#LoginButton {
    width: 25%;
    height: 30px;
    min-width: 25%;
    min-height: 30px;
    background-color: var(--button);

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
