<template>
  <div id="app">
    <LandingPage />
    <NavBar ref="navbarComponent" />
    <Main v-on:hidenavbar="hideNav" />
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
  },
  methods: {
    hideNav() {
      // console.log(this.$refs)
      this.$refs.navbarComponent.toggleNav(1)
    }
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
body {
  background: rgb(46, 46, 46);
}

#services, #articlespress, #footer {
          background: gold !important;
}

@media screen and (max-width: 765px) {
  #about, #services, #currentpositions, #articlespress, #contact {
    padding-top: 30px !important;
  }
  .about-section-header, .services-header, .current-positions-header, .articles-section-title, .contact-section-title {
    padding-top: 20px !important;
    margin-top: 0 !important;
  }
}


@media screen and (min-width: 765px) {
        #about, #services, #currentpositions, #articlespress, #contact {
            padding-top: 65px !important;
        }
}
</style>
