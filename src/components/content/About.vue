<template>
    <div id="about">
        <h2 class="about-section-header">About Me</h2>
        <sui-card-content class="card-content">
            <sui-image class="avatar-img" v-bind:src="'./images/craig-profile-photo.png'" fluid />
            <!-- <img src="../../assets/craig-profile-pic.png" /> -->
            <sui-card-description class="profile-desc">
                I'm Craig Stobo. During the course of my career, I have been a diplomat, economist, CIO and CEO.
                Outside of work, after a number of years mountaineering in NZ and overseas, I continue to be a keen 
                cross-country and downhill skier, sailor and ocean swimmer. I also enjoy drama and literature, 
                and I am a long-time supporter of Otago rugby and the Highlanders. Finally, I am 
                Chair of the Waitaki Boys High School Foundation, and I'm a Taumata of Otago University’s 
                Business School.
            </sui-card-description>
        </sui-card-content>
    </div>
</template>

<script>
    export default {
        name: 'About'
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
        margin-top: 20px;
    }

    .profile-desc {
        text-align: start;
        padding-top: 20px;
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
        .profile-desc {
            font-size: 1.1rem;
            margin: 0;
            padding: 0;
        }
    }
    @media screen and (min-width: 1050px) {
        .profile-desc {
            font-size: 1.2rem;
        }

        .card-content {
            justify-content: space-around;
            max-width: 940px;
        }
    }

</style>