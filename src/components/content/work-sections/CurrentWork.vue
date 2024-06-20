<template>
    <div>
        <ul class="positions-list currWork">
            <a href="https://www.nzwindfarms.co.nz/" target="_blank" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/nz_wind_farms.png'" alt="Local Government Funding Agency Logo">
                <h3 class="company-title">NZ Windfarms Ltd</h3>
                <span class="job-title">Chairman</span>
            </a>
            <a href="https://www.fma.govt.nz/" target="_blank" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/fma_logo.png'" alt="Local Government Funding Agency Logo">
                <h3 class="company-title">Financial Markets Authority</h3>
                <span class="job-title">Chairman</span>
            </a>
            <a href="https://www.lgfa.co.nz/" target="_blank" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/lgfa.png'" alt="Local Government Funding Agency Logo">
                <h3 class="company-title">Local Government Funding Agency</h3>
                <span class="job-title">Chairman</span>
            </a>
            <a href="https://appelloservices.co.nz/" target="_blank" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/apello-services-logo.png'" alt="Apello Services Ltd Logo">
                <h3 class="company-title">Apello Services Ltd</h3>
                <span class="job-title">Director</span>
            </a>
            <a href="http://www.biomarine.co.nz/" target="_blank" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/biomarine-logo.png'" alt="Biomarine Ltd Logo">
                <h3 class="company-title">Biomarine Ltd</h3>
                <span class="job-title">Chairman and Shareholder</span>
            </a>
            <a href="https://www.cressida.co.nz/" target="_blank" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/cressida-capital-logo.png'" alt="Cressida Capital Logo">
                <h3 class="company-title">Cressida Capital</h3>
                <span class="job-title">Board Advisor</span>
            </a>
            <a href="http://www.urbanpartners.co.nz" target="_blank" class="positions-list-item">
                <img class="company-logo" v-bind:src="'./images/company-logos/current/urban-partners.png'" alt="Urban Partners NZ Logo">
                <h3 class="company-title">Urban Partners</h3>
                <span class="job-title">Board Advisor</span>
            </a>
            <span class="positions-list-item">
                <img class="company-logo legend-terrace-logo" v-bind:src="'./images/company-logos/current/legend-terrace-logo.png'" alt="Legend Terrace NZ Logo">
                <h3 class="company-title">Legend Terrace</h3>
                <span class="job-title">Chairman and Shareholder</span>
            </span>
        </ul>
    </div>
</template>

<script>
export default {
    name: "CurrentWork"
    // data() {
    //     return {
    //         isOn: true,
    //         children: [],
    //     }
    // },
    // methods: {
    //     showMe() {
    //         this.isOn = true
    //         console.log(this.isOn)
    //         this.children = this.isOn ? this.children = document.querySelector('.currWork').children : []
    //         this.children.forEach((child, idx, arr) => {
    //             child.setAttribute("style", `animation-delay: ${(idx+1)/(arr.length/1.5)}s`)
    //         })
    //     },
    //     turnOff() {
    //         this.isOn = false
    //         console.log(this.isOn)
    //     }
    // },
    // beforeMount() {
    //     this.showMe()
    // },
    // beforeDestroy() {
    //     this.turnOff()
    // },
    // beforeUpdate() {
    //     this.showMe()
    // }
}
</script>