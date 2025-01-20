import { reactive } from 'vue';

export const Formatting = reactive({

    async GetFormattedPicture(currentProfilePicture)
    {
        console.log("Getting formated prof pic for wasatext")
        if(currentProfilePicture == null)
            currentProfilePicture = "https://cdn-icons-png.flaticon.com/128/668/668709.png";

        let formattedProfilePic;
        try{
            // make sure profile picture is in base 64
            if(typeof currentProfilePicture === "string"
                && currentProfilePicture.startsWith("https")) 
            {
                const response = await fetch(currentProfilePicture);
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
                console.log("Getting formated prof pic for wasatext 4")
                formattedProfilePic = currentProfilePicture;
            }
        } catch (e) {
            console.error(e.toString());
            alert("profile pic conversion attempt failed!")
        }
        console.log("Getting formated prof pic for wasatext returning: ", formattedProfilePic)

        return formattedProfilePic;
    }

});