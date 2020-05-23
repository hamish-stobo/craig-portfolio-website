<template>
<div style="position: sticky; top: 0; z-index: 1;">
    <div v-if="isMobile" class="nav-wrapper" v-bind:class="{'bg-on': isShowing}">
        <button v-if="!isShowing" @click="toggleNav"><i class="fas fa-bars hamburger"></i></button>
        <div class="wrapper-div" v-else-if="isShowing">
            <button  @click="toggleNav"><i class="fas fa-times cross"></i></button>
            <ul>
                <a @click="toggleNav" href="#about">About</a>
                <a @click="toggleNav" href="#services">Services</a>
                <a @click="toggleNav" href="#currentpositions">Current Work</a>
                <a @click="toggleNav" href="#articlespress">Articles & Press</a>
                <a @click="toggleNav" href="#contact">Contact</a>
            </ul>
        </div>
    </div>
    <div v-else-if="!isMobile" class="desktop-nav-wrapper">
        <ul>
                <a href="#about">About</a>
                <a href="#services">Services</a>
                <a href="#currentpositions">Current Work</a>
                <a href="#articlespress">Articles & Press</a>
                <a href="#contact">Contact</a>
            </ul>
    </div>
</div>
</template>

<script>
export default {
    name: 'NavBar',
    data() {
        return {
            isMobile: true,
            isShowing: false
        }
    },
    methods: {
      onResize() {
          const width = window.innerWidth
          width >= 765 ? this.isMobile = false : this.isMobile = true
        },
        toggleNav() {
            this.isShowing = !this.isShowing
            console.log('navbar bool ', this.isShowing)
        }
    },
    created() {
        this.onResize()
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
    .nav-wrapper {
        border-radius: 10px;
        padding-right: 15px;
        padding-top: 15px;
        z-index: 2;
        position: absolute;
        right: 1vw;
        top: 3vh;
        width: 200px;
    }

    .bg-on {
        background: rgba(46, 46, 46, 0.8);
        padding-right: 0;
    }


    .hamburger {
        background: rgba(46, 46, 46, 0.8);
        width: 20px;
        height: 20px;
        padding: 2px;
        border-radius: 5px;
    }

    .hamburger, .cross {
        transform: scale(2);
        color: gold;
    }

    .cross {
        width: 18px;
        height: 18px;
    }

    button {
        background: none;
        border: none;
        float: right;
        outline: none;
        cursor: pointer;
    }

    .wrapper-div > button {
        display: block;
        position: absolute;
        right: 15px;
    }

    ul {
        list-style-type: none;
        padding: 0 !important;
        display: flex;
        padding-top: 15px;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: flex-end;
        width: 100%;
        margin-block-start: 0 !important;
        margin-block-end: 0 !important;
    }

    a {
        text-decoration: none;
        color: gold;
        border-bottom: 1px solid rgba(255, 217, 0);
        width: 100%;
        padding: 5px 5px 5px 15px;
        height: 40px;
        line-height: 200%;
    }

    ul > a:nth-child(5) {
        border: none;
    }

    ul > *:active {
        color:rgba(46, 46, 46, 0.8);
        background: gold;
        border-bottom: rgba(46, 46, 46, 0.8);
    }

    ul > a:nth-child(5):active {
        border: none;
        border-radius: 0 0 10px 10px;
    }

@media screen and (min-width: 765px) {

    ul {
        flex-flow: row nowrap;
        justify-content: space-evenly;
        align-items: center;
        padding: 0;
        background: rgba(46, 46, 46, 0.8);
    }
    ul a {
        display: inline-block;
        white-space: nowrap;
        height: 100%;
        border: none;
        width: auto;
        color: gold;
        width: 100%;
        text-align: center;
        padding: 15px 0;
    }

    ul a:hover {
        background: gold;
        color:rgba(46, 46, 46, 0.8);
        font-weight: 600;
        font-size: 1.2rem;
    }
}

  @media screen and (min-width: 1050px) {
      ul > a {
          font-size: 1.2rem;
      }
      ul a:hover {
        font-size: 1.3rem;
    }
  } 
</style>