<template>
  <div class="container">
    <h1>skillset-visualizer</h1>
    <bar-chart
      v-if="loaded"
      :chart-data="chartdata"
      :options="options"/>
    <br>
    <select v-model="selectedCategory">
      <option v-for="category in categories" v-bind:key="category">{{ category }}</option>
    </select>
    <br>
    <input type="range" v-model="selectedNum" min="0" max="5">
    <br>
    <button v-on:click="putChartData">Push!</button>
  </div>
</template>

<script>
import BarChart from './Chart.js'

export default {
  name: 'BarChartContainer',
  components: { BarChart },
  data: () => ({
    loaded: false,
    chartdata: null,
    categories: null,
    nums: null,
    selectedNum: 3,
    selectedCategory: ''
  }),
  async mounted () {
    this.loaded = false
    try {
      const graphData = await fetch('http://localhost:8081/fetch')
        .then(function (resp) {
          return resp.json()
        })
        .then(function (json) {
          const strJson = JSON.stringify(json)
          return JSON.parse(strJson)
        })

      console.log('--debug!--')
      console.log(graphData)

      this.categories = Object.keys(graphData)
      this.nums = Object.values(graphData)

      this.chartdata = {
        labels: this.categories,
        datasets: [{
          label: 'now(sample)',
          data: this.nums,
          borderWidth: 1,
          borderColor: [
            'rgba(85, 222, 22, 1)',
            'rgba(244, 67, 146, 1)',
            'rgba(67, 114, 244, 1)',
            'rgba(164, 162, 146, 1)',
            'rgba(29, 240, 191, 1)',
            'rgba(240, 29, 78, 1)'
          ],
          backgroundColor: [
            'rgba(85, 222, 22, 0.4)',
            'rgba(244, 67, 146, 0.4)',
            'rgba(67, 114, 244, 0.4)',
            'rgba(164, 162, 146, 0.4)',
            'rgba(29, 240, 191, 0.4)',
            'rgba(240, 29, 78, 0.4)'
          ]
        }]
      }
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
  },
  methods: {
    reInputChartdata: function (category, num) {
      const cb = function (cat) {
        return cat === category
      }
      const index = this.categories.findIndex(cb)

      if (index === -1) { return }

      this.nums[index] = num
      this.chartdata = {
        labels: this.categories,
        datasets: [{
          label: 'now(sample)',
          data: this.nums,
          borderWidth: 1,
          borderColor: [
            'rgba(85, 222, 22, 1)',
            'rgba(244, 67, 146, 1)',
            'rgba(67, 114, 244, 1)',
            'rgba(164, 162, 146, 1)',
            'rgba(29, 240, 191, 1)',
            'rgba(240, 29, 78, 1)'
          ],
          backgroundColor: [
            'rgba(85, 222, 22, 0.4)',
            'rgba(244, 67, 146, 0.4)',
            'rgba(67, 114, 244, 0.4)',
            'rgba(164, 162, 146, 0.4)',
            'rgba(29, 240, 191, 0.4)',
            'rgba(240, 29, 78, 0.4)'
          ]
        }]
      }
    },
    putChartData: function () {
      const data = this.mappingChartData(this.categories, this.nums)

      // ref: https://github.com/axios/axios#instance-methods
      const apiClient = require('axios').create({
        headers: {
          'Content-Type': 'application/json'
        }
      })
      apiClient.put('http://localhost:8081/put/1', data)
        .then(function (resp) {
          console.log('debug resp', resp)
        }).catch(function (err) {
          console.log('ERROR...', err)
        }).finally(function () {
          // FIXME: 再描画がうまくいかない。機能としてはいまのところ不要だけど、、
          // ref: https://qiita.com/shoridevel/items/11638860eb04dfe56df7
          // this.$router.go({path: this.$router.currentRoute.path, force: true})
        })// .bind(this))
    },
    mappingChartData: function (keys, values) {
      if (keys.length !== values.length) {
        console.log('no match keys/values')
        return {}
      }

      let kv = {}
      for (var i = 0; i < keys.length; i++) {
        kv[keys[i]] = values[i]
      }
      return kv
    }
  },
  watch: {
    selectedNum () {
      this.reInputChartdata(this.selectedCategory, this.selectedNum)
    }
  }
}
</script>
