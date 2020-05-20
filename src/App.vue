<template>
  <div id="app">
    <LandingPage />
    <NavBar />
    <Main />
  </div>
</template>

<script>
import LandingPage from './components/LandingPage.vue'
import NavBar from './components/layout/NavBar'
import Main from './components/Main'
export default {
  name: 'App',
  components: {
    LandingPage,
    NavBar,
    Main
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  font-family: 'Open Sans', sans-serif; 
}
html {
  scroll-behavior: smooth;
}
</style>
