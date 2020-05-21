<template>
<div ref="width">
    <div class="landing-wrapper">
        <h1>
          I'm Craig Stobo.
          <h3>Welcome to my personal website.</h3>
        </h1>
        <a href="#about">
          <i class="fas fa-chevron-down down-scroll"></i>
        </a>
    </div>
    <div v-if="isDesktop" class="desktop-bg"></div>
    <ul v-else class="crossfade">
        <li></li>
        <li></li>
        <li></li>
        <li></li>
    </ul>
</div>
</template>

<script>
export default {
    name: 'LandingPage',
    data() {
      return {
        isDesktop: false
      }
    },
    methods: {
      onResize() {
          const width = window.innerWidth
          width >= 1200 ? this.isDesktop = true : this.isDesktop = false
        }
    },

    mounted() {
      this.onResize()
      // Register an event listener when the Vue component is ready
      window.addEventListener('resize', this.onResize) 
    },

    beforeDestroy() {
      // Unregister the event listener before destroying this Vue instance
      window.removeEventListener('resize', this.onResize)
    }
    }
</script>

<style scoped>
/* mobile styling */
 .landing-wrapper {
     height: 100vh;
     padding-top: 40vh;
     padding-bottom: 4vh;
     font-family: 'Raleway', sans-serif;
     background-image: linear-gradient(rgba(255,255,255,0.3), rgba(63,101,198,0.4));
     display: flex;
     flex-flow: column nowrap;
     justify-content: space-between;
     align-items: center
 }

 .landing-wrapper h1 {
     vertical-align: middle;
     text-align: center;
     color: gold;
     text-shadow: 2px 2px 4px rgba(50, 81, 158, 0.8);
     font-family: 'Raleway', sans-serif;
 }

 .landing-wrapper h3 {
   font-family: 'Raleway', sans-serif;
 }

 .down-scroll {
     color: rgba(63,101,198,0.8);
     display: block;
     margin: 0 auto;
     transform: scale(2.5);
     text-shadow: 1px 1px 4px rgba(50, 81, 158, 0.5);
     animation: arrowAnimation 2s ease-in-out infinite;
 }

 @keyframes arrowAnimation {
   0% {
     transform: scale(2.5)
   }
   50% {
     transform: scale(3);
     color: gold
   }
   100% {
     transform: scale(2.5)
   }
 }
/* background-animation */
.crossfade li { 
    width: 100%;
    height: 100%;
    position: absolute;
    top: 0;
    left: 0;
    background-size: cover;
    background-position: 50% 50%;
    background-repeat: no-repeat;
    opacity: 0;
    z-index: -1;
    animation: imageAnimation 20s linear infinite;
  }

  .crossfade li:nth-child(1) { 
    background-image: url('../assets/waitaki-boys.jpg')
  }
  .crossfade li:nth-child(2) { 
    background-image: url('../assets/otago-uni-clocktower.jpg');
    animation-delay: 5s;
  }
  .crossfade li:nth-child(3) { 
    background-image: url('../assets/Queenstown.jpg');
    animation-delay: 10s;
  }
  .crossfade li:nth-child(4) { 
    background-image: url('../assets/skyline-Auckland.jpg');
    animation-delay: 15s;
  }

  @keyframes imageAnimation { 
    0% { 
      opacity: 0;
    }
    10% {
      opacity: 1;
      animation-timing-function: ease-in;
    }
    20% {
      opacity: 1;
      animation-timing-function: ease-out;
    }
    50% {
      opacity: 0
    }
    100% {
      opacity: 0
    }
    
  }

  /* Older browser support - .no-cssanimations class added by modernizr */
.crossfade li {
	opacity: 1;
}
/* styling for medium size screens */
@media screen and (min-width: 750px) { 
  .landing-wrapper h1 {
     font-size: 3rem;
 }
 .down-scroll {
     transform: scale(3.5);
 }
}

/* Desktop styling 1000 and up*/
@media screen and (min-width: 1200px) {

  .landing-wrapper {
    background: none;
  }

  .desktop-bg {
      width: 100%;
      height: 100%;
      position: absolute;
      top: 0;
      left: 0;
      background-image: url('../assets/desktop-bg.png');
      background-size: cover;
      background-position: 50% 50%;
      background-repeat: no-repeat;
      z-index: -1;
  }
}
 
</style>