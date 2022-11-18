<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <div>
      <i-svg-vue style="font-size: 50px; fill: red;" />
      <i-icon-vue style="font-size: 50px; fill: red;" />
    <a href="https://vitejs.dev" target="_blank">
      <img src="@images/hero3.png" class="logo" alt="Vite logo" />
    </a>
            <i-svg-vue />
            <a-button>Default Button</a-button>
    <a href="https://vuejs.org/" target="_blank">
    </a>
  </div>
  <HelloWorld msg="Vite + Vue" />
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  /* background: url("@/assets/images/bg-login.jpg") center; */
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
