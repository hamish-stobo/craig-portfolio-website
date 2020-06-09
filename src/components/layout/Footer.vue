<template>
    <div id="footer" class="footer-wrapper">
        <span>© Hamish Stobo 2020</span>
        <a href="https://www.linkedin.com/in/hamishstobo/" target="_blank"><i class="fab fa-linkedin"></i></a>
        <a href="https://github.com/hamish-stobo" target="_blank"><i class="fab fa-github"></i></a>
    </div>
</template>

<script>
export default {
    name: 'Footer'
}
</script>

<style scoped>
    #footer {
        height: 35px;
    }

    .footer-wrapper {
        text-align: center;
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
    }
    .footer-wrapper a, .footer-wrapper span {
        margin: 0 5vw;
    }

    .footer-wrapper a {
        width: 25px;
    }

    .footer-wrapper svg {
        transform: scale(1.8);
        color: rgb(46, 46, 46);
    }
    @media screen and (min-width: 800px) {
        .footer-wrapper a, .footer-wrapper span {
            margin: 0 35px;
        }
    }
</style>