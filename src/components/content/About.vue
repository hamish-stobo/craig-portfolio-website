<template>
    <div id="about">
        <sui-card class="card">
            <sui-card-content class="card-content">
                <sui-image v-bind:src="'./images/craig-profile-photo.png'" fluid />
                <!-- <img src="../../assets/craig-profile-pic.png" /> -->
                <sui-card-header class="about-header">Craig Stobo</sui-card-header>
                <sui-card-description class="profile-desc">
                    During the course of his career, Craig has been a diplomat, economist, CIO and CEO. <br />Outside of work, after a number of years mountaineering in NZ and overseas, he continues to be a keen cross-country and downhill skier, sailor and ocean swimmer. <br />He also enjoys drama and literature and is a long-time supporter of Otago rugby and the Highlanders. <br />Finally, he is Chair of the Waitaki Boys High School Foundation, and is a Taumata of Otago University’s Business School.
                </sui-card-description>
            </sui-card-content>
        </sui-card>
    </div>
</template>

<script>
    export default {
        name: 'About'
    }
</script>

<style scoped>
    #about {
        background: rgba(63,101,198,0.8);
        padding: 20px;
    }

    .card {
        margin: 0 auto;
        background: rgba(255,255,255,0.9);
        width: 100%;
        max-width: 550px;
    }

    .card-content {
        display: flex;
        flex-flow: column nowrap;
        align-items: center;
        justify-content: space-evenly;
    }

    .card-content > .header {
        padding-top: 10px;
    }

    .profile-desc {
        color: black !important;   
    }

    @media screen and (min-width: 500px) {
        .about-header {
            font-size: 2rem !important;
        }
        .profile-desc {
            max-width: 410px;
            margin: 0 auto;
        }
    }
</style>