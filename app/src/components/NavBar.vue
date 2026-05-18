<template>
  <header class="nav">
    <Logo />

    <button class="nav__burger" @click="open = !open" :aria-label="open ? 'Закрыть меню' : 'Открыть меню'">
      <Icon :name="open ? 'close' : 'menu'" :size="20" />
    </button>

    <nav class="nav__items" :class="{ 'nav__items--open': open }">
      <router-link
        v-for="item in items"
        :key="item.id"
        :to="item.to"
        class="nav__item"
        active-class="nav__item--active"
        @click="open = false"
      >
        <Icon :name="item.icon" :size="16" />
        <span>{{ item.label }}</span>
        <span v-if="item.badge" class="nav__badge ap-mono">{{ item.badge }}</span>
      </router-link>
    </nav>

    <div class="nav__user">
      <div class="nav__avatar">Н</div>
      <div class="nav__userText">
        <div class="nav__name">Никита</div>
        <div class="nav__role">Администратор</div>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref } from 'vue';
import Logo from './Logo.vue';
import Icon from './Icon.vue';

const open = ref(false);

const items = [
  { id: 'courses',    label: 'Курсы',        icon: 'course', to: '/courses' },
  { id: 'statistics', label: 'Статистика',   icon: 'chart',  to: '/statistics' },
  { id: 'tickets',    label: 'Тикеты',       icon: 'ticket', to: '/tickets', badge: 8 },
  { id: 'users',      label: 'Пользователи', icon: 'users',  to: '/users' },
];
</script>

<style scoped>
.nav {
  height: var(--nav-h);
  border-bottom: 1px solid var(--border);
  background: var(--surface);
  display: flex;
  align-items: center;
  gap: 24px;
  padding: 0 28px;
  position: sticky;
  top: 0;
  z-index: 10;
}

.nav__items {
  display: flex;
  align-items: center;
  gap: 4px;
  flex: 1;
  margin-left: 12px;
}

.nav__item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 14px;
  border-radius: var(--radius-md);
  color: var(--text-dim);
  font-size: 14px;
  font-weight: 500;
  position: relative;
  transition: background .15s, color .15s;
  white-space: nowrap;
}
.nav__item:hover { color: var(--text); }
.nav__item--active {
  color: var(--text);
  background: var(--surface-3);
}
.nav__item--active::after {
  content: '';
  position: absolute;
  left: 14px; right: 14px;
  bottom: -17px;
  height: 2px;
  background: var(--accent);
  border-radius: 2px;
}

.nav__badge {
  min-width: 20px;
  height: 18px;
  padding: 0 6px;
  border-radius: 999px;
  background: var(--accent);
  color: var(--bg);
  font-size: 11px;
  font-weight: 700;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  margin-left: 4px;
}

.nav__user {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 4px 12px 4px 4px;
  border-radius: 999px;
  background: var(--bg);
  border: 1px solid var(--border);
}
.nav__avatar {
  width: 30px;
  height: 30px;
  border-radius: 999px;
  background: linear-gradient(135deg, var(--accent), var(--accent-dim));
  color: var(--bg);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 700;
}
.nav__userText { line-height: 1.2; }
.nav__name { font-size: 13px; font-weight: 600; }
.nav__role { font-size: 11px; color: var(--text-muted); }
.nav__chev { color: var(--text-muted); }

.nav__burger {
  display: none;
  width: 38px;
  height: 38px;
  align-items: center;
  justify-content: center;
  border-radius: var(--radius-md);
  background: var(--bg);
  border: 1px solid var(--border);
  color: var(--text-dim);
  margin-left: auto;
}

@media (max-width: 960px) {
  .nav { padding: 0 18px; gap: 16px; }
  .nav__item { padding: 8px 10px; font-size: 13px; }
  .nav__item--active::after { display: none; }
  .nav__userText { display: none; }
  .nav__chev    { display: none; }
}

@media (max-width: 720px) {
  .nav {
    padding: 0 16px;
    gap: 12px;
    flex-wrap: wrap;
    height: auto;
    min-height: var(--nav-h);
  }

  .nav__burger { display: flex; order: 2; }
  .nav__user   { order: 3; padding: 4px; }
  .nav__avatar { width: 28px; height: 28px; font-size: 11px; }

  .nav__items {
    order: 4;
    width: 100%;
    flex-direction: column;
    align-items: stretch;
    gap: 4px;
    margin: 0;
    padding: 0 0 12px;
    border-top: 1px solid var(--border);
    margin-top: 8px;
    display: none;
  }
  .nav__items--open { display: flex; }

  .nav__item {
    padding: 12px 14px;
    font-size: 14px;
    border-radius: var(--radius-md);
  }
  .nav__item--active { background: var(--surface-3); }
}
</style>
