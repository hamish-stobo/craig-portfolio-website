<template>
    <div>
        <ul class="positions-list">
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/previous/fliway-group-logo.png'" alt="Fliway Group Logo">
                <h3 class="company-title">Fliway Group</h3>
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/previous/bureau-logo.png'" alt="Buereau Ltd Logo">
                <h3 class="company-title">Bureau Ltd</h3> 
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/previous/ocg-logo.png'" alt="OCG Consulting Ltd Logo">
                <h3 class="company-title">OCG Consulting Ltd</h3>
                <span class="job-title">Chairman</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/previous/callaghan-innovation-logo.png'" alt="Industrial Resarch Ltd Logo">
                <h3 class="company-title">Industrial Research Ltd</h3>
                <span class="job-title">Director</span>
            </div>
            <div class="positions-list-item">
                <img style="height: 80px; width: 125px;" class="company-logo" v-bind:src="'./images/company-logos/previous/NZGIF-logo.png'" alt="NZGIF Logo">
                <h3 class="company-title">New Zealand Green Investment Finance Ltd</h3>
                <span class="job-title">Advisor</span>
            </div>
            <div class="positions-list-item">
                <img style="height: 90px;" class="company-logo" v-bind:src="'./images/company-logos/previous/treasury-logo.png'" alt="NZ Treasury Logo">
                <h3 class="company-title">New Zealand Treasury</h3>
                <span class="job-title">Advisor</span>
            </div>
            <div class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/previous/lgnz-logo.png'" alt="Local Government New Zealand Logo">
                <h3 class="company-title">Local Government New Zealand</h3>
                <span class="job-title">Advisor</span>
            </div>
        </ul>
    </div>
</template>

<script>
export default {
    name: "PreviousWork"
}
</script>