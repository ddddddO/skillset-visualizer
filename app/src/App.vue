<template>
  <div id="app">
    <h1>{{ title }}</h1>
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App',
  data: function () {
    return {
      title: 'Skillset Visualizer'
    }
  }
}
</script>

<style>
#app {
  text-align: center;
  margin-top: 60px;
}
</style>
