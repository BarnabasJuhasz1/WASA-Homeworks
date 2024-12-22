import { reactive } from 'vue';
import axios from "./axios.js"

export const sharedData = reactive({
  UserSession: {
    Username: null,
    ProfilePicture: null,
    SessionToken: null,
  },
  profilePictureCache: {}, // Cache inside sharedData

  // Method to fetch the profile picture
  async fetchProfilePicture(username) {
    try {

      let response = await axios.get("http://localhost:3000/user",
      {
        params: { Username: username },
        headers: {
          "Authorization": "Bearer " + sharedData.UserSession.SessionToken,
          "Content-Type": "application/json",
        },
      });

      return response.data.ProfilePicture;

    } catch (e) {
        console.error(e.toString());
        alert("Fetching other user attempt failed!")
    }
  },

  // Method to get the profile picture from cache or fetch it
  async getProfilePicture(username) {

    if (this.profilePictureCache[username]) {
      console.log('Cache hit for', username);
      return this.profilePictureCache[username];
    }

    console.log('Cache miss for', username);
    const profilePicture = await this.fetchProfilePicture(username);
    this.profilePictureCache[username] = profilePicture;
    return profilePicture;
  }
});
