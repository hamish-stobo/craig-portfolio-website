<template>
    <div id="currentpositions">
        <h2 class="current-positions-header">Current Positions</h2>
        <ul class="positions-list">
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/Saturn-Advice-Logo_new 1.png'" alt="Saturn Advice Ltd Logo">
                <h3 class="company-title">Saturn Advice Ltd</h3>
                <span class="job-title">Chairman and Non-Executive Director</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/elevation-capital-logo.png'" alt="Elevation Capital Logo">
                <h3 class="company-title">Elevation Capital</h3>
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/fliway-group-logo.png'" alt="Fliway Group Logo">
                <h3 class="company-title">Fliway Group</h3>
                <span class="job-title">Independent Non-Executive Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/lgfa.png'" alt="Local Government Funding Agency Logo">
                <h3 class="company-title">Local Government Funding Agency</h3>
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/precinct-properties-logo.png'" alt="Precinct Properties Ltd Logo">
                <h3 class="company-title">Precinct Properties Ltd</h3>
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/AIG-logo.png'" alt="AIG Insurance New Zealand Logo">
                <h3 class="company-title">AIG Insurance New Zealand</h3>
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/apello-services-logo.png'" alt="Apello Services Ltd Logo">
                <h3 class="company-title">Apello Services Ltd</h3>
                <span class="job-title">Managing Director</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/biomarine-logo.png'" alt="Biomarine Ltd Logo">
                <h3 class="company-title">Biomarine Ltd</h3>
                <span class="job-title">Director</span>
            </div>
        </ul>
    </div>
</template>

<script>
    export default {
        name: 'CurrentPositions'
    }
</script>

<style scoped>
    #currentpositions {
        text-align: center;
    }

    .current-positions-header {
        padding-top: 20px;
        font-weight: 400;
        color: gold;
    }

    .positions-list {
        list-style-type: none;
        margin: 50px 10px 50px 10px;
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
    }

    .company-logo {
        grid-row: 1 / 3;
        grid-column: 1 / 2;
        display: block;
        margin: auto;
        padding-left: 4px;
        width: 130px;
    }

    .company-title {
        grid-row: 1 / 2;
        grid-column: 2 / 3;
    }

    .job-title {
        grid-row: 2 / 3;
        grid-column: 2 / 3;
    }

    .company-logo:nth-child(1) {
        max-width: 130px;
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
        /* .positions-list {
            -moz-column-count: 2;
            -webkit-column-count: 2;
            column-count: 2;
        }

        .positions-list-item:nth-child(1) {
            margin-top: 0;
        } */

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