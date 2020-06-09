<template>
<ul class="article-list">
        <a href="https://www.nzherald.co.nz/business/news/article.cfm?c_id=3&objectid=12330405" style="text-decoration: none; color: gold" target="_blank" class="article-list-item">
            <img class="article-img" v-bind:src="'./images/articles -press-images/nz-herald-2020.jpg'" alt="Worker in Hi-Vis and Hardhat">
            <span class="article-title">Shovel-ready, but not work-ready</span>
            <span class="article-source">NZ Herald</span>
            <span class="article-date">9 May 2020</span>
        </a>
        <a href="https://www.youtube.com/watch?v=fDo0HFOe8EU" style="text-decoration: none; color: gold" target="_blank" class="article-list-item">
            <img class="article-img" v-bind:src="'./images/articles -press-images/UoO-2017-business-school-interview.png'" alt="Otago University Business School video thumbnail">
            <span class="article-title">Q & A with Craig Stobo</span>
            <span class="article-source">University of Otago Business School</span>
            <span class="article-date">20 Oct 2017</span>
        </a>
        <a href="https://www.nzherald.co.nz/business/news/article.cfm?c_id=3&objectid=11415720" style="text-decoration: none; color: gold" target="_blank" class="article-list-item">
            <img class="article-img" v-bind:src="'./images/articles -press-images/nz-herald-2015.jpg'" alt="Aircraft Landing">
            <span class="article-title">Government assets sheet needs right balance</span>
            <span class="article-source">NZ Herald</span>
            <span class="article-date">12 Mar 2015</span>
        </a>
        </ul>
</template>

<script>
export default {
    name: 'Press'
}
</script>