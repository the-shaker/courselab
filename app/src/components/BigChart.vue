<template>
  <div class="chart">
    <Line :data="chartData" :options="chartOptions" />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { Line } from 'vue-chartjs';
import {
  Chart as ChartJS,
  LineElement,
  PointElement,
  LinearScale,
  CategoryScale,
  Filler,
  Tooltip,
} from 'chart.js';
import { fetchChart } from '../api.js';

ChartJS.register(LineElement, PointElement, LinearScale, CategoryScale, Filler, Tooltip);

const max = 15;
const visits    = ref([]);
const completed = ref([]);

onMounted(async () => {
  const data = await fetchChart();
  visits.value    = data.visits;
  completed.value = data.completed;
});

const makeLabels = (arr) => arr.map((_, i) => {
  const date = new Date(2025, 4, 1);
  date.setDate(date.getDate() + i);
  return date.toLocaleDateString('ru-RU', { day: 'numeric', month: 'short' });
});

const chartData = computed(() => ({
  labels: makeLabels(visits.value),
  datasets: [
    {
      label: 'Посещения',
      data: visits.value,
      borderColor: '#3DDC97',
      borderWidth: 2,
      fill: true,
      backgroundColor: (ctx) => {
        const gradient = ctx.chart.ctx.createLinearGradient(0, 0, 0, ctx.chart.height);
        gradient.addColorStop(0, 'rgba(61,220,151,0.32)');
        gradient.addColorStop(1, 'rgba(61,220,151,0)');
        return gradient;
      },
      tension: 0,
      pointRadius: (ctx) => ctx.dataIndex === visits.value.length - 1 ? 4 : 0,
      pointHoverRadius: 4,
      pointBackgroundColor: '#3DDC97',
      pointBorderColor: '#0F1F18',
      pointBorderWidth: 2,
    },
    {
      label: 'Завершено уроков',
      data: completed.value,
      borderColor: '#6BB7E8',
      borderWidth: 1.6,
      borderDash: [5, 4],
      fill: false,
      tension: 0,
      pointRadius: 0,
      pointHoverRadius: 4,
      pointBackgroundColor: '#6BB7E8',
    },
  ],
}));

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  animation: false,
  interaction: { mode: 'index', intersect: false },
  plugins: {
    legend: { display: false },
    tooltip: {
      backgroundColor: '#0F1F18',
      borderColor: '#1F3D2E',
      borderWidth: 1,
      titleColor: '#9DB5A7',
      bodyColor: '#E6F1EA',
      padding: 10,
      callbacks: {
        label: (ctx) => ` ${ctx.dataset.label}: ${ctx.parsed.y}`,
      },
    },
  },
  scales: {
    x: {
      grid: { display: false },
      border: { display: false },
      ticks: {
        color: '#5F7A6C',
        font: { size: 11, family: 'Inter, system-ui, sans-serif' },
        maxTicksLimit: 5,
        maxRotation: 0,
      },
    },
    y: {
      min: 0,
      max,
      grid: {
        color: '#173024',
        borderDash: [2, 4],
      },
      border: { display: false, dash: [2, 4] },
      ticks: {
        color: '#5F7A6C',
        font: { size: 11, family: 'ui-monospace, monospace' },
        stepSize: 3,
        count: 6,
      },
    },
  },
};
</script>

<style scoped>
.chart {
  width: 100%;
  height: 100%;
  min-height: 240px;
}
</style>
