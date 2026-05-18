<template>
  <div class="spark">
    <Line :data="chartData" :options="chartOptions" />
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { Line } from 'vue-chartjs';
import {
  Chart as ChartJS,
  LineElement,
  PointElement,
  LinearScale,
  CategoryScale,
  Filler,
} from 'chart.js';

ChartJS.register(LineElement, PointElement, LinearScale, CategoryScale, Filler);

const props = defineProps({
  color: { type: String, default: 'var(--accent)' },
  trend: { type: String, default: 'up' },
});

const upPoints    = [10, 8, 12, 9, 14, 11, 17, 14, 20, 18, 24, 22, 28];
const downPoints  = [22, 24, 21, 23, 19, 21, 17, 19, 15, 16, 13, 14, 10];

const points = computed(() => props.trend === 'up' ? upPoints : downPoints);

const chartData = computed(() => ({
  labels: points.value.map(() => ''),
  datasets: [{
    data: points.value,
    borderColor: props.color,
    borderWidth: 1.5,
    fill: true,
    backgroundColor: (ctx) => {
      const gradient = ctx.chart.ctx.createLinearGradient(0, 0, 0, ctx.chart.height);
      gradient.addColorStop(0, props.color + '40');
      gradient.addColorStop(1, props.color + '00');
      return gradient;
    },
    tension: 0,
    pointRadius: 0,
  }],
}));

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  animation: false,
  plugins: { legend: { display: false }, tooltip: { enabled: false } },
  scales: {
    x: { display: false },
    y: { display: false },
  },
};
</script>

<style scoped>
.spark {
  width: 100%;
  height: 32px;
  margin-top: 10px;
}
</style>
