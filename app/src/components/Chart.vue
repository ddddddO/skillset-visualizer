import { Bar, mixins } from 'vue-chartjs'
const { reactiveProp } = mixins

export default {
  extends: Bar,
  mixins: [reactiveProp],
  props: {
    chartdata: {
      type: Object,
      default: null
    }
    options: {
      type: Object,
      default: null
    }
/*
    options = {
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
*/
  },
  mounted () {
    this.renderChart(this.chartData, this.options)
  }
}
