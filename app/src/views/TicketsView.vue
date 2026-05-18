<template>
  <div class="page">
    <NavBar />

    <main class="page__main">
      <header class="page__header">
        <div class="ap-eyebrow">Поддержка</div>
        <h1 class="ap-h1">
          Тикеты <span class="muted">· {{ tickets.length }} всего, {{ openCount }} открытых</span>
        </h1>
      </header>

      <div class="toolbar">
        <div class="tabs">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            class="tabs__tab"
            :class="{ 'tabs__tab--active': activeTab === tab.id }"
            @click="activeTab = tab.id"
          >
            {{ tab.label }}
            <span class="tabs__count" :class="{ 'tabs__count--active': activeTab === tab.id }">
              {{ tab.id === 'all' ? tickets.length : tab.id === 'open' ? openCount : closedCount }}
            </span>
          </button>
        </div>

        <div class="search">
          <Icon name="search" :size="15" class="search__icon" />
          <input
            v-model="searchQuery"
            class="search__input"
            placeholder="Поиск по ID, теме или e-mail..."
          />
        </div>
      </div>

      <div class="grid">
        <Card :padding="0" class="list">
          <div class="list__head">
            <span class="list__hcheck">
              <span class="list__checkbox" />
            </span>
            <span>ID</span>
            <span>Тема и пользователь</span>
            <span>Курс</span>
            <span>Приоритет</span>
            <span>Статус</span>
          </div>
          <div class="list__body ap-scroll">
            <TicketRow
              v-for="t in filteredTickets"
              :key="t.id"
              :t="t"
              :active="selectedTicket?.id === t.id"
              @click="selectedTicket = t"
            />
            <div v-if="filteredTickets.length === 0" class="list__empty">
              Тикеты не найдены
            </div>
          </div>
        </Card>

        <Card :padding="0" class="detail">
          <TicketDetail
            v-if="selectedTicket"
            :ticket="selectedTicket"
            @close="selectedTicket.status = 'closed'"
          />
          <div v-else class="detail__empty">
            <Icon name="ticket" :size="32" />
            <p>Выберите тикет из списка</p>
          </div>
        </Card>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import NavBar       from '../components/NavBar.vue';
import Card         from '../components/Card.vue';
import Icon         from '../components/Icon.vue';
import TicketRow    from '../components/TicketRow.vue';
import TicketDetail from '../components/TicketDetail.vue';
import { fetchTickets } from '../api.js';

const tickets = ref([]);

onMounted(async () => {
  const data = await fetchTickets();
  tickets.value = data.map(t => ({ ...t, sel: false }));
  selectedTicket.value = tickets.value[1] ?? null;
});


const selectedTicket = ref(null);
const activeTab      = ref('open');
const searchQuery    = ref('');

const openCount   = computed(() => tickets.value.filter(t => t.status === 'open').length);
const closedCount = computed(() => tickets.value.filter(t => t.status === 'closed').length);

const filteredTickets = computed(() => {
  const q = searchQuery.value.toLowerCase().trim();
  return tickets.value.filter(t => {
    const matchTab =
      activeTab.value === 'all' ||
      (activeTab.value === 'open'   && t.status === 'open')   ||
      (activeTab.value === 'closed' && t.status === 'closed');
    const matchSearch =
      !q ||
      t.id.toLowerCase().includes(q) ||
      t.subj.toLowerCase().includes(q) ||
      t.user.toLowerCase().includes(q) ||
      t.email.toLowerCase().includes(q);
    return matchTab && matchSearch;
  });
});

const tabs = [
  { id: 'all',    label: 'Все'       },
  { id: 'open',   label: 'Открытые'  },
  { id: 'closed', label: 'Закрытые'  },
];
</script>

<style scoped>
.page {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}
.page__main {
  flex: 1;
  padding: 24px 32px;
  display: flex;
  flex-direction: column;
  min-height: 0;
}

.page__header { margin-bottom: 18px; }

.toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 14px;
  gap: 16px;
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
.tabs__count {
  padding: 1px 7px;
  font-size: 11px;
  font-weight: 600;
  border-radius: 999px;
  background: var(--bg);
  color: var(--text-dim);
}
.tabs__count--active { background: var(--accent); color: var(--bg); }

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
  flex: 1;
  display: grid;
  grid-template-columns: 1.4fr 1fr;
  gap: 14px;
  min-height: 0;
}
.list, .detail {
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.list__head {
  display: grid;
  grid-template-columns: 24px 80px 1fr 160px 100px 90px;
  gap: 12px;
  padding: 11px 18px;
  font-size: 11px;
  font-weight: 500;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
  border-bottom: 1px solid var(--border);
  background: var(--surface);
}
.list__hcheck { display: flex; align-items: center; }
.list__checkbox {
  width: 14px; height: 14px;
  border: 1.5px solid var(--border);
  border-radius: 3px;
  display: inline-block;
}
.list__body {
  flex: 1;
  overflow-y: auto;
  min-height: 0;
}
.list__empty {
  padding: 40px;
  text-align: center;
  color: var(--text-muted);
  font-size: 13px;
}

.detail__empty {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 12px;
  color: var(--text-muted);
  font-size: 13px;
  height: 100%;
}

@media (max-width: 1100px) {
  .grid { grid-template-columns: 1fr; }
  .detail { min-height: 520px; }
}
@media (max-width: 900px) {
  .list__head { display: none; }
  .page__main { padding: 18px 16px; }
  .search { width: 100%; }
}
</style>
