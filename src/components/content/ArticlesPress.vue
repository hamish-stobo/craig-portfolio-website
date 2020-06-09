<template>
    <div id="articlespress">
        <h2 class="articles-section-title">Articles, Press & Poetry</h2>
        <div class="button-div">
            <h3 v-bind:class="{'underline': showingArticles}" @click="toggler(0)">Articles</h3>
            <h3 v-bind:class="{'underline': showingPress}" @click="toggler(1)">Press</h3>
            <h3 v-bind:class="{'underline': showingPoetry}" @click="toggler(2)">Poetry</h3>
        </div>
        <Articles v-if="showingArticles" />
        <Press v-else-if="showingPress" />
        <Poetry v-else-if="showingPoetry" />
    </div>
</template>

<script>
import Articles from './subcomponents/Articles'
import Press from './subcomponents/Press'
import Poetry from './subcomponents/Poetry'
    export default {
        name: 'ArticlesPress',
        components: {
            Articles,
            Press,
            Poetry
        },
        data() {
            return {
                showingArticles: true,
                showingPress: false,
                showingPoetry: false
            }
        },
        methods: {
            toggler(num) {
                if(num === 0) {
                    this.showingArticles = true
                    this.showingPress = false
                    this.showingPoetry = false
                } else if(num === 1) {
                    this.showingArticles = false
                    this.showingPress = true
                    this.showingPoetry = false
                } else if(num === 2) {
                    this.showingArticles = false
                    this.showingPress = false
                    this.showingPoetry = true
                }
            }
        }
    }
</script>

<style>
    #articlespress {
        text-align: center;
        padding: 40px 0 100px 0;
    }

    .articles-section-title {
        padding-top: 20px;
        font-weight: 700;
        color: rgb(46, 46, 46);
    }

    .button-div {
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: center;
        margin-top: 20px;
        cursor: pointer;
        min-width: 320px;
    }

     .button-div h3 {
        padding: 5px 10px;;
        margin: 0 15px 5px 15px;
        border: 1px solid rgb(46, 46, 46);
        border-radius: 5px;
    }

    .underline {
        background: rgb(46, 46, 46);
        color: gold;
        transition: background 0.7s ease-in;
    }

    .article-list {
        list-style-type: none;
        margin: 50px auto 0;
        margin-block-start: 0 !important;
        margin-block-end: 0 !important;
    }

    .article-list-item {
        display: grid;
        grid-template-columns: 45% 55%;
        grid-template-rows: 50% 25% 25%;
        margin: 10px auto;
        background: rgb(46, 46, 46);
        text-align: left;
        height: 110px;
        padding-bottom: 3px;
        box-shadow: 2px 2px 4px rgba(46, 46, 46, 0.6);
        max-width: 400px;
        border-radius: 10px;
    }

    .article-img {
        grid-column: 1 / 2;
        grid-row: 1 / 4;
        margin: 0;
        max-height: 100%;
        max-width: 140px;
        display: inline-block;
        border-radius: 10px 0 0 0;
    }

    .article-title {
        font-size: 1rem;
        grid-row: 1 / 2;
        grid-column: 2 / 3;
        align-self: center;
        margin: 0;
        display: inline-block;
        line-height: 1.2rem;
        margin: 2px 2px 0 0;
    }

    .article-source {
        grid-row: 2 / 3;
        grid-column: 2 / 3;
        font-size: 0.8rem;
        line-height: 1rem;
        color: white;
        align-self: center;
    }

    .article-date {
        grid-row: 3 / 4;
        grid-column: 2 / 3;
        font-size: 0.7rem;
        align-self: start;
        color: white;
    }

    .article-desc {
        display: none;
    }

    @media screen and (min-width: 400px) {
        .article-list-item {
            grid-template-columns: 150px auto;
        }
        .space-me {
            margin-top: 5px;
        }
    }

    @media screen and (min-width: 700px) {
        .article-list-item {
            height: 120px;
        }
    }

    @media screen and (min-width: 850px) {
        
        .article-list {
            display: flex;
            flex-flow: row wrap;
            justify-content: space-evenly;
            align-items: center;
            max-width: 920px;
        }

        .article-list-item {
            margin: 30px 0;
            width: 400px;
        }
    }

    @media screen and (min-width: 920px) { 
        #articlespress {
            padding: 40px 4vw 100px;
        }
    }
    @media screen and (min-width: 1050px) { 
        #articlespress {
            padding: 40px 8vw 100px;
        }
    }
    @media screen and (min-width: 1200px) { 
        #articlespress {
            padding: 40px 12vw 100px;
        }
    }

</style>