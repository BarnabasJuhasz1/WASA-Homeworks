
import { reactive } from 'vue';

export const sharedData = reactive({
  UserSession: {
    Username: null,
    ProfilePicture: null,
    SessionToken: null,
  },
});


// import Vuex from 'vuex';

// export default new Vuex.Store({
//   state: {
//     UserSession: {
//       Username: null,
//       ProfilePicture: null,
//       SessionToken: null,
//     },
//   },
//   mutations: {
//     setUserSession(state, payload) {
//       state.UserSession = payload;
//     },
//   },
//   actions: {
//     updateUserSession({ commit }, UserSession) {
//       commit('setUserSession', UserSession);
//     },
//   },
//   getters: {
//     getUserSession: (state) => state.UserSession,
//   },
// });