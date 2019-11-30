<template>
  <div class="container">
    <bar-chart
      v-if="loaded"
      :chartdata="chartdata"
      :options="options"/>
  </div>
</template>

<script>
// import Chart from './Chart.vue'
import Bar from './Chart.vue'
// import Bar from 'vue-chartjs'

export default {
  name: 'BarChart',
  // extends: Bar,
  components: { Bar },
  data: () => ({
    loaded: false,
    chartdata: null
    // options: null
  }),
  async mounted () {
    this.loaded = false
    try {
      const graphData = await fetch('http://localhost:8081/fetch')
        .then(function (resp) {
          return resp.json()
        })
        .then(function (json) {
          return JSON.stringify(json)
        })

      console.log('--debug!--')
      console.log(graphData)

      this.chartData = graphData
      this.options = {
        scales: {
          xAxes: [{
            scaleLabel: {
              labelString: 'category',
              display: true
            }
          }],
          yAxes: [{
            ticks: {
              beginAtZero: true,
              min: 0,
              max: 5,
              stepSize: 1
            }
          }]
        }
      }
      this.loaded = true
    } catch (e) {
      console.error(e)
    }
  }
}
</script>
