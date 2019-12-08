<script>
import { Bar } from 'vue-chartjs'

let data = {
  labels: [
    'frontend',
    'backend',
    'database',
    'infrastructure',
    'network',
    'facilitation'
  ],
  datasets: [{
    label: 'now(sample)',
    data: [1, 2, 3, 3, 2, 4],
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

let options = {
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

export default {
  name: 'HelloBar',
  extends: Bar,
  mounted () {
    this.renderChart(data, options)
  }
}
</script>
