<template>
  <div class="logo" :style="{ '--logo-size': size + 'px' }">
    <div class="logo__mark">
      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2"
           stroke-linecap="round" stroke-linejoin="round">
        <path d="M3 7l9-4 9 4-9 4-9-4z" />
        <path d="M7 9.5V14l5 2 5-2V9.5" />
      </svg>
    </div>
    <div v-if="withWord" class="logo__text">
      <span class="logo__name">Курслаб</span>
      <span class="logo__sub">Admin</span>
    </div>
  </div>
</template>

<script setup>
defineProps({
  size:     { type: Number, default: 28 },
  withWord: { type: Boolean, default: true },
});
</script>

<style scoped>
.logo {
  display: flex;
  align-items: center;
  gap: 10px;
  user-select: none;
}

.logo__mark {
  width:  var(--logo-size);
  height: var(--logo-size);
  border-radius: 8px;
  background: linear-gradient(135deg, var(--accent) 0%, var(--accent-dim) 100%);
  color: var(--bg);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 14px -2px rgba(61, 220, 151, 0.33);
}
.logo__mark svg {
  width:  calc(var(--logo-size) * 0.6);
  height: calc(var(--logo-size) * 0.6);
}

.logo__text {
  display: flex;
  flex-direction: column;
  line-height: 1;
}
.logo__name {
  font-weight: 700;
  font-size: 16px;
  letter-spacing: -0.3px;
}
.logo__sub {
  font-size: 10px;
  color: var(--text-muted);
  letter-spacing: 1.2px;
  text-transform: uppercase;
  margin-top: 3px;
}
</style>
