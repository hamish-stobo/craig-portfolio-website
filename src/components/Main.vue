<template>
    <div @click="hideNavbar">
        <About />
        <Services />
        <CurrentPositions />
        <ArticlesPress />
        <Contact />
        <Footer />
    </div>
</template>

<script>

import About from './content/About'
import Services from './content/Services'
import CurrentPositions from './content/CurrentPositions'
import ArticlesPress from './content/ArticlesPress'
import Contact from './content/Contact'
import Footer from './layout/Footer'

export default {
    name: 'Main',
    components: {
        About, 
        ArticlesPress, 
        CurrentPositions, 
        Services,
        Contact,
        Footer
    },
    methods: {
        hideNavbar() {
            this.$emit('hidenavbar')
        }
    }
}
</script>