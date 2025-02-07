
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
  },
  setup(props, { emit }) {

    const editableUsername = ref(props.username);
    const editableProfilePicture = ref(props.profilePicture);

    watch(() => props.username, (newValue) => {
      editableUsername.value = newValue;
    });

    watch(() => props.profilePicture, (newValue) => {
      editableProfilePicture.value = newValue;
    });

    watch(editableUsername, (newValue) => {
      emit("updateUsername", newValue)
    });

    watch(editableProfilePicture, (newValue) => {
      emit("updateProfilePicture", newValue)
    });

    // Return the reactive variable to use it in the template
    return { editableUsername, editableProfilePicture };
  },
  data() {
    return {
    };
  },
  methods: {
    // Trigger the file input click programmatically
    openFileExplorer() {
      this.$refs.fileInput.click();
    },
    
    // Handle file change
    handleFileChange(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();

        // Read file as binary string
        reader.readAsArrayBuffer(file);

        // When the file is loaded
        reader.onload = (e) => {
          const arrayBuffer = e.target.result; // This is an ArrayBuffer
          const base64String = btoa(
              new Uint8Array(arrayBuffer)
                  .reduce((data, byte) => data + String.fromCharCode(byte), "")
          );
          this.editableProfilePicture = base64String;
          // console.log("File uploaded as Base64:", this.editableProfilePicture);
        };

        reader.onerror = (e) => {
            console.error("Error reading file:", e);
            alert("Error reading file!")
        };
      }
    }
  },
  computed: {
    // Prepend the Base64 string with the required prefix
    formattedProfilePicture() {

      if (typeof this.editableProfilePicture === "string"
        && this.editableProfilePicture.startsWith("http")) 
      {
        return this.editableProfilePicture;
      }
      else
      {
        return `data:image/png;base64,${this.editableProfilePicture}`; // Return formatted Base64 string
      }
    },
  }
};
</script>


<template>
  
  <div id="popUpProfileHeaderContainer">

		<div id=ProfilePicture class="image-container"
        :style="{ minHeight: `${100}px`, minWidth: `${100}px` }"
        @click="openFileExplorer">

				<img :src="formattedProfilePicture"/>
		</div>

    <div style="display: block; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;">
      <!-- <input
            id="ProfileName"
            class="transparentContainer"
            style="font-size: 35px; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;"
            :value="editableUsername"
            @input="onInput"
            placeholder="Enter Username"
            /> -->
      <!-- <div id="ProfileName"
          class="transparentContainer"
          contenteditable="true"
          style="width: 200px; height: 30px; overflow: hidden; white-space: nowrap; text-overflow: ellipsis;">
        This is the initial text that will be displayed in the box.
      </div> -->
      <input 
      id="ProfileName"
      v-model="editableUsername"
      type="text"
      class="transparentContainer popUpHeaderProfileText"
      style="font-size: 30px;" 
      placeholder="New Username"
      />

      <input type="file" ref="fileInput" @change="handleFileChange" style="display: none;" />

    </div>
	</div>

</template>


<style>

#popUpProfileHeaderContainer {
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 150px;
  width: 300px;
  margin-top: 15px;
  margin-left: auto;
  margin-right: auto;

  border: 0;
  outline: 0;
}

.popUpHeaderProfileText {
  width:300px;
  text-align: center;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}


</style>


