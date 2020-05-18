<template>
<div ref="width">
    <div class="landing-wrapper">
        <h1>Welcome</h1>
        <i class="fas fa-chevron-down down-scroll"></i>
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
          width >= 1000 ? this.isDesktop = true : this.isDesktop = false
        }
    },

    mounted() {
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
     font-family: 'Raleway', sans-serif;
     background-image: linear-gradient(rgba(255,255,255,0.3), rgba(63,101,198,0.4));
 }

 .landing-wrapper h1 {
     vertical-align: middle;
     line-height: 90vh;
     text-align: center;
     color: white;
     text-shadow: 1px 1px 4px rgba(50, 81, 158, 0.5);
 }

 .down-scroll {
     color: white;
     display: block;
     margin: 0 auto;
     transform: scale(2.5);
     text-shadow: 1px 1px 4px rgba(50, 81, 158, 0.5);
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
    30% {
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

/* Desktop styling 1000 and up*/
@media screen and (min-width: 1000px) {
  
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