<script>
export default {
	props: ['username', 'profilePic', 'content', 'timestamp', 'msgStyle'],
	methods: {
	},
	mounted() {

  },
	computed: {
		TimeStyle(){
			return this.msgStyle.wasSentByUser ? "margin-right: 15px;" : "margin-right: 15px;"
		},
		ExtraSpace(){
				return this.msgStyle.wasSentByUser ? "display: block; margin-right: 0; margin-left: auto;" : ""
		},
    MessageStyle(){
				//'rgb(255, 221, 70)'
				//'rgb(255, 209, 0)'
        let bgColor = this.msgStyle.wasSentByUser ? 'rgb(255, 221, 70)' : 'rgb(238, 238, 238)'
				let alignment = this.msgStyle.wasSentByUser ? 'right' : 'left'

        return 'color:' + this.msgStyle.color +';' +
				' font-size:'+ this.msgStyle.fontSize + 'rem;' +
				" background-color:" + bgColor + ";" +
				" text-align:"+ alignment+";"+
				" font-weight: 400"
    }
  }
}
</script>


<template>

	<div id="Parent">

		<div v-if="!this.msgStyle.wasSentByUser" id="profileImageOfOtherPerson {{content}}" class="image-container">
				<img :src="this.profilePic"/>
		</div>

		<div id= "extraSpace" :style="ExtraSpace"> </div>		

		<div id="ComplexMessage" :style="MessageStyle">
			<div v-if="!this.msgStyle.wasSentByUser" id="username">
					{{ username }}
			</div>

			<div class="message" :style="TimeStyle">
					{{ content }}
			</div>

			<div class="time">
					{{ timestamp }}
			</div>
		</div>

	</div>
	

</template>


<style>

#Parent {
	display: flex;
	overflow-wrap: break-word; /* Modern equivalent for word-wrap */

}

#profileImageOfOtherPerson {
	visibility: visible;
}

.time {
	stroke-width: 50%;
	font-weight: 10;
	font-style: italic;
	margin-top: 5px;
	text-align: right;
}

#username {
	font-weight: bold;
	text-align: left;
	/*margin-left: 10px;*/
}

#ComplexMessage{
		display: inline-block; /* Makes the div size fit its content */
		/*border: 2px solid rgba(0, 0, 0, .25);
		background-color: rgba(0, 0, 0, .25);*/
		border-radius: 15px;
		padding-left: 8px;
		padding-right: 8px;
		padding-top: 5px;
		padding-bottom: 5px;
		margin-bottom: 5px;

		max-width: 75%;
}

.message {
	display: inline-block; /* Makes the div size fit its content */

  word-break: break-word; /* Breaks long words */
  overflow-wrap: break-word; /* Modern equivalent for word-wrap */

}
</style>
