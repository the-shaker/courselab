<template>
  <div class="page">
    <NavBar />

    <main class="page__main">
      <header class="page__header">
        <div class="ap-eyebrow">База пользователей</div>
        <h1 class="ap-h1">
          Пользователи <span class="muted">· {{ users.length }} всего</span>
        </h1>
      </header>

      <section class="stats">
        <div v-for="s in stats" :key="s.label" class="stats__cell">
          <span class="stats__icon" :style="{ '--c': s.color }">
            <Icon name="user" :size="16" />
          </span>
          <div>
            <div class="stats__value ap-mono">{{ s.value }}</div>
            <div class="stats__label">{{ s.label }}</div>
          </div>
        </div>
      </section>

      <div class="toolbar">
        <div class="tabs">
          <button
            v-for="r in roles"
            :key="r.key"
            class="tabs__tab"
            :class="{ 'tabs__tab--active': activeRole === r.key }"
            @click="activeRole = r.key"
          >
            {{ r.l }}
            <span class="tabs__count ap-mono">{{ r.c }}</span>
          </button>
        </div>

        <div class="search">
          <Icon name="search" :size="14" class="search__icon" />
          <input
            v-model="searchQuery"
            class="search__input"
            placeholder="Имя, email или ID..."
          />
        </div>
      </div>

      <Transition name="bulk-bar">
        <div v-if="selectedUsers.length" class="bulk">
          <span class="bulk__text">
            <span class="bulk__count ap-mono">{{ selectedUsers.length }}</span>
            {{ selectedUsers.length === 1 ? 'пользователь выбран' : 'пользователя выбрано' }}
          </span>
          <span class="bulk__sep" />
          <button class="bulk__action" @click="openRoleModal(selectedUsers[0])">
            <Icon name="user" :size="13" /> Изменить роль
          </button>
          <button class="bulk__action bulk__action--danger" @click="toggleBanSelected">
            <Icon name="trash" :size="13" />
            {{ selectedUsers.every(u => u.status === 'banned') ? 'Разблокировать' : 'Заблокировать' }}
          </button>
        </div>
      </Transition>

      <Card :padding="0" class="table">
        <div class="table__head">
          <span class="table__hcheck">
            <span
              class="table__checkbox"
              :class="{ 'table__checkbox--checked': allSelected && filteredUsers.length > 0 }"
              @click="toggleAll"
            >
              <Icon v-if="allSelected && filteredUsers.length" name="check" :size="9" :stroke="2.5" />
            </span>
          </span>
          <span>Пользователь</span>
          <span>Роль</span>
          <span>Регистрация</span>
          <span>Был онлайн</span>
          <span>Статус</span>
          <span />
        </div>
        <div class="table__body ap-scroll">
          <UserRow
            v-for="u in filteredUsers"
            :key="u.email"
            :u="u"
            @toggle-select="toggleSelect(u)"
            @change-role="openRoleModal(u)"
            @toggle-ban="toggleBan(u)"
          />
          <div v-if="filteredUsers.length === 0" class="table__empty">
            Пользователи не найдены
          </div>
        </div>
      </Card>
    </main>

    <RoleModal
      v-model="roleModalOpen"
      :user="roleModalUser"
      @confirm="applyRole"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import NavBar     from '../components/NavBar.vue';
import Card       from '../components/Card.vue';
import Icon       from '../components/Icon.vue';
import UserRow    from '../components/UserRow.vue';
import RoleModal  from '../components/RoleModal.vue';
import { fetchUsers, fetchUserStats } from '../api.js';

const users = ref([]);
const stats = ref([]);

onMounted(async () => {
  const [rawUsers, rawStats] = await Promise.all([fetchUsers(), fetchUserStats()]);
  users.value = rawUsers.map(u => ({ ...u, sel: false }));
  stats.value = rawStats;
});

const roles = [
  { key: 'all',     l: 'Все',      c: computed(() => users.value.length) },
  { key: 'student', l: 'Студенты', c: computed(() => users.value.filter(u => u.roleKey === 'student').length) },
  { key: 'author',  l: 'Авторы',   c: computed(() => users.value.filter(u => u.roleKey === 'author').length)  },
];

const activeRole  = ref('all');
const searchQuery = ref('');

const filteredUsers = computed(() => {
  const q = searchQuery.value.toLowerCase().trim();
  return users.value.filter(u => {
    const matchRole = activeRole.value === 'all' || u.roleKey === activeRole.value;
    const matchSearch = !q || u.name.toLowerCase().includes(q) || u.email.toLowerCase().includes(q);
    return matchRole && matchSearch;
  });
});

