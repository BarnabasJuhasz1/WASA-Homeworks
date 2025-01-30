import { reactive } from 'vue';
import axios from './axios.js';

function loadUserSessionFromLocalStorage() {
  const saved = localStorage.getItem('UserSession');
  return saved ? JSON.parse(saved) :
  {
    UserID: null, 
    Username: null,
    ProfilePicture: null,
    SessionToken: null
  };
}

function saveUserSessionToLocalStorage(session) {
  localStorage.setItem('UserSession', JSON.stringify(session));
}

export const sharedData = reactive({
  UserSession: loadUserSessionFromLocalStorage(), // Load UserSession from localStorage

  userProfileCache: {}, // Transient cache, not persisted

  async fetchUserProfile(userID) {
    try {
      const response = await axios.get('/user', {
        params: { UserID: userID },
        headers: {
          Authorization: 'Bearer ' + this.UserSession.SessionToken,
          'Content-Type': 'application/json',
        },
      });

      // console.log("user with id: ", userID, "fetched: ", response.data)

      if(response.data.ProfilePicture == null){
        // console.log("has no prof pic: ", response.data)
        response.data.ProfilePicture = await getDefaultImage();
      }
      // console.log("USER FETCHED: ", response.data)
      return response.data;
      
    } catch (e) {
      console.error('Error fetching user:', e);
      // alert('Fetching other user attempt failed!');
    }
  },

  async getUserProfile(userID) {
    if (this.userProfileCache[userID]) {
      // console.log('Cache hit for', userID, " GOT: ", this.userProfileCache[userID]);
      return this.userProfileCache[userID];
    }

    // console.log('Cache miss for', userID);
    const userProfile = await this.fetchUserProfile(userID);
    if (userProfile) {
      this.userProfileCache[userID] = userProfile;
    }
    return userProfile;
  },
});

async function getDefaultImage() {
  const response = await fetch('/images/default/basic_user_picture.png');
  const blob = await response.blob();

  return new Promise((resolve) => {
      const reader = new FileReader();
      reader.onloadend = () => {
        const base64String = reader.result.split(',')[1];  
        resolve(base64String)
      };
      reader.readAsDataURL(blob);
  });
}

// Watch for changes to UserSession and persist them to localStorage
import { watch } from 'vue';

watch(
  () => sharedData.UserSession,
  (newSession) => {
    saveUserSessionToLocalStorage(newSession);
  },
  { deep: true }
);
