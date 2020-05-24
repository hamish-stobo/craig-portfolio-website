<template>
<div ref="width" style="background: white;">
    <div class="landing-wrapper">
        <h1>
          I'm Craig Stobo.
          <h3>Welcome to my website.</h3>
        </h1>
        <a href="#about">
          <i class="fas fa-chevron-down down-scroll"></i>
        </a>
    </div>
    <div v-if="isDesktop" class="desktop-bg" style="background-image: url('./images/landing-page/desktop-bg.png');"></div>
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
     align-items: center;
     z-index: 2;
 }

 .landing-wrapper  * {
   z-index: 3;
 }

 .landing-wrapper h1 {
     vertical-align: middle;
     text-align: center;
     color: gold;
     font-family: 'Raleway', sans-serif;
 }

 .landing-wrapper h3 {
   font-family: 'Raleway', sans-serif;
 }

 .landing-wrapper > h1, .landing-wrapper > h3 {
   text-shadow: 3px 3px 8px rgba(46, 46, 46);
 }

 .down-scroll {
     color: gold;
     display: block;
     margin: 0 auto;
     transform: scale(2.5);
     animation: arrowAnimation 2s ease-in-out infinite;
 }

 @keyframes arrowAnimation {
   0% {
     transform: scale(2.5)
   }
   50% {
     transform: scale(3);
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
    z-index: 1;
    animation: imageAnimation 20s linear infinite 0s;
  }

  .crossfade li:nth-child(1) { 
    background-image: url('https://inboxdesign.ibcdn.nz/media/2018_02_07_2016_07_27_waitaki-boys-high-school-entrance.jpg');
    background-position: 30% 50%;
  }

  .crossfade li:nth-child(2) { 
    background-image: url('https://julietetelandresen.com/wp-content/uploads/2018/01/7027833-queenstown-new-zealand-wakatipu-lake.jpg');
    animation-delay: 5s;
  }
  .crossfade li:nth-child(3) { 
    background-image: url('https://www.timeshighereducation.com/sites/default/files/styles/institution_12col_card/public/institution_images/clocktower_blossom_2.jpg');
    animation-delay: 10s;
  }
  .crossfade li:nth-child(4) { 
    background-image: url('https://upload.wikimedia.org/wikipedia/commons/6/60/Bayswater_Marina_Auckland_New_Zealand.jpg');
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
      opacity: 0;
    }
    100% {
      opacity: 0;
    }
    
  }
/* styling for medium size screens */

@media screen and (min-width: 500px) { 
  .landing-wrapper h1 {
     font-size: 2.5rem;
  }

  .landing-wrapper h3 {
    font-size: 1.4rem;
  }

  @keyframes arrowAnimation {
   0% {
     transform: scale(3);
   }
   50% {
     transform: scale(3.5);
   }
   100% {
     transform: scale(3);
   }
 }
}

@media screen and (min-width: 750px) { 
  .landing-wrapper h1 {
     font-size: 3.6rem;
  }

  .landing-wrapper h3 {
    font-size: 1.6rem;
  }

 @keyframes arrowAnimation {
   0% {
     transform: scale(3.5);
   }
   50% {
     transform: scale(4);
   }
   100% {
     transform: scale(3.5);
   }
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
      background-size: cover;
      background-position: 50% 50%;
      background-repeat: no-repeat;
      z-index: 1;
  }

  .landing-wrapper h1 {
     font-size: 4rem;
 }
}
 
</style>