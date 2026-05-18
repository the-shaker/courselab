<template>
  <div class="card" :style="{ '--card-pad': pad + 'px' }">
    <slot />
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  padding: { type: [Number, String], default: 20 },
});

const pad = computed(() => props.padding);
</script>

<style scoped>
.card {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-xl);
  padding: var(--card-pad);
}
</style>
