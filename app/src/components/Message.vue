<template>
  <div class="msg" :class="{ 'msg--you': you }">
    <div class="msg__from">{{ from }}</div>
    <div class="msg__bubble">
      <slot>{{ text }}</slot>
    </div>
  </div>
</template>

<script setup>
defineProps({
  from: { type: String, required: true },
  text: { type: String, default: '' },
  you:  { type: Boolean, default: false },
});
</script>

<style scoped>
.msg {
  display: flex;
  flex-direction: column;
  gap: 6px;
  align-items: flex-start;
}
.msg--you { align-items: flex-end; }

.msg__from {
  font-size: 11px;
  color: var(--text-dim);
}

.msg__bubble {
  max-width: 85%;
  padding: 10px 14px;
  border-radius: 12px;
  border-top-left-radius: 4px;
  background: var(--bg);
  border: 1px solid var(--border);
  color: var(--text);
  font-size: 13px;
  line-height: 1.5;
}
.msg--you .msg__bubble {
  background: var(--accent-soft);
  border-color: rgba(61, 220, 151, 0.2);
  border-top-left-radius: 12px;
  border-top-right-radius: 4px;
}
</style>
