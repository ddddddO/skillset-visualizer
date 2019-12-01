// import Bar from 'vue-chartjs'
import { Bar, mixins } from 'vue-chartjs'
const { reactiveProp } = mixins

export default {
  name: 'BarChart',
  extends: Bar,
  mixins: [reactiveProp],
  props: {
    chartdata: {
      type: Object,
      default: null
    },
    options: {
      type: Object,
      default: null
    }
  },
  mounted () {
    console.log('-in BarChart-')
    this.renderChart(this.chartdata, this.options)
    console.log('-in BarChart end-')
  }
}
