<template>
<div id="services">
    <sui-header>How I can help you:</sui-header>
    <sui-card class="raised service-list-item">
      <sui-card-content>
        <i class="fas fa-business-time"/>
        Business advice
      </sui-card-content>
    </sui-card>
    <sui-card class="raised service-list-item">
      <sui-card-content>
        <i class="fas fa-user-tie" />
        Investment Direction
      </sui-card-content>
    </sui-card>
    <sui-card class="raised service-list-item">
      <sui-card-content>
        <sui-icon
          class="fas fa-business-time"
          shape="circular"
        />
        Government Relations
      </sui-card-content>
    </sui-card>
    <sui-card class="raised service-list-item">
      <sui-card-content>
        <sui-icon
          class="fas fa-business-time"
          shape="circular"
        />
        Capital Raising
      </sui-card-content>
    </sui-card>
    <sui-card class="raised service-list-item">
      <sui-card-content>
        <sui-icon
          class="fas fa-business-time"
          shape="circular"
        />
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
        height: 100vh;
        background: rgba(0,0,0,0.15);
        text-align: center;
        padding: 10px;
    }
    .service-list-item {
        width: 100%;
        text-align: justify;
        padding-left: 20%;
    }
    .service-list-item > i {
        transform: scale(1.5);
    }
</style>