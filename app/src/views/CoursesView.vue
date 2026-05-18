<template>
  <div class="page">
    <NavBar />

    <main class="page__main">
      <div class="page__header">
        <div>
          <div class="ap-eyebrow">Контент платформы</div>
          <h1 class="ap-h1">
            Курсы <span class="muted">· {{ filteredCourses.length }} из {{ courses.length }}</span>
          </h1>
        </div>
      </div>

      <div class="toolbar">
        <div class="tabs">
          <button
            v-for="f in filters"
            :key="f.key"
            class="tabs__tab"
            :class="{ 'tabs__tab--active': activeFilter === f.key }"
            @click="activeFilter = f.key"
          >
            {{ f.l }}
            <span class="tabs__count ap-mono">{{ f.c }}</span>
          </button>
        </div>

        <div class="search">
          <Icon name="search" :size="15" class="search__icon" />
          <input
            v-model="searchQuery"
            class="search__input"
            placeholder="Найти курс по названию или автору..."
          />
        </div>
      </div>

      <div v-if="filteredCourses.length" class="grid">
        <CourseCard
          v-for="c in filteredCourses"
          :key="c.title"
          v-bind="c"
        />
      </div>
      <div v-else class="empty">
        Курсы не найдены
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import NavBar     from '../components/NavBar.vue';
import Icon       from '../components/Icon.vue';
import CourseCard from '../components/CourseCard.vue';
import { fetchCourses } from '../api.js';

const STATUS = {
  published: { l: 'Опубликован',  c: 'var(--accent)',     s: 'var(--accent-soft)' },
  review:    { l: 'На модерации', c: 'var(--warn)',       s: 'var(--warn-soft)'   },
  draft:     { l: 'Черновик',     c: 'var(--text-muted)', s: 'var(--bg)'          },
};

const courses      = ref([]);
const activeFilter = ref('all');
const searchQuery  = ref('');

onMounted(async () => {
  courses.value = await fetchCourses();
});

const filters = [
  { key: 'all',       l: 'Все курсы',      c: computed(() => courses.value.length) },
  { key: 'published', l: 'Опубликованные', c: computed(() => courses.value.filter(c => c.statusKey === 'published').length) },
  { key: 'review',    l: 'На модерации',   c: computed(() => courses.value.filter(c => c.statusKey === 'review').length)    },
  { key: 'draft',     l: 'Черновики',      c: computed(() => courses.value.filter(c => c.statusKey === 'draft').length)     },
];

const filteredCourses = computed(() => {
  const q = searchQuery.value.toLowerCase().trim();
  return courses.value
    .filter(c => activeFilter.value === 'all' || c.statusKey === activeFilter.value)
    .filter(c => !q || c.title.toLowerCase().includes(q) || c.author.toLowerCase().includes(q))
    .map(c => ({ ...c, status: STATUS[c.statusKey] }));
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
}

.page__header {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  margin-bottom: 22px;
  flex-wrap: wrap;
  gap: 16px;
}

.toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  margin-bottom: 18px;
  flex-wrap: wrap;
}

.tabs {
  display: flex;
  gap: 4px;
  padding: 4px;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  overflow-x: auto;
  scrollbar-width: none;
}
.tabs::-webkit-scrollbar { display: none; }

.tabs__tab {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 7px 14px;
  border-radius: var(--radius-sm);
  color: var(--text-dim);
  font-size: 13px;
  font-weight: 500;
  white-space: nowrap;
  cursor: pointer;
}
.tabs__tab:hover { color: var(--text); }
.tabs__tab--active { background: var(--surface-3); color: var(--text); }
.tabs__count { font-size: 11px; color: var(--text-muted); }
.tabs__tab--active .tabs__count { color: var(--accent); }

.search {
  display: flex;
  align-items: center;
  gap: 10px;
  background: var(--surface);
  border: 1px solid var(--border);
  padding: 8px 14px;
  border-radius: var(--radius-md);
  width: 320px;
  max-width: 100%;
}
.search__icon  { color: var(--text-muted); flex-shrink: 0; }
.search__input {
  flex: 1;
  min-width: 0;
  background: transparent;
  border: 0;
  outline: 0;
  color: var(--text);
  font-size: 13px;
}
.search__input::placeholder { color: var(--text-muted); }

.grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 16px;
}
.empty {
  padding: 60px;
  text-align: center;
  color: var(--text-muted);
  font-size: 13px;
}

@media (max-width: 1200px) { .grid { grid-template-columns: repeat(3, minmax(0, 1fr)); } }
@media (max-width: 900px)  { .grid { grid-template-columns: repeat(2, minmax(0, 1fr)); } }
@media (max-width: 560px)  { .grid { grid-template-columns: 1fr; } }

@media (max-width: 720px) {
  .page__main { padding: 20px 16px; }
  .search { width: 100%; }
}
</style>
