<template>
  <label class="field">
    <span class="field__label">{{ label }}</span>
    <div class="field__box">
      <Icon v-if="icon" :name="icon" :size="16" class="field__icon" />
      <input
        :type="type"
        :placeholder="placeholder"
        :value="modelValue"
        @input="$emit('update:modelValue', $event.target.value)"
        class="field__input"
      />
      <button v-if="reveal" type="button" class="field__trailing">
        <Icon name="eye" :size="16" />
      </button>
    </div>
  </label>
</template>

<script setup>
import Icon from './Icon.vue';

defineProps({
  label:       { type: String, required: true },
  icon:        { type: String, default: '' },
  placeholder: { type: String, default: '' },
  modelValue:  { type: String, default: '' },
  type:        { type: String, default: 'text' },
  reveal:      { type: Boolean, default: false },
});

defineEmits(['update:modelValue']);
</script>

<style scoped>
.field {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.field__label {
  font-size: 12px;
  color: var(--text-dim);
  font-weight: 500;
}

.field__box {
  display: flex;
  align-items: center;
  gap: 10px;
  height: 46px;
  padding: 0 14px;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  transition: border-color .15s, background .15s;
}
.field__box:focus-within {
  border-color: var(--accent);
  background: var(--surface-2);
}

.field__icon { color: var(--text-muted); }

.field__input {
  flex: 1;
  min-width: 0;
  background: transparent;
  border: 0;
  outline: 0;
  color: var(--text);
  font-size: 14px;
}
.field__input::placeholder { color: var(--text-muted); }

.field__trailing {
  color: var(--text-muted);
  padding: 4px;
  display: flex;
  align-items: center;
  border-radius: 4px;
}
.field__trailing:hover { color: var(--text-dim); }
</style>
