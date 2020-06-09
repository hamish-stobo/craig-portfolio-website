<template>
    <div>
        <ul class="positions-list">
            <a href="https://www.fliway.com/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/previous/fliway-group-logo.png'" alt="Fliway Group Logo">
                <h3 class="company-title">Fliway Group</h3>
                <span class="job-title">Chairman</span>
            </a>
            <a href="https://www.bureau.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/previous/bureau-logo.png'" alt="Buereau Ltd Logo">
                <h3 class="company-title">Bureau Ltd</h3> 
                <span class="job-title">Chairman</span>
            </a>
            <a href="https://www.ocg.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/previous/ocg-logo.png'" alt="OCG Consulting Ltd Logo">
                <h3 class="company-title">OCG Consulting Ltd</h3>
                <span class="job-title">Chairman</span>
            </a>
            <a href="https://www.callaghaninnovation.govt.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/previous/callaghan-innovation-logo.png'" alt="Industrial Resarch Ltd Logo">
                <h3 class="company-title">Industrial Research Ltd</h3>
                <span class="job-title">Director</span>
            </a>
            <a href="https://nzgif.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img style="height: 80px; width: 125px;" class="company-logo" v-bind:src="'./images/company-logos/previous/NZGIF-logo.png'" alt="NZGIF Logo">
                <h3 class="company-title">New Zealand Green Investment Finance</h3>
                <span class="job-title">Advisor</span>
            </a>
            <a href="https://treasury.govt.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img style="height: 90px;" class="company-logo" v-bind:src="'./images/company-logos/previous/treasury-logo.png'" alt="NZ Treasury Logo">
                <h3 class="company-title">New Zealand Treasury</h3>
                <span class="job-title">Advisor</span>
            </a>
            <a href="https://www.lgnz.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/previous/lgnz-logo.png'" alt="Local Government New Zealand Logo">
                <h3 class="company-title">Local Government New Zealand</h3>
                <span class="job-title">Advisor</span>
            </a>
        </ul>
    </div>
</template>

<script>
export default {
    name: "PreviousWork"
}
</script>