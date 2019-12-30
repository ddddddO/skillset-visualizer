import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/components/Main'
import LineChart from '@/components/RandomChart/RandomChart'
import SSBarChart from '@/components/SSBarChart/SSBarChart'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Main',
      component: Main
    },
    {
      path: '/line',
      name: 'LineChart',
      component: LineChart
    },
    {
      path: '/ss-bar-chart',
      name: 'SSBarChart',
      component: SSBarChart
    }
  ]
})
