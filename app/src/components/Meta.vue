<template>
  <div class="meta" :class="{ 'meta--sep': sep }">
    <span class="meta__value ap-mono">{{ value }}</span>
    <span class="meta__label">{{ label }}</span>
  </div>
</template>

<script setup>
defineProps({
  value: { type: [Number, String], required: true },
  label: { type: String, required: true },
  sep:   { type: Boolean, default: false },
});
</script>

<style scoped>
.meta {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0 4px;
}
.meta--sep { border-left: 1px solid var(--border-soft); }

.meta__value {
  font-size: 14px;
  font-weight: 600;
  color: var(--text);
}
.meta__label {
  font-size: 10px;
  color: var(--text-muted);
  margin-top: 1px;
}
</style>
