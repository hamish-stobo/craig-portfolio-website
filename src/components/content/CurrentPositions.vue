<template>
    <div id="currentpositions">
        <h2 class="current-positions-header">Current Positions</h2>
        <ul class="positions-list">
            <div class="positions-list-item">
                <img class="company-logo" src="" alt="Saturn Advice Ltd Logo">
                <h3 class="company-title">Saturn Advice Ltd</h3>
                <span class="job-title">Chairman and Non-Executive Director</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" src="" alt="Elevation Capital Logo">
                <h3 class="company-title">Elevation Capital</h3>
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" src="" alt="Fliway Group Logo">
                <h3 class="company-title">Fliway Group</h3>
                <span class="job-title">Independent Non-Executive Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" src="" alt="Local Government Funding Agency Logo">
                <h3 class="company-title">Local Government Funding Agency</h3>
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" src="" alt="Precinct Properties Ltd Logo">
                <h3 class="company-title">Precinct Properties Ltd</h3>
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" src="" alt="AIG Insurance New Zealand Logo">
                <h3 class="company-title">AIG Insurance New Zealand</h3>
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" src="" alt="Apello Services Ltd Logo">
                <h3 class="company-title">Apello Services Ltd</h3>
                <span class="job-title">Managing Director</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" src="" alt="Biomarine Ltd Logo">
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
        margin: 10px 0;
        border-radius: 10px;
        box-shadow: 1px 1px 4px rgba(255, 217, 0, 0.637);
        display: grid;
        grid-template-columns: 1fr 2fr;
        grid-template-rows: 1fr 1fr;
    }

    .company-logo {
        grid-row: 1 / 3;
        grid-column: 1 / 2;
    }

    .company-title {
        grid-row: 1 / 2;
        grid-column: 2 / 3;
    }

    .job-title {
        grid-row: 2 / 3;
        grid-column: 2 / 3;
    }
</style>