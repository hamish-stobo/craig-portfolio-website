<template>
    <div id="about">
        <h2 class="about-section-header">About Me</h2>
        <sui-card-content class="card-content">
            <sui-image class="avatar-img" v-bind:src="'./images/craig-profile-photo.png'" fluid />
            <div class="dynamic-content-wrapper">
            <div class="button-div">
                <h3 v-bind:class="{'underline': !showingQualifications}" @click="displayQualifications(0)">Personal Summary</h3>
                <h3 v-bind:class="{'underline': showingQualifications}" @click="displayQualifications(1)">Qualifications</h3>
            </div>
                <sui-card-description class="profile-desc" v-bind:class="{'HideAbout': showingQualifications}">
                    I’m Craig Stobo. I am a professional director and business owner. I have previously worked 
                    as a diplomat, economist, Chief Investment Officer and CEO. Outside of work I have climbed 
                    mountains in NZ, Peru, Pakistan, and Nepal. I continue to ski, sail and compete in ocean swims. 
                    I also enjoy drama and literature, and I am a long-time supporter of Otago rugby and the 
                    Highlanders. Finally, I am Chair of the Waitaki Boys High School Foundation, and I'm a 
                    Taumata of Otago University’s Business School.
                </sui-card-description>
                <ul class="qualifications-wrapper" v-bind:class="{'ShowQualifications': showingQualifications}">
                    <li>BA (Hons) First Class Economics Otago</li>
                    <li>Wharton Business School (Advanced Management Programme)</li>
                    <li>Associate Member CFA Society New Zealand</li>
                    <li>Chartered Fellow NZ Institute of Directors</li>
                </ul>
            </div>
        </sui-card-content>
    </div>
</template>

<script>
    export default {
        name: 'About',
        data() {
            return {
                showingQualifications: false
            }
        },
        methods: {
            displayQualifications(num) {
                return num === 1 ? this.showingQualifications = true : this.showingQualifications = false
            }
        }
    }
</script>

<style scoped>
    .card-content {
        display: flex;
        flex-flow: column nowrap;
        align-items: center;
        justify-content: space-between;
        margin: 20px auto 50px auto;
        width: 80%;
        text-align: center;
    }

    .card-content > * {
     color: gold;
    }

    .about-section-header {
        margin: 30px auto 0 auto;
        font-weight: 400;
        text-align: center;
        color: gold;
    }

    .avatar-img {
        border-radius: 50%;
        max-width: 260px;
        height: auto;
        margin: 20px auto 0 auto;
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
        padding: 5px;
        margin: 0 5px 5px 5px;
        border: 1px solid gold;
        border-radius: 5px;
    }

    .underline {
        background: gold;
        color: rgb(46, 46, 46);
        transition: background 0.7s ease-in;
    }

    .selectorArrow {
        animation: selectorArrowAnimation 2s ease-in-out infinite;
    }

    @keyframes selectorArrowAnimation {
        0% {
            transform: scale(1.3)
        }
        50% {
            transform: scale(1.5);
        }
        100% {
            transform: scale(1.3)
        }
    }

    .profile-desc, .qualifications-wrapper {
        text-align: start;
        padding: 25px;
        transition: opacity 0.7s ease-in;
    }

    .profile-desc {
        opacity: 1;
    }

    .HideAbout {
        opacity: 0 !important;
        visibility: hidden !important;
        height: 0px !important;
        width: 0px !important;
        margin: 0 !important;
        padding: 0 !important;
    }

    .qualifications-wrapper {
        opacity: 0;
        visibility: hidden;
        height: 0px;
        width: 0px;
        margin: 0;
        padding: 0;
    }

    .ShowQualifications {
        opacity: 1 !important;
        visibility: visible !important;
        width: auto !important;
        height: auto !important;
        padding: 25px !important;
    }

    .qualifications-wrapper li {
        margin: 10px 0;
    }

    @media screen and (min-width: 500px) {
        .card-content {
            margin-bottom: 100px;
        }

        .about-section-header {
            font-size: 2rem !important;
        }
        .profile-desc {
            max-width: 410px;
            margin: 0 auto;
        }
    }

    @media screen and (min-width: 765px) {
        .card-content {
            margin: 50px auto 200px auto;
            flex-flow: row nowrap;
            justify-content: space-evenly;
            align-items: center;
            width: 95%;
        }
        .about-section-header {
            margin-top: 0;
        }
        .dynamic-content-wrapper {
            width: 410px;
            height: 190px;
        }

        .profile-desc, .qualifications-wrapper li {
            font-size: 1.1rem;
        }
    }
    @media screen and (min-width: 1050px) {
        .profile-desc, .qualifications-wrapper li {
            font-size: 1.2rem;
        }

        .card-content {
            justify-content: space-around;
            max-width: 940px;
        }
    }

</style>