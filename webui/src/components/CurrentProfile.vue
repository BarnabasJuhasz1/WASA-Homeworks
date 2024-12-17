
<script>

export default {
  props: ['Username', 'ProfilePicture'],
  data() {
    return {
    };
  },

};
</script>


<template>
  
  <div id="ProfileHeaderParent" >

		<div id=ProfilePicture class="image-container">
				<img :src="ProfilePicture"/>
		</div>

    <div id="MyUsername">
        {{ Username }}
    </div>
    
      <div id="settingsDots" class="image-container">
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
  background-color: rgba(0, 0, 0, .25);
  height: 60px;
  margin-top: 5px;

  border: 3px solid rgba(0, 0, 0, .25);
  background-color: rgba(0, 0, 0, .25);
}

#ProfilePicture {
  height: 50px;
  width: 50px;

  min-width: 50px;
	min-height: 50px;
}

#MyUsername {
  color: rgb(146, 146, 146);
	font-weight: bold;
  font-size: larger;

  white-space: nowrap;
  overflow: hidden; 
  text-overflow: ellipsis; 
  display: block; 
}

#HeaderMemberCount {
  color: rgb(146, 146, 146);
  font-style: italic;
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


