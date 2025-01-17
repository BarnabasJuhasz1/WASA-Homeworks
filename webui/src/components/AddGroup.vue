
<script>

export default {
  data() {
    return {
    };
  },

};
</script>

<template>
  
  <div id="AddToGroupParent" >

		<div class="image-container2" @click="this.$emit('openOverlayInMode', 'CREATE_CONVERSATION')">
				<img src="https://cdn-icons-png.flaticon.com/128/7887/7887095.png"/>
		</div>

    <div class="image-container2" @click="this.$emit('openOverlayInMode', 'WASA_TEXT_USERS')">
				<img src="https://cdn-icons-png.flaticon.com/128/200/200941.png"/>
		</div>

	</div>

</template>


<style>

#AddToGroupParent {
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 15px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
  height: 40px;

  border: 3px solid rgba(0, 0, 0, .25);
  background-color: var(--panel-bg);

  gap: 25px;
}

.image-container2 {
  padding: 5px;

  width: 35px;
  height: 35px;

	min-width: 30px;
	min-height: 30px;

  border-radius: 50%;
  overflow: hidden;
  display: flex;
  justify-content: center;
  align-items: center;
}

.image-container2 img {
  width: 100%; 
  height: 100%;
  object-fit: cover;
}

</style>


