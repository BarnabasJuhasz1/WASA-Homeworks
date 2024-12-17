
<script>

export default {
  props: ['picture', 'groupName', 'memberCount'],
  data() {
    return {
    };
  },

};
</script>


<template>
  
  <div id="HeaderParent">

		<div id=HeaderPicture class="image-container">
				<img :src="picture"/>
		</div>

    <div class="NameAndMessage">
    
      <div id="HeaderGroupName">
          {{ groupName }}
      </div>

      <div id="HeaderMemberCount">
          {{ memberCount }} members
      </div>

    </div>
    
    <div id="settingsDots" class="image-container">
        <img src="https://icon-library.com/images/android-three-dots-icon/android-three-dots-icon-0.jpg"/>
    </div>

	</div>

</template>


<style>

#HeaderParent {
  display: flex;
  flex-direction: row;
  
}


#HeaderPicture {
  height: 65px;
  width: 65px;

  min-width: 65px;
	min-height: 65px;
}

#HeaderGroupName {
  color: rgb(200, 200, 200);
	font-weight: bold;
  font-size: xx-large;

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


</style>


