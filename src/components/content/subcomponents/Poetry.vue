<template>
<ul class="article-list">
        <a href="./images/docs/Covid 19 Verse.pdf" style="text-decoration: none; color: gold" target="_blank" class="article-list-item">
            <span class="article-title">COVID-19 Vignettes</span>
            <span class="article-date">13 May 2020</span>
        </a>
        <a href="./images/docs/For David.pdf" style="text-decoration: none; color: gold" target="_blank" class="article-list-item">
            <span class="article-title">For David</span>
            <span class="article-date">May 2020</span>
        </a>
        </ul>
</template>

<script>
export default {
    name: 'poetry'
}
</script>

<style scoped>
    .article-list-item {
        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
    }

    .article-date {
        align-self: center;
    }
</style>