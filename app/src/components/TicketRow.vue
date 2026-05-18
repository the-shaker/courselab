<template>
  <div class="row" :class="{ 'row--sel': t.sel, 'row--active': active }">
    <span class="row__sel" :class="{ 'row__sel--on': t.sel }">
      <Icon v-if="t.sel" name="check" :size="10" :stroke="3" />
    </span>

    <span class="row__id ap-mono">{{ t.id }}</span>

    <div class="row__subject">
      <div class="row__subj">{{ t.subj }}</div>
      <div class="row__user">{{ t.user }}</div>
    </div>

    <span class="row__course" data-label="Курс">{{ t.course }}</span>

    <div class="row__prio" data-label="Приоритет">
      <Pill :color="prio.c" :soft="prio.soft">{{ prio.l }}</Pill>
    </div>

    <span class="row__status" :style="{ color: stat.c }" data-label="Статус">
      <StatusDot :color="stat.dot" />{{ stat.l }}
    </span>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import Icon from './Icon.vue';
import Pill from './Pill.vue';
import StatusDot from './StatusDot.vue';

const props = defineProps({
  t:      { type: Object,  required: true },
  active: { type: Boolean, default: false },
});

const PRIO = {
  urgent: { l: 'Срочно',  c: 'var(--danger)',   soft: 'var(--danger-soft)' },
  high:   { l: 'Высокий', c: 'var(--warn)',     soft: 'var(--warn-soft)'   },
  med:    { l: 'Средний', c: 'var(--info)',     soft: 'var(--info-soft)'   },
  low:    { l: 'Низкий',  c: 'var(--text-dim)', soft: 'var(--bg)'          },
};
const STATUS = {
  open:   { l: 'Открыт', c: 'var(--accent)',     dot: 'var(--accent)'     },
  closed: { l: 'Закрыт', c: 'var(--text-muted)', dot: 'var(--text-muted)' },
};

const prio = computed(() => PRIO[props.t.prio]);
const stat = computed(() => STATUS[props.t.status]);
</script>

<style scoped>
.row {
  display: grid;
  grid-template-columns:
    24px
    80px
    minmax(0, 1fr)
    160px
    100px
    90px;
  gap: 12px;
  padding: 14px 18px;
  align-items: center;
  border-bottom: 1px solid var(--border-soft);
  position: relative;
  cursor: pointer;
  transition: background .12s;
}
.row:hover   { background: var(--surface-2, var(--surface)); }
.row--sel    { background: var(--surface-3); }
.row--active {
  background: var(--surface-3);
}
.row--active::before,
.row--sel::before {
  content: '';
  position: absolute;
  left: 0; top: 0; bottom: 0;
  width: 2px;
  background: var(--accent);
}

.row__sel {
  width: 14px; height: 14px;
  border-radius: 3px;
  border: 1.5px solid var(--border);
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: var(--bg);
}
.row__sel--on {
  border-color: var(--accent);
  background: var(--accent);
}

.row__id     { font-size: 12px; color: var(--text-dim); }
.row__subject{ min-width: 0; }
.row__subj   {
  font-size: 13px;
  font-weight: 500;
  color: var(--text);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.row__user   {
  font-size: 11px;
  color: var(--text-muted);
  margin-top: 2px;
}
.row__course {
  font-size: 12px;
  color: var(--text-dim);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.row__status {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  font-weight: 500;
}

@media (max-width: 900px) {
  .row {
    grid-template-columns: 1fr auto;
    grid-template-areas:
      'id     status'
      'subj   subj'
      'course course'
      'prio   prio';
    gap: 6px 12px;
    padding: 14px 16px;
  }
  .row__sel { display: none; }
  .row__id      { grid-area: id; }
  .row__subject { grid-area: subj; }
  .row__course  { grid-area: course; }
  .row__prio    { grid-area: prio; }
  .row__status  { grid-area: status; justify-self: end; }
  .row__course::before {
    content: attr(data-label) ': ';
    color: var(--text-muted);
  }
}
</style>
