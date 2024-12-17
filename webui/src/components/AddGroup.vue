
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

		<div class="image-container2">
				<img src="https://cdn-icons-png.flaticon.com/512/107/107078.png"/>
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
  background-color: rgba(0, 0, 0, .5);
  height: 40px;

  border: 3px solid rgba(0, 0, 0, .25);;
  background-color: rgba(0, 0, 0, .25);
}

.image-container2 {
  padding: 5px;

  width: 30px;
  height: 30px;

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


