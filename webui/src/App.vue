<script setup>

// import { RouterLink, RouterView } from 'vue-router'
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
  <div style="background-color: rgb(50, 50, 50); display: flex;  flex-direction: column; height: 100vh;">

    <div style="display: block;">
      <header>
        <div id="header">
          <div class="content">
            <h1>WasaText</h1>
          </div>
        </div>
      </header>
      <main>
        <RouterView />
      </main>
    </div>

    <div style="display: flex; flex-grow: 1; height: auto; margin-bottom: 0; margin-top: auto;">

    </div>


    <div id="pageFooter" style="display: flex; flex-direction: column;"> 
      <p style="text-align: center;"> Made by Barnabas Juhasz</p>
    </div>

  </div>

</template>

<style>

body, header, h1 {
	margin: 0;
	padding: 0;
}


#header {
  padding-top: 1rem;
	text-align: center;
	height: 4rem;
  color: rgb(146, 146, 146);
}


#pageFooter {
	height: 10vh;
  color: rgb(146, 146, 146);
}

</style>