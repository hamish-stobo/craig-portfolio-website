<template>
    <div id="contact">
        <h2 class="contact-section-title">Contact Me</h2>
        <div class="social-icons-wrapper">
            <a href="https://nz.linkedin.com/in/craig-stobo-4b40361a2" target="_blank"><i class="fab fa-linkedin-in"></i></a>
            <a href="mailto:stobo@xtra.co.nz" target="_blank"><i class="far fa-envelope"></i></a>
            <a href="tel:+6421733751" target="_blank"><i class="fas fa-phone"></i></a>
        </div>
    </div>
</template>

<script>
export default {
    name: "contact"
}
</script>

<style scoped>
    #contact {
        text-align: center;
        height: 40vh;
    }

    .contact-section-title {
        margin-top: 20px;
        padding-top: 20px;
        margin-bottom: 50px;
        font-weight: 700;
        color: gold;
    }
    
    .social-icons-wrapper {
        display: flex;
        flex-flow: row wrap;
        justify-content: center;
        align-items: center;
    }

    .social-icons-wrapper a {
        border-radius: 50%;
        width: 65px;
        height: 65px;
        margin: 10px 5vw;
    }

    .social-icons-wrapper svg {
        color: rgb(46, 46, 46);
        background: gold;
        border-radius: 50%;
        width: 65px;
        height: 65px;
        padding: 15px;
    }
</style>