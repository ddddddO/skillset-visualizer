<template>
  <div class="main">
    <h2>{{ title }}</h2>
    <h3>Links!</h3>
    <ul>
      <li><router-link to="line">RandomLine</router-link></li>
      <li><router-link to="ss-bar-chart">SSBarChart</router-link></li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Main',
  data () {
    return {
      title: 'TOP'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
