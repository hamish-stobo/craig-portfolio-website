<template>
<div id="services">
    <sui-header>How I can help you:</sui-header>
    <sui-card class="raised service-list-item">
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
    </sui-card>
</div>
</template>

<script>
    export default {
        name: 'Services'
    }
</script>

<style scoped>
    #services {
        background: rgba(0,0,0,0.15);
        text-align: center;
        padding: 30px 10px;
    }
    .service-list-item {
        width: 100%;
        text-align: justify;
        padding-left: 20%;
        max-width: 500px;
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
    }

/* @media screen and (min-width: 375px) {
    .item-content {
        margin-left: 20vw !important;
    }
}

@media screen and (min-width: 450px) {
    .item-content {
        margin-left: 25vw !important;
    }
} */
</style>