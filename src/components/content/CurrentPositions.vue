<template>
    <div id="currentpositions">
        <h2 class="current-positions-header">My Work</h2>
        <div class="positions-toggler" >
            <h3 @click="changeSelected(1)" v-bind:class="{'underline': currentWorkSelected}">Current</h3>
            <h3 @click="changeSelected(0)" v-bind:class="{'underline': !currentWorkSelected}">Previous</h3>
        </div>
        <CurrentWork v-if="currentWorkSelected" />
        <PreviousWork v-else-if="!currentWorkSelected" />
    </div>
</template>

<script>
import CurrentWork from './work-sections/CurrentWork'
import PreviousWork from './work-sections/PreviousWork'
    export default {
        name: 'CurrentPositions',
        components: {
            CurrentWork,
            PreviousWork
        },
        data() {
          return {
                currentWorkSelected: true
              }
        },
        methods: {
            changeSelected(currentSelection) {
                //if bool is true, return false, else return true.
                //prevents overly-zealous toggling
                currentSelection === 1 ? this.currentWorkSelected = true : this.currentWorkSelected = false
            }
        }
    }
</script>

<style>
    #currentpositions {
        text-align: center;
    }

    .current-positions-header {
        padding-top: 20px;
        font-weight: 400;
        color: gold;
    }

    .positions-toggler {
        color: gold;
        display: flex;
        flex-flow: row nowrap;
        width: 100%;
        justify-content: space-around;
        align-items: center;
        margin: 10px auto;
        max-width: 400px;
    }

    .positions-toggler h3 {
        border: 1px solid gold;
        border-radius: 5px;
        padding: 5px 10px;
        margin: 5px;
    }

    .underline {
        background: gold;
        color: rgb(46, 46, 46);
        transition: background 0.4s ease-in;
    }

    .positions-toggler h3:hover {
        cursor: pointer;
    }

    .positions-list {
        list-style-type: none;
        margin: 50px 10px 50px 10px;
        transition: opacity 0.4s ease-in;
    }

    .positions-list-item {
        background: gold;
        margin: 10px auto;
        border-radius: 15px;
        box-shadow: 1px 1px 4px rgba(255, 217, 0, 0.637);
        display: grid;
        grid-template-columns: 1fr 1.5fr;
        grid-template-rows: 1fr 1fr;
        height: 110px;
        max-width: 400px;
        transition: opacity 0.5s ease-in;
    }

    .company-logo {
        grid-row: 1 / 3;
        grid-column: 1 / 2;
        display: block;
        margin: auto;
        padding-left: 4px;
        max-width: 130px;
        align-self: center;
        max-height: 100%;
        transition: opacity 0.6s ease-in;
    }

    .company-title {
        grid-row: 1 / 2;
        grid-column: 2 / 3;
        transition: opacity 0.7s ease-in;
    }

    .job-title {
        grid-row: 2 / 3;
        grid-column: 2 / 3;
        transition: opacity 0.7s ease-in;
    }

    @media screen and (min-width: 700px) {
        .positions-list-item {
            height: 120px;
        }
        .company-logo {
            width: 140px;
        }
    }

    @media screen and (min-width: 850px) {
        .positions-list {
            display: flex;
            flex-flow: row wrap;
            justify-content: space-evenly;
            align-items: center;
        }

        .positions-list-item {
            width: 390px;
        }
    }   

    @media screen and (min-width: 920px) { 
        #currentpositions {
            margin: 0 4vw;
        }
    }
    @media screen and (min-width: 1050px) { 
        #currentpositions {
            margin: 0 8vw;
        }
    }
    @media screen and (min-width: 1200px) { 
        #currentpositions {
            margin: 0 12vw;
        }
    }
</style>