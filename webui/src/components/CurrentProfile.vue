
<script>

export default {
  props: {
    username: {
      type: String,
      required: true,
    },
    profilePicture: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
    };
  },
  computed: {
    // add required prefix to string
    formattedProfilePicture() {
      if (this.profilePicture) {
        return `data:image/png;base64,${this.profilePicture}`;
      }
      return '';
    }
  }
};
</script>


<template>
  
  <div id="ProfileHeaderParent" >

		<div id=ProfilePicture class="image-container">
			<img :src="formattedProfilePicture"/>
		</div>

    <div id="MyUsername">
      {{ username }}
    </div>

    <div id="settingsDots" class="image-container" @click="this.$emit('openOverlayInMode', 'USER', username, this.profilePicture)">
      <img src="https://icon-library.com/images/android-three-dots-icon/android-three-dots-icon-0.jpg"/>
    </div>

	</div>

</template>


<style>

#ProfileHeaderParent {
  display: flex;
  flex-direction: row;
  align-items: center;
  border-radius: 15px;
  background-color: var(--panel-bg);
  height: 60px;
  margin-top: 5px;

  border: 3px solid rgba(0, 0, 0, .25);
}

#ProfilePicture {
  height: 50px;
  width: 50px;

  min-width: 50px;
	min-height: 50px;
}

#MyUsername {
  color: var(--font-light);
	font-weight: bold;
  font-size: larger;

  white-space: nowrap;
  overflow: hidden; 
  text-overflow: ellipsis; 
  display: block; 
}


#settingsDots {
    margin-left: auto;
    margin-right: -10px;
}

</style>


