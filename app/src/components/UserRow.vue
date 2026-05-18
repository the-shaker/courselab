<template>
  <div class="row" :class="{ 'row--sel': u.sel }" @click="$emit('toggle-select')">
    <span class="row__sel" :class="{ 'row__sel--on': u.sel }">
      <Icon v-if="u.sel" name="check" :size="11" :stroke="3" />
    </span>

    <div class="row__user">
      <div class="row__avatar" :style="{ '--c': u.color }">{{ u.avatar }}</div>
      <div class="row__userText">
        <div class="row__name">{{ u.name }}</div>
        <div class="row__email">{{ u.email }}</div>
      </div>
    </div>

    <div class="row__cell row__cell--role" data-label="Роль">
      <Pill :color="u.roleC" :soft="`color-mix(in srgb, ${u.roleC} 10%, transparent)`">{{ u.role }}</Pill>
    </div>

    <div class="row__cell row__cell--reg" data-label="Регистрация">{{ u.registered }}</div>

    <div class="row__cell row__cell--seen" data-label="Был онлайн">{{ u.lastSeen }}</div>

    <div class="row__cell row__cell--status" data-label="Статус">
      <span class="row__status" :style="{ color: statusColor }">
        <StatusDot :color="statusColor" />{{ statusLabel }}
      </span>
    </div>

    <div class="row__more-wrap" @click.stop>
      <button class="row__more" aria-label="Действия" @click="menuOpen = !menuOpen">
        <Icon name="moreV" :size="16" />
      </button>
      <div v-if="menuOpen" class="row__menu" v-click-outside="() => menuOpen = false">
        <button class="row__menuItem" @click="menuOpen = false; $emit('change-role')">
          <Icon name="user" :size="13" /> Изменить роль
        </button>
        <button class="row__menuItem row__menuItem--danger" @click="menuOpen = false; $emit('toggle-ban')">
          <Icon name="trash" :size="13" />
          {{ u.status === 'banned' ? 'Разблокировать' : 'Заблокировать' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import Icon from './Icon.vue';
import Pill from './Pill.vue';
import StatusDot from './StatusDot.vue';

const props = defineProps({
  u: { type: Object, required: true },
});

defineEmits(['toggle-select', 'change-role', 'toggle-ban']);

const menuOpen = ref(false);

const vClickOutside = {
  mounted(el, binding) {
    el._handler = (e) => { if (!el.contains(e.target)) binding.value(e); };
    document.addEventListener('click', el._handler, true);
  },
  unmounted(el) {
    document.removeEventListener('click', el._handler, true);
  },
};

const statusColor = computed(() => ({
  active:  'var(--accent)',
  pending: 'var(--warn)',
  banned:  'var(--danger)',
}[props.u.status]));

const statusLabel = computed(() => ({
  active:  'Активен',
  pending: 'Ожидает',
  banned:  'Заблокирован',
}[props.u.status]));
</script>

<style scoped>
.row {
  display: grid;
  grid-template-columns:
    40px
    minmax(0, 2.4fr)
    minmax(0, 1.4fr)
    140px
    110px
    100px
    36px;
  padding: 12px 18px;
  align-items: center;
  border-bottom: 1px solid var(--border-soft);
  background: transparent;
  position: relative;
  cursor: pointer;
  transition: background .12s;
}
.row:hover { background: var(--surface-2, var(--surface)); }
.row--sel  { background: var(--surface-3); }
.row--sel::before {
  content: '';
  position: absolute;
  left: 0; top: 0; bottom: 0;
  width: 2px;
  background: var(--accent);
}

.row__sel {
  width: 16px; height: 16px;
  border-radius: 3px;
  border: 1.5px solid var(--border);
  background: transparent;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: var(--bg);
  flex-shrink: 0;
}
.row__sel--on {
  border-color: var(--accent);
  background: var(--accent);
}

.row__user {
  display: flex;
  align-items: center;
  gap: 12px;
  min-width: 0;
}
.row__avatar {
  width: 36px;
  height: 36px;
  border-radius: 999px;
  background: linear-gradient(135deg, var(--c), color-mix(in srgb, var(--c) 70%, transparent));
  color: var(--bg);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 700;
  flex-shrink: 0;
}
.row__userText { min-width: 0; }
.row__name {
  font-size: 13px;
  font-weight: 500;
  color: var(--text);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.row__email {
  font-size: 11px;
  color: var(--text-muted);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.row__cell { font-size: 12px; color: var(--text-dim); }

.row__status {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  font-weight: 500;
}

.row__more-wrap {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}
.row__more {
  color: var(--text-muted);
  padding: 4px;
  display: flex;
  align-items: center;
  border-radius: var(--radius-sm);
}
.row__more:hover { color: var(--text-dim); background: var(--surface-3); }

.row__menu {
  position: absolute;
  right: 0;
  top: calc(100% + 4px);
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  box-shadow: 0 8px 24px rgba(0,0,0,.3);
  min-width: 170px;
  z-index: 100;
  overflow: hidden;
}
.row__menuItem {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
  padding: 9px 14px;
  font-size: 13px;
  color: var(--text-dim);
  text-align: left;
}
.row__menuItem:hover { background: var(--surface-3); color: var(--text); }
.row__menuItem--danger { color: var(--danger); }
.row__menuItem--danger:hover { background: var(--danger-soft); }

@media (max-width: 900px) {
  .row {
    grid-template-columns: 1fr auto;
    grid-template-areas:
      'user   more'
      'role   role'
      'reg    seen'
      'status status';
    gap: 8px 12px;
    padding: 14px 16px;
  }
  .row__sel { display: none; }
  .row__user         { grid-area: user; }
  .row__more-wrap    { grid-area: more; align-self: start; justify-self: end; }
  .row__cell--role   { grid-area: role; }
  .row__cell--reg    { grid-area: reg; }
  .row__cell--seen   { grid-area: seen; text-align: right; }
  .row__cell--status { grid-area: status; text-align: right; }
  .row__cell--reg,
  .row__cell--seen { font-size: 11px; }
  .row__cell--reg::before,
  .row__cell--seen::before,
  .row__cell--status::before {
    content: attr(data-label) ': ';
    color: var(--text-muted);
    margin-right: 4px;
  }
  .row__cell--status::before { content: ''; }
}
</style>
