<template>
<div>
    <div class="landing-wrapper">
        <h1>Welcome</h1>
        <i class="fas fa-chevron-down"></i>
    </div>
    <ul class="crossfade">
        <li></li>
        <li></li>
        <li></li>
        <li></li>
    </ul>
</div>
</template>

<script>
export default {
    name: 'LandingPage'
}
</script>

<style scoped>
/* mobile styling */
 .landing-wrapper {
     height: 100vh;
     font-family: 'Raleway', sans-serif;
     background-image: linear-gradient(rgba(255,255,255,0.6), rgba(63,101,198,0.6));
 }

 .landing-wrapper h1 {
     vertical-align: middle;
     line-height: 90vh;
     text-align: center;
     color: white;
 }

 .fa-chevron-down {
     color: white;
     display: block;
     margin: 0 auto;
     transform: scale(2.5);
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
    animation: imageAnimation 16s linear infinite;
  }

  .crossfade li:nth-child(1) { 
    background-image: url('../assets/waitaki-boys.jpg')
  }
  .crossfade li:nth-child(2) { 
    background-image: url('../assets/otago-uni-clocktower.jpg');
    animation-delay: 4s;
  }
  .crossfade li:nth-child(3) { 
    background-image: url('../assets/Queenstown.jpg');
    animation-delay: 8s;
  }
  .crossfade li:nth-child(4) { 
    background-image: url('../assets/skyline-Auckland.jpg');
    animation-delay: 12s;
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
 
</style>