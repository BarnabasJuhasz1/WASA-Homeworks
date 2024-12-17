<script>
export default {
	props: ['index', 'picture', 'name', 'lastMessage'],
	emits: ['SelectNewConversationAtGroupList'],
	methods: {
		SelectThisGroup()
		{
			console.log('group should be selected:', this.index);
			this.$emit('SelectNewConversationAtGroupList', this.index)
		}
	},
	computed: {

  }
}
</script>


<template>

	<div id="Parent" @click="SelectThisGroup">

		<div class="image-container">
				<img :src="picture"/>
		</div>

		<div class="NameAndMessage"> 
			<div id="GroupName">
					{{ name }}
			</div>

			<div id="LastMessage">
					{{ lastMessage.username }}: {{ lastMessage.content }}
			</div>
		</div>

	</div>
	
</template>


<style>

#Parent {
  overflow: hidden;
  text-overflow: ellipsis;
	display: flex;
}

.NameAndMessage {
	overflow: hidden;
  text-overflow: ellipsis;
	display: block;
}

#GroupName {
	color: rgb(200, 200, 200);
	font-weight: bold;
	padding-right: 5px;
	padding-top: 5px;

	white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis; 
  display: block;
}

#LastMessage {
	color: rgb(146, 146, 146);
	padding-right: 5px;
	padding-top: 5px;

	white-space: nowrap;
  overflow: hidden; 
  text-overflow: ellipsis;
  display: block;
}

.image-container {
  padding: 5px;

  width: 50px;
  height: 50px;

	min-width: 50px;
	min-height: 50px;

  border-radius: 50%;
  overflow: hidden;
  display: flex;
  justify-content: center;
  align-items: center;
}

.image-container img {
  width: 100%; 
  height: 100%;
  object-fit: cover;
}

</style>
