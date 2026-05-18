<template>
  <div class="legend">
    <span class="legend__line" :class="{ 'legend__line--dashed': dashed }" :style="{ '--c': color }" />
    {{ label }}
  </div>
</template>

<script setup>
defineProps({
  color:  { type: String, default: 'var(--accent)' },
  label:  { type: String, required: true },
  dashed: { type: Boolean, default: false },
});
</script>

<style scoped>
.legend {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
  color: var(--text-dim);
  white-space: nowrap;
}
.legend__line {
  width: 18px;
  height: 0;
  border-top: 2px solid var(--c);
}
.legend__line--dashed {
  border-top-style: dashed;
}
</style>
