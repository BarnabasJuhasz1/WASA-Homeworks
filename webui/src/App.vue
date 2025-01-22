<script setup>

import { RouterLink, RouterView } from 'vue-router'
import { createApp } from 'vue'

</script>

<script>
export default {
  data() {
    return {
    }
  }
}
</script>


<template>
  <div style="background-color: var(--main-outline); display: flex; flex-direction: column; height: 100vh;">

    <div style="display: block; height: 100vh">
      <header id="header">
        <div class="content" style="color: var(--font-light); margin-bottom: 0; align-items: center; font-size: x-large;">
          WasaText
        </div>
      </header>
      <main style="display: block;height: 90vh">
        <div style="display: flex; ">
          <div style="width:5vh; background-color: var(--main-outline);">
          </div>

          <RouterView />
        </div>
       
      </main>
    </div>

    <!-- <div style="display: flex; flex-grow: 1; height: auto; margin-bottom: 0; margin-top: auto;">

    </div> -->

    <!-- <div id="pageFooter" style="display: flex; flex-direction: column;"> 
     
    </div> -->
    <footer style="height: 5vh">
      <div id="pageFooter">
        Made by Barnabas Juhasz
      </div>
    </footer>

  </div>

</template>

<style>

body, header, h1 {
	margin: 0;
	padding: 0;
  background-color: var(--background-light);
}


#header {
	text-align: center;
  vertical-align: middle;
  display: flex;
  height: 5vh;
  padding:0;
  align-items: center;
  justify-content: center;
  background-color: var(--main-outline);
}


#pageFooter {
  display: flex;
  height: 5vh;
  align-items: center;
  justify-content: center;
  color: var(--font-light);
  background-color: var(--main-outline);
}

</style>