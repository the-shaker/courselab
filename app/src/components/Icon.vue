<template>
  <svg
    :width="size"
    :height="size"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    :stroke-width="stroke"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="icon"
  >
    <path :d="d" />
  </svg>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  name:   { type: String, required: true },
  size:   { type: [Number, String], default: 18 },
  stroke: { type: [Number, String], default: 1.6 },
});

const PATHS = {
  chart:       'M4 19V5M4 19h16M8 16V11M12 16V8M16 16V13',
  ticket:      'M3 9a2 2 0 012-2h14a2 2 0 012 2v2a2 2 0 100 4v2a2 2 0 01-2 2H5a2 2 0 01-2-2v-2a2 2 0 100-4V9zM12 8v2M12 13v2M12 17v0',
  users:       'M16 21v-2a4 4 0 00-4-4H6a4 4 0 00-4 4v2M9 11a4 4 0 100-8 4 4 0 000 8zM22 21v-2a4 4 0 00-3-3.87M16 3.13a4 4 0 010 7.75',
  course:      'M4 19.5A2.5 2.5 0 016.5 17H20V3H6.5A2.5 2.5 0 004 5.5v14zM4 19.5A2.5 2.5 0 006.5 22H20v-5',
  search:      'M11 19a8 8 0 100-16 8 8 0 000 16zM21 21l-4.35-4.35',
  plus:        'M12 5v14M5 12h14',
  arrowRight:  'M5 12h14M13 5l7 7-7 7',
  arrowUp:     'M12 19V5M5 12l7-7 7 7',
  arrowDown:   'M12 5v14M19 12l-7 7-7-7',
  check:       'M20 6L9 17l-5-5',
  chevronDown: 'M6 9l6 6 6-6',
  chevronRight:'M9 6l6 6-6 6',
  eye:         'M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8zM12 15a3 3 0 100-6 3 3 0 000 6z',
  lock:        'M5 11h14a2 2 0 012 2v7a2 2 0 01-2 2H5a2 2 0 01-2-2v-7a2 2 0 012-2zM7 11V7a5 5 0 0110 0v4',
  mail:        'M4 4h16a2 2 0 012 2v12a2 2 0 01-2 2H4a2 2 0 01-2-2V6a2 2 0 012-2zM2 6l10 7 10-7',
  user:        'M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2M12 11a4 4 0 100-8 4 4 0 000 8z',
  trash:       'M3 6h18M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2',
  clock:       'M12 21a9 9 0 100-18 9 9 0 000 18zM12 7v5l3 2',
  send:        'M22 2L11 13M22 2l-7 20-4-9-9-4 20-7z',
  moreV:       'M12 5h.01M12 12h.01M12 19h.01',
  menu:        'M3 6h18M3 12h18M3 18h18',
  close:       'M6 6l12 12M6 18L18 6',
};

const d = computed(() => PATHS[props.name] || PATHS.chart);
</script>

<style scoped>
.icon { display: block; flex-shrink: 0; }
</style>
