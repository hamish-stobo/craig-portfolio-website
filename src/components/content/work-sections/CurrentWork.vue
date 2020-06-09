<template>
    <div>
        <ul class="positions-list currWork" v-bind:class="{'isOn': isOn}">
            <a href="https://saturnadvice.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/Saturn-Advice-Logo_new 1.png'" alt="Saturn Advice Ltd Logo">
                <h3 class="company-title">Saturn Advice Ltd</h3>
                <span class="job-title">Chairman and Shareholder</span>
            </a>
            <a href="https://www.pssnz.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/saturn-invest-logo.png'" alt="Saturn Invest Ltd Logo">
                <h3 class="company-title">Saturn Invest Ltd</h3>
                <span class="job-title">Chairman and Shareholder</span>
            </a>
            <a href="https://www.elevationcapital.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/elevation-capital-logo.png'" alt="Elevation Capital Logo">
                <h3 class="company-title">Elevation Capital</h3>
                <span class="job-title">Chairman and Shareholder</span>
            </a>
            <a href="https://www.lgfa.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/lgfa.png'" alt="Local Government Funding Agency Logo">
                <h3 class="company-title">Local Government Funding Agency</h3>
                <span class="job-title">Chairman</span>
            </a>
            <a href="https://www.precinct.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/precinct-properties-logo.png'" alt="Precinct Properties Ltd Logo">
                <h3 class="company-title">Precinct Properties Ltd</h3>
                <span class="job-title">Chairman</span>
            </a>
            <a href="https://www.aig.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/AIG-logo.png'" alt="AIG Insurance New Zealand Logo">
                <h3 class="company-title">AIG Insurance New Zealand</h3>
                <span class="job-title">Chairman</span>
            </a>
            <a href="https://appelloservices.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/apello-services-logo.png'" alt="Apello Services Ltd Logo">
                <h3 class="company-title">Apello Services Ltd</h3>
                <span class="job-title">Director</span>
            </a>
            <a href="http://www.biomarine.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/biomarine-logo.png'" alt="Biomarine Ltd Logo">
                <h3 class="company-title">Biomarine Ltd</h3>
                <span class="job-title">Chairman and Shareholder</span>
            </a>
            <a href="https://www.cressida.co.nz/" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/cressida-capital-logo.png'" alt="Cressida Capital Logo">
                <h3 class="company-title">Cressida Capital</h3>
                <span class="job-title">Board Advisor</span>
            </a>
            <a href="http://www.urbanpartners.co.nz" target="_blank" style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/urban-partners.png'" alt="Urban Partners NZ Logo">
                <h3 class="company-title">Urban Partners</h3>
                <span class="job-title">Board Advisor</span>
            </a>
            <span style="text-decoration: none; color: inherit" class="positions-list-item">
                <img class="company-logo legend-terrace-logo" v-bind:src="'./images/company-logos/current/legend-terrace-logo.png'" alt="Legend Terrace NZ Logo">
                <h3 class="company-title">Legend Terrace</h3>
                <span class="job-title">Chairman and Shareholder</span>
            </span>
        </ul>
    </div>
</template>

<script>
export default {
    name: "CurrentWork",
}
</script>