const selectedUsers = computed(() => filteredUsers.value.filter(u => u.sel));
const allSelected   = computed(() =>
  filteredUsers.value.length > 0 && filteredUsers.value.every(u => u.sel),
);

function toggleSelect(user) {
  user.sel = !user.sel;
}

function toggleAll() {
  const next = !allSelected.value;
  filteredUsers.value.forEach(u => { u.sel = next; });
}

function toggleBan(user) {
  user.status = user.status === 'banned' ? 'active' : 'banned';
}

function toggleBanSelected() {
  const allBanned = selectedUsers.value.every(u => u.status === 'banned');
  selectedUsers.value.forEach(u => { u.status = allBanned ? 'active' : 'banned'; });
}

const roleModalOpen = ref(false);
const roleModalUser = ref(null);

function openRoleModal(user) {
  roleModalUser.value = user;
  roleModalOpen.value = true;
}

function applyRole({ user, roleKey, roleLabel }) {
  const target = users.value.find(u => u.email === user.email);
  if (!target) return;
  target.roleKey = roleKey;
  target.role    = roleLabel;
  target.roleC   = roleKey === 'author' ? 'var(--accent)' : 'var(--info)';
}
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

.page__header { margin-bottom: 20px; }

.stats {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 12px;
  margin-bottom: 18px;
}
.stats__cell {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  padding: 14px 16px;
  display: flex;
  align-items: center;
  gap: 12px;
}
.stats__icon {
  width: 36px;
  height: 36px;
  border-radius: var(--radius-md);
  background: color-mix(in srgb, var(--c) 12%, transparent);
  color: var(--c);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}
.stats__value {
  font-size: 20px;
  font-weight: 600;
  letter-spacing: -0.4px;
}
.stats__label {
  font-size: 11px;
  color: var(--text-muted);
  margin-top: 1px;
}

.toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 14px;
  margin-bottom: 14px;
  flex-wrap: wrap;
}
.tabs {
  display: flex;
  gap: 4px;
  padding: 4px;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
}
.tabs__tab {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 7px 13px;
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
  padding: 8px 12px;
  border-radius: var(--radius-md);
  width: 280px;
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

.bulk {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 16px;
  margin-bottom: 12px;
  background: var(--accent-soft);
  border: 1px solid var(--accent-bd);
  border-radius: var(--radius-lg);
  flex-wrap: wrap;
}
.bulk__text  { font-size: 13px; color: var(--text); }
.bulk__count { color: var(--accent); font-weight: 600; }
.bulk__sep   { width: 1px; height: 18px; background: var(--border); }
.bulk__action {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: var(--text-dim);
  padding: 4px 8px;
  border-radius: var(--radius-sm);
  cursor: pointer;
}
.bulk__action:hover         { color: var(--text); background: rgba(255,255,255,.04); }
.bulk__action--danger       { color: var(--danger); margin-left: auto; }
.bulk__action--danger:hover { background: var(--danger-soft); }

.table {
  flex: 1;
  display: flex;
  flex-direction: column;
  min-height: 0;
  overflow: hidden;
}
.table__head {
  display: grid;
  grid-template-columns:
    40px
    minmax(0, 2.4fr) minmax(0, 1.4fr)
    140px 110px 100px 36px;
  padding: 12px 18px;
  font-size: 11px;
  font-weight: 500;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
  border-bottom: 1px solid var(--border);
  background: var(--surface);
  align-items: center;
  flex-shrink: 0;
}
.table__hcheck { display: flex; align-items: center; }
.table__checkbox {
  width: 16px; height: 16px;
  border: 1.5px solid var(--border);
  border-radius: 3px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: var(--bg);
}
.table__checkbox--checked {
  border-color: var(--accent);
  background: var(--accent);
}

.table__body {
  flex: 1;
  overflow-y: auto;
  overflow-x: hidden;
  min-height: 0;
}

.table__empty {
  padding: 40px;
  text-align: center;
  color: var(--text-muted);
  font-size: 13px;
}

.bulk-bar-enter-active,
.bulk-bar-leave-active { transition: all .2s ease; }
.bulk-bar-enter-from,
.bulk-bar-leave-to { opacity: 0; transform: translateY(-6px); margin-bottom: 0; padding-top: 0; padding-bottom: 0; }

@media (max-width: 1100px) {
  .stats { grid-template-columns: repeat(2, minmax(0, 1fr)); }
}
@media (max-width: 900px) {
  .table__head { display: none; }
  .page__main  { padding: 20px 16px; }
  .search      { width: 100%; }
}
@media (max-width: 480px) {
  .stats { grid-template-columns: 1fr; }
}
</style>
