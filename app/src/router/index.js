import Vue from 'vue'
import Router from 'vue-router'
import Top from '@/components/Top'
import LineChart from '@/components/RandomChart/RandomChart'
import SSBarChart from '@/components/SSBarChart/SSBarChart'
import TodoList from '@/components/Todo/Todo'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Top',
      component: Top
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
    },
    {
      path: '/todo',
      name: 'TODOLIST',
      component: TodoList
    }
  ]
})
