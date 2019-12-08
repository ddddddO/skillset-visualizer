import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import HelloBar from '@/components/BaseChart'
import LineChart from '@/components/RandomChart'
import BarChartContainer from '@/components/ChartContainer'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/bar',
      name: 'HelloBar',
      component: HelloBar
    },
    {
      path: '/line',
      name: 'LineChart',
      component: LineChart
    },
    {
      path: '/bar-chart',
      name: 'BarChart',
      component: BarChartContainer
    }
  ]
})
