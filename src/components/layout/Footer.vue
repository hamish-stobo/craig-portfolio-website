<template>
    <div id="footer" class="footer-wrapper">
        <span>© Hamish Stobo2020</span>
        <i class="fab fa-linkedin"></i>
        <i class="fab fa-github"></i>
    </div>
</template>

<script>
export default {
    name: 'Footer'
}
</script>

<style scoped>
    .footer-wrapper {
        background: salmon;
        text-align: center;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;
    }



</style>