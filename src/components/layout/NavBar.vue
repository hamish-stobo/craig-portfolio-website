<template>
<div>
    <div v-if="isMobile" class="nav-wrapper" v-bind:class="{'bg-on': isShowing}">
        <button v-if="!isShowing" @click="toggleNav"><i class="fas fa-bars hamburger"></i></button>
        <div class="wrapper-div" v-else-if="isShowing">
            <button  @click="toggleNav"><i class="fas fa-times cross"></i></button>
            <ul>
                <a href="#about">About</a>
                <a href="#services">Services</a>
                <a href="#currentpositions">Current Work</a>
                <a href="#articlespress">Articles & Press</a>
                <a href="#contact">Contact</a>
            </ul>
        </div>
    </div>
    <div v-else-if="!isMobile" class="desktop-nav-wrapper">
        <ul>
            <li></li>
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
        toggleNav() {
            this.isShowing = !this.isShowing
            console.log('navbar bool ', this.isShowing)
        }
    }
}
</script>

<style scoped>
    .nav-wrapper {
        position: absolute;
        margin-right: 15px;
        margin-top: 5px;
        padding: 10px;
        border-radius: 10px;
        right: 0;
        z-index: 1;
    }

    .bg-on {
        background: rgba(255,255,255,0.8);
    }

    .hamburger, .cross {
        transform: scale(1.5);
        color: gold;
    }

    button {
        background: none;
        border: none;
        float: right;
        outline: none;
        cursor: pointer;
    }

    .wrapper-div > button {
        display: inline-block;
    }

    ul {
        list-style-type: none;
        padding: none;
        display: flex;
        padding-top: 15px;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: flex-end;
    }

    a {
        text-decoration: none;
        color: black;
        border-bottom: 1px solid black;
        width: 100%;
    }


</style>