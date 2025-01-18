
<script>
import { ref, watch } from "vue";

export default {
  props: {
    username: {
      type: String,
      required: true,
      default: "New Group",
    },
    profilePicture: {
      type: String,
      required: true,
      default: "https://cdn-icons-png.flaticon.com/128/14721/14721998.png",
    },
    editable: {
      type: Boolean,
      required: true,
      default: true,
    },
    size: {
      default: 1,
    },
    isInContainer: {
      default: false,
    },
    hasCheckBox: {
      default: false,
    }
  },
  setup(props) {
    // console.log("username is ", props.username);
    const editableUsername = ref(props.username);

     watch(editableUsername, (newValue) => {
      // console.log("updated value: ", newValue)
      emit("update:username", newValue);
    });

    // console.log("editable username is: ", editableUsername.value)
    // editableUsername = props.username;
    return { editableUsername };
  },
  data() {
    return {
      // editableUsername: username,
      checked: false,
    };
  },
  computed: {
    // Prepend the Base64 string with the required prefix
    formattedProfilePicture() {
      const isValidUrl = this.profilePicture.startsWith("http");

      if (isValidUrl) {
        return this.profilePicture; // Return URL directly if it's valid
      } else {
        return `data:image/png;base64,${this.profilePicture}`; // Return formatted Base64 string
      }
    },
    // editableUsername() {
    //   return props.username;
    // }
  }
};
</script>


<template>
  
  <div id="profileObjectParent"
  :class="isInContainer ? 'container' : 'transparentContainer'">

		<div id=ProfilePicture  class="image-container" :style="{ minHeight: `${50 * size}px`, minWidth: `${50 * size}px` }">
				<img :src="formattedProfilePicture"/>
		</div>

    <input v-if="editable"
          id="ProfileName"
          class="transparentContainer"
          type="text"
          v-model="editableUsername"
          placeholder="Enter Username"
          :style="{ height: `${50 * size}px`, fontSize: `${18 * size}px` }"
          />

    <div v-if="!editable" text id="ProfileName"
    :style="{ height: `${50 * size}px`, width: `${200 * size}px` }"
    style="display: flex; align-items: center;">
        {{ editableUsername }}
    </div>

    <input
    type="checkbox"
    v-if="hasCheckBox"
    v-model="checked"
    @click="this.$emit('clickedCheckbox')"
    style="display: flex; align-items: center; width: 25px; height: 25px; margin-right: 20px;"
    />

  
	</div>

</template>


<style>

#profileObjectParent{
  display: flex;
  height: 60px;
  width: calc(100% - 10px);
  margin-top: 5px;

  justify-content: center;
  align-items: center;

  padding-left: 25px;
}

.transparentContainer {
  display: flex;
  flex-direction: row;
  align-items: center;

  border: 0;
  outline: 0;
  background-color: rgba(0, 0, 0, 0);
}

.container {
  display: flex;
  flex-direction: row;
  align-items: center;
  border-radius: 5px;

  border: 3px solid rgba(0, 0, 0, .25);
  background-color: rgba(0, 0, 0, .25);
}

.roundedContainer {
  display: flex;
  flex-direction: row;
  align-items: center;
  border-radius: 10px;

  border: 3px solid rgba(0, 0, 0, .25);
  background-color: rgba(0, 0, 0, .25);
}

.simpleCenteredText {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-left: auto;
  margin-right: auto;

  color: var(--font-light);  
}

#ProfileName {
  color: var(--font-light);
	font-weight: bold;
  font-size: larger;
  max-height: 50px;

  white-space: nowrap;
  overflow: hidden; 
  text-overflow: ellipsis; 
  display: block; 
}


</style>


