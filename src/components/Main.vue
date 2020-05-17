<template>
    <div>
        <About />
        <Services />
        <CurrentPositions />
        <ArticlesPress />
        <Footer />
    </div>
</template>

<script>

import About from './content/About'
import Services from './content/Services'
import CurrentPositions from './content/CurrentPositions'
import ArticlesPress from './content/ArticlesPress'
import Footer from './layout/Footer'

export default {
    name: 'Main',
    components: {
        About, 
        ArticlesPress, 
        CurrentPositions, 
        Services,
        Footer
    }
}
</script>