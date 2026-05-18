<template>
  <span class="pill" :style="cssVars">
    <slot />
  </span>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  color: { type: String, default: 'var(--accent)' },
  soft:  { type: String, default: 'var(--accent-soft)' },
});

const cssVars = computed(() => ({
  '--pill-c':    props.color,
  '--pill-soft': props.soft,
}));
</script>

<style scoped>
.pill {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  font-weight: 500;
  padding: 3px 9px;
  border-radius: 999px;
  background: var(--pill-soft);
  color: var(--pill-c);
  white-space: nowrap;
}
</style>
