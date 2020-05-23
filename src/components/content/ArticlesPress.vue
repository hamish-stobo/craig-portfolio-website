<template>
    <div id="articlespress">
        <h2 class="articles-section-title">Articles & Press</h2>
        <ul class="article-list">
        <a href="https://www.nzherald.co.nz/business/news/article.cfm?c_id=3&objectid=12330405" style="text-decoration: none; color: gold" target="_blank" class="article-list-item">
            <img class="article-img" v-bind:src="'./images/articles -press-images/nz-herald-2020.jpg'" alt="Worker in Hi-Vis and Hardhat">
            <span class="article-title">Craig Stobo: Shovel-ready, but not work-ready</span>
            <span class="article-source">NZ Herald</span>
            <span class="article-date">9 May 2020</span>
            <p class="article-desc"></p>
        </a>
        <a href="https://www.youtube.com/watch?v=fDo0HFOe8EU" style="text-decoration: none; color: gold" target="_blank" class="article-list-item">
            <img class="article-img" v-bind:src="'./images/articles -press-images/UoO-2017-business-school-interview.png'" alt="Otago University Business School video thumbnail">
            <span class="article-title">Q & A with Craig Stobo (Professional Director and Business Owner)</span>
            <span class="article-source">University of Otago Business School</span>
            <span class="article-date">20 Oct 2017</span>
            <p class="article-desc"></p>
        </a>
        <a href="https://www.nzherald.co.nz/business/news/article.cfm?c_id=3&objectid=11415720" style="text-decoration: none; color: gold" target="_blank" class="article-list-item">
            <img class="article-img" v-bind:src="'./images/articles -press-images/nz-herald-2015.jpg'" alt="Aircraft Landing">
            <span class="article-title">Craig Stobo: Government assets sheet needs right balance</span>
            <span class="article-source">NZ Herald</span>
            <span class="article-date">12 Mar 2015</span>
            <p class="article-desc"></p>
        </a>
        <a href="https://taxpolicy.ird.govt.nz/sites/default/files/2004-other-report-stobo-taxation-investment-income.pdf" style="text-decoration: none; color: gold" target="_blank" class="article-list-item">
            <img class="article-img" v-bind:src="'./images/articles -press-images/report-to-minister-of-finance.png'" alt="">
            <span class="article-title">Towards Consensus on the Taxation of Investment Income</span>
            <span class="article-source">Inland Revenue Department</span>
            <span class="article-date">October 2004</span>
            <p class="article-desc"></p>
        </a>
        </ul>
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
        padding: 40px 0 100px 0;
    }

    .articles-section-title {
        padding-top: 20px;
        font-weight: 700;
        color: rgb(46, 46, 46);
    }

    .article-list {
        list-style-type: none;
        margin-top: 50px;
    }

    .article-list-item {
        display: grid;
        grid-template-columns: 45% 55%;
        grid-template-rows: 50% 50%;
        margin: 10px auto;
        background: rgb(46, 46, 46);
        text-align: left;
        height: 100px;
        padding-bottom: 3px;
    }

    .article-img {
        grid-column: 1 / 2;
        grid-row: 1 / 4;
        margin: 0;
        width: 130px;
        display: inline-block;
    }

    .article-title {
        font-size: 1rem;
        grid-row: 1 / 2;
        grid-column: 2 / 3;
        align-self: start;
        margin: 0;
        display: inline-block;
        line-height: 1.2rem;
    }

    .article-source {
        grid-row: 2 / 3;
        grid-column: 2 / 3;
        font-size: 0.8rem;
        line-height: 1rem;
        color: white;
    }

    .article-date {
        grid-row: 2 / 3;
        grid-column: 2 / 3;
        font-size: 0.7rem;
        align-self: end;
        color: white;
    }

    .article-desc {
        display: none;
    }
</style>