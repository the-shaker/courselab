<template>
  <div class="page">
    <NavBar />

    <main class="page__main">
      <header class="page__header">
        <div class="ap-eyebrow">Аналитика платформы</div>
        <h1 class="ap-h1">Статистика</h1>
      </header>

      <section class="kpi">
        <Card v-for="k in kpis" :key="k.l" :padding="18">
          <div class="kpi__head">
            <span class="kpi__label">{{ k.l }}</span>
            <span
              class="kpi__delta"
              :class="k.g ? 'kpi__delta--up' : 'kpi__delta--down'"
            >
              <Icon :name="k.g ? 'arrowUp' : 'arrowDown'" :size="10" :stroke="2.6" />
              {{ k.d }}
            </span>
          </div>
          <div class="kpi__value ap-mono">{{ k.v }}</div>
          <div class="kpi__sub">{{ k.s }}</div>
          <Sparkline
            :trend="k.g ? 'up' : 'down'"
            :color="k.g ? '#3DDC97' : '#F47174'"
          />
        </Card>
      </section>

      <Card :padding="0" class="chart-card">
        <div class="chart-card__head">
          <div>
            <h3 class="chart-card__title">Активность пользователей</h3>
            <div class="chart-card__sub">
              Уникальные посещения и завершения уроков · 30 дней
            </div>
          </div>
          <div class="chart-card__legends">
            <Legend color="var(--accent)" label="Посещения" />
            <Legend color="var(--info)"   label="Завершено уроков" dashed />
          </div>
        </div>
        <div class="chart-card__body">
          <BigChart />
        </div>
      </Card>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import NavBar    from '../components/NavBar.vue';
import Card      from '../components/Card.vue';
import Icon      from '../components/Icon.vue';
import Sparkline from '../components/Sparkline.vue';
import Legend    from '../components/Legend.vue';
import BigChart  from '../components/BigChart.vue';
import { fetchKPI } from '../api.js';

const kpis = ref([
  { l: 'DAU',                v: '—', d: '—',  g: true,  s: 'против пред. периода' },
  { l: 'Время в сессии',     v: '—', d: '—',  g: true,  s: 'медиана' },
  { l: 'Завершение модулей', v: '—', d: '—',  g: true,  s: 'все курсы' },
  { l: 'Отток (7д)',         v: '—', d: '—',  g: false, s: 'тревога: норма ≤ 2.5%' },
]);

onMounted(async () => {
  const data = await fetchKPI();
  kpis.value = kpis.value.map((kpi, i) => ({ ...kpi, v: data[i].value, d: data[i].delta, g: data[i].growing, s: data[i].subtitle }));
});
</script>

<style scoped>
.page {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}
.page__main {
  flex: 1;
  padding: 28px 32px;
  display: flex;
  flex-direction: column;
  min-height: 0;
}
.page__header { margin-bottom: 22px; }

.kpi {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 12px;
  margin-bottom: 18px;
}
.kpi__head {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.kpi__label {
  font-size: 12px;
  color: var(--text-dim);
}
.kpi__delta {
  display: inline-flex;
  align-items: center;
  gap: 3px;
  font-size: 11px;
  font-weight: 600;
  padding: 2px 7px;
  border-radius: 999px;
}
.kpi__delta--up   { color: var(--accent); background: var(--accent-soft); }
.kpi__delta--down { color: var(--danger); background: var(--danger-soft); }

.kpi__value {
  font-size: 26px;
  font-weight: 600;
  margin-top: 6px;
  letter-spacing: -0.6px;
}
.kpi__sub {
  font-size: 11px;
  color: var(--text-muted);
  margin-top: 2px;
}

.chart-card {
  display: flex;
  flex-direction: column;
  height: 380px;
}
.chart-card__head {
  padding: 18px 22px;
  border-bottom: 1px solid var(--border);
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 12px;
}
.chart-card__title {
  font-size: 15px;
  font-weight: 600;
}
.chart-card__sub {
  font-size: 12px;
  color: var(--text-muted);
  margin-top: 2px;
}
.chart-card__legends {
  display: flex;
  gap: 16px;
  align-items: center;
}
.chart-card__body {
  flex: 1;
  min-height: 0;
  padding: 20px 22px 16px;
  display: flex;
}

@media (max-width: 1100px) {
  .kpi { grid-template-columns: repeat(2, minmax(0, 1fr)); }
}
@media (max-width: 720px) {
  .page__main { padding: 18px 16px; }
  .kpi { gap: 10px; }
  .kpi__value { font-size: 22px; }
}
@media (max-width: 480px) {
  .kpi { grid-template-columns: 1fr; }
  .chart-card__legends { width: 100%; }
}
</style>
