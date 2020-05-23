<template>
    <div id="articlespress">
        <h2 class="articles-section-title">Articles & Press</h2>
        
    </div>
</template>

<script>
    export default {
        name: 'ArticlesPress'
    }
</script>

<style scoped>
    #articlespress {
        text-align: center;
    }

    .articles-section-title {
        padding-top: 20px;
        font-weight: 700;
        color: rgb(46, 46, 46);
    }

    .articles-list {
        list-style-type: none;
    }

    .article-list-item {
        display: grid;
        grid-template-columns: 30% 70%;
        grid-template-rows: repeat(4, 25%);
        height: 200px;
        border-radius: 10px;
        margin: 10px auto;
    }

    .article-img {

    }

    .article-title {

    }

    .article-source {

    }

    .article-date {

    }

    .article-desc {
        display: none;
    }
</style>