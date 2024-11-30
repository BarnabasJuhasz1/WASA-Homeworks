
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
	</div>

</template>


<style>

#HeaderParent {
  display: block;
  flex-direction: row;
  
}

#HeaderPicture {
  height: 65px;
  width: 65px;

  min-width: 65px;
	min-height: 65px;
}

#HeaderGroupName {
  color: rgb(146, 146, 146);
	font-weight: bold;
  font-size: xx-large;

  white-space: nowrap; /* Prevent text from wrapping */
  overflow: hidden; /* Hide overflowing content */
  text-overflow: ellipsis; /* Show ellipsis for hidden text */
  display: block; /* Ensure proper box behavior */
}

#HeaderMemberCount {
  color: rgb(146, 146, 146);
  font-style: italic;
  font-size: larger;

  white-space: nowrap; /* Prevent text from wrapping */
  overflow: hidden; /* Hide overflowing content */
  text-overflow: ellipsis; /* Show ellipsis for hidden text */
  display: block; /* Ensure proper box behavior */
}

#MainGroupList {
  display: block; /* Enable Flexbox */
  /*justify-content: center; /* Center horizontally */
  border: 2px solid rgba(0, 0, 0, .25);
  background-color: rgba(0, 0, 0, .25);
  border-radius: 15px;
  margin-bottom: 5px;
  /*margin-left: 5px;
  margin-right: 5px;*/
}

</style>


