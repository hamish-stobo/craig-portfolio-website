<template>
<div style="position: sticky; top: 0; z-index: 1;">
    <div v-if="isMobile" class="nav-wrapper" v-bind:style="navBackgroundStyle">
        <button @click="toggleNav" v-bind:class="['nav-icon-position', !isShowing ? 'show' : 'hide-icon']" class="nav-icon"><i v-bind:class="['hamburger', !isShowing ? 'show' : 'hide-icon']" class="fas fa-bars"></i></button>
        <div v-bind:class="['nav-list-transition', isShowing ? 'show-list' : 'hide-list']">
            <button @click="toggleNav" v-bind:class="['cross-icon-position', isShowing ? 'show' : 'hide-icon']" class="nav-icon"><i v-bind:class="['cross', isShowing ? 'hide-icon' : 'show']" class="fas fa-times"></i></button>
            <ul v-bind:class="['', isShowing ? 'show' : 'hide-icon']">
                <a @click="toggleNav" href="#about">About</a>
                <a @click="toggleNav" href="#services">Services</a>
                <a @click="toggleNav" href="#currentpositions">My Work</a>
                <a @click="toggleNav" href="#articlespress">Articles & Press</a>
                <a @click="toggleNav" href="#contact">Contact</a>
            </ul>
        </div>
    </div>
    <div v-else-if="!isMobile" class="desktop-nav-wrapper">
        <ul>
                <a href="#about">About</a>
                <a href="#services">Services</a>
                <a href="#currentpositions">My Work</a>
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
            isShowing: false,
            navBackgroundStyle: !this.isShowing ? 'background: rgba(46, 46, 46, 0.8);' : 'background: none;'
        }
    },
    methods: {
      onResize() {
          const width = window.innerWidth
          width >= 765 ? this.isMobile = false : this.isMobile = true
        },
        toggleNav() {
            this.isShowing = !this.isShowing
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
        padding-right: 5px;
        z-index: 2;
        position: absolute;
        right: 1vw;
        top: 3vh;
        width: 200px;
        opacity: 1;
        transition: opacity 0.7s ease-in;
    }

    .nav-icon-position {
        position: absolute;
        top: 20px !important;
        right: 10px !important;
    }

    .cross-icon-position {
        position: absolute;
        top: 20px !important;
        right: 10px !important;
    }

    .show {
        opacity: 1;
        visibility: visible;
        transition: all 0.5s ease-in;
    }

    .hide-icon {
        position: absolute;
        right: 10px;
        opacity: 0 !important;
        visibility: hidden !important;
        margin: 0 !important;
        padding: 0 !important;
        height: 0px !important;
        transition: all 0.5s ease-in;
        background: none;
    }

    .nav-list-transition {
        transition: all 0.7s ease-in;
    }

    .show-list {
        visibility: visible;
        height: auto;
        opacity: 1;
        border-radius: 10px;
        padding-top: 20px;
        padding-right: 20px;
    }

    .hide-list {
        opacity: 0 !important;
        background: none !important;
        visibility: none;
        border-radius: 10px;
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
        outline: none;
        cursor: pointer;
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
        border: none;
        width: auto;
        color: gold;
        width: 100%;
        text-align: center;
        padding: 15px 0;
        height: 60px;
    }

    ul a:hover {
        background: gold;
        color:rgba(46, 46, 46, 0.8);
        font-weight: 600;
        font-size: 1.2rem;
        transition: background 0.4s ease-in;
    }

    ul a:nth-child(5):active {
        border-radius: 0px !important;
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