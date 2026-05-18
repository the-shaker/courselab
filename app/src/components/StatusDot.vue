<template>
  <span class="dot" :style="{ '--dot-c': color }" />
</template>

<script setup>
defineProps({
  color: { type: String, default: 'var(--accent)' },
});
</script>

<style scoped>
.dot {
  display: inline-block;
  width: 7px;
  height: 7px;
  border-radius: 999px;
  background: var(--dot-c);
}
</style>
