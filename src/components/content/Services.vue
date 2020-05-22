<template>
<div id="services">
    <h2 class="ui header services-header">How I can help you:</h2>
    <ul class="list-wrapper">
      <div>
        <img class="services-img" v-bind:src="'./images/services-photos/photo-1450101499163-c8848c66ca85.jpg'" alt="Man signing paper"/>
        <span class="service-title">Business advice</span>
      </div>
      <div>
        <span class="service-title">Investment Direction</span>
        <img class="services-img" v-bind:src="'./images/services-photos/photo-1454165804606-c3d57bc86b40.jpg'" alt="Men discussing sheet of paper"/>
      </div>
      <div>
        <img class="services-img" v-bind:src="'./images/services-photos/photo-1460925895917-afdab827c52f.jpg'" alt="Laptop display graph with figures" />
        <span class="service-title">Government Relations</span>
      </div>
      <div>
        <span class="service-title">Capital Raising</span>
        <img class="services-img" v-bind:src="'./images/services-photos/photo-1516321318423-f06f85e504b3.jpg'" alt="People pointing at laptop screen"/>
      </div>
      <div>
        <img class="services-img" v-bind:src="'./images/services-photos/photo-1517048676732-d65bc937f952.jpg'" alt="People sitting around table"/>
        <span class="service-title">B2B Consulting</span>
      </div>
      <div>
        <span class="service-title">Cash Flow Optimisation</span>
        <img class="services-img" v-bind:src="'./images/services-photos/photo-1523875194681-bedd468c58bf.jpg'" alt="Chess board"/>
      </div>
    </ul>
    <!-- <sui-card class="raised service-list-item">
      <sui-card-content class="item-content">
        <i class="fas fa-business-time"/>
        Business advice
      </sui-card-content>
    </sui-card>
    <sui-card class="raised service-list-item">
      <sui-card-content class="item-content">
        <i class="fas fa-user-tie" />
        Investment Direction
      </sui-card-content>
    </sui-card>
    <sui-card class="raised service-list-item">
      <sui-card-content class="item-content">
        <i class="fas fa-comments" />
        Government Relations
      </sui-card-content>
    </sui-card>
    <sui-card class="raised service-list-item">
      <sui-card-content class="item-content">
        <i class="fas fa-project-diagram" />
        Capital Raising
      </sui-card-content>
    </sui-card>
    <sui-card class="raised service-list-item">
      <sui-card-content class="item-content">
        <i class="fas fa-ruler-horizontal" />
        Board minutes
      </sui-card-content>
    </sui-card> -->
</div>
</template>

<script>
    export default {
        name: 'Services'
    }
</script>

<style scoped>
    #services {
        text-align: center;
        padding: 30px 0 0 0;
    }

    .services-header {
      color: rgb(46, 46, 46);
    }

    .list-wrapper {
      width: 100vw;
      list-style-type: none;
      margin-top: 50px;
    }

    .list-wrapper > div {
      width: 100%;
      display: flex;
      flex-flow: row nowrap;
      justify-content: center;
      align-items: center;
      background: rgb(46, 46, 46);
    }

    

    ul > div:nth-child(2n) {
        background: gold;
    }

    ul > div:nth-child(2n) > span {
        color: rgb(46, 46, 46);
        font-weight: bold;
    }

    .service-title {
      width: 50%;
      margin: 0;
      display: block;
      color: gold;
      font-size: 1rem;
    }

    .services-img {
      width: 50%;
      margin: 0;
    }

    @media screen and (min-width: 650px) {
      .service-title {
        font-size: 1.2rem;
      }
    }

    @media screen and (min-width: 650px) {
      .service-title {
        font-size: 1.5rem;
      }
    }






    .service-list-item {
        width: 100%;
        text-align: justify;
        padding-left: 20%;
        max-width: 550px;
        margin: 10px auto;
        padding: 0;
    }
    .item-content {
        display: flex;
        flex-flow: row nowrap;
        justify-content: flex-start;
        align-items: center;
        width: 220px;
        margin: 0 auto !important;
    }
    .item-content > svg {
        margin-right: 30px;
        transform: scale(2) !important;
        color: cornflowerblue;
    }
</style>