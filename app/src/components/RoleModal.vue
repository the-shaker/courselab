<template>
  <Teleport to="body">
    <Transition name="modal">
      <div v-if="modelValue" class="overlay" @click.self="$emit('update:modelValue', false)">
        <div class="modal" role="dialog" aria-modal="true">
          <div class="modal__head">
            <div>
              <div class="modal__eyebrow">Изменение роли</div>
              <h3 class="modal__title">{{ user?.name }}</h3>
            </div>
            <button class="modal__close" @click="$emit('update:modelValue', false)">
              <Icon name="close" :size="18" />
            </button>
          </div>

          <div class="modal__body">
            <p class="modal__hint">Текущая роль: <strong>{{ user?.role }}</strong></p>

            <div class="modal__options">
              <label
                v-for="opt in roleOptions"
                :key="opt.value"
                class="modal__option"
                :class="{ 'modal__option--active': selected === opt.value }"
              >
                <input
                  type="radio"
                  :value="opt.value"
                  v-model="selected"
                  class="modal__radio"
                />
                <span class="modal__optIcon" :style="{ '--c': opt.color }">
                  <Icon :name="opt.icon" :size="16" />
                </span>
                <div>
                  <div class="modal__optName">{{ opt.label }}</div>
                  <div class="modal__optDesc">{{ opt.desc }}</div>
                </div>
                <span v-if="selected === opt.value" class="modal__check">
                  <Icon name="check" :size="12" :stroke="2.5" />
                </span>
              </label>
            </div>
          </div>

          <div class="modal__foot">
            <button class="modal__btn modal__btn--ghost" @click="$emit('update:modelValue', false)">
              Отмена
            </button>
            <button
              class="modal__btn modal__btn--primary"
              :disabled="selected === user?.roleKey"
              @click="confirm"
            >
              Применить
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { ref, watch } from 'vue';
import Icon from './Icon.vue';

const props = defineProps({
  modelValue: { type: Boolean, required: true },
  user: { type: Object, default: null },
});

const emit = defineEmits(['update:modelValue', 'confirm']);

const roleOptions = [
  {
    value: 'student',
    label: 'Студент',
    desc: 'Может проходить курсы и оставлять отзывы',
    icon: 'user',
    color: 'var(--info)',
  },
  {
    value: 'author',
    label: 'Автор курсов',
    desc: 'Может создавать и публиковать курсы',
    icon: 'course',
    color: 'var(--accent)',
  },
];

const selected = ref(null);

watch(() => props.user, (u) => {
  selected.value = u?.roleKey ?? null;
}, { immediate: true });

watch(() => props.modelValue, (v) => {
  if (v && props.user) selected.value = props.user.roleKey;
});

function confirm() {
  const opt = roleOptions.find(o => o.value === selected.value);
  emit('confirm', { user: props.user, roleKey: selected.value, roleLabel: opt?.label });
  emit('update:modelValue', false);
}
</script>

<style scoped>
.overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, .55);
  backdrop-filter: blur(3px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 20px;
}

.modal {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-xl);
  width: 100%;
  max-width: 420px;
  box-shadow: 0 24px 60px rgba(0,0,0,.4);
}

.modal__head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 12px;
  padding: 20px 22px 16px;
  border-bottom: 1px solid var(--border);
}
.modal__eyebrow {
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: .6px;
  color: var(--text-muted);
  margin-bottom: 4px;
}
.modal__title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text);
}
.modal__close {
  color: var(--text-muted);
  padding: 4px;
  border-radius: var(--radius-sm);
  flex-shrink: 0;
}
.modal__close:hover { color: var(--text); background: var(--surface-3); }

.modal__body { padding: 18px 22px; }
.modal__hint {
  font-size: 13px;
  color: var(--text-dim);
  margin-bottom: 14px;
}
.modal__hint strong { color: var(--text); }

.modal__options { display: flex; flex-direction: column; gap: 8px; }

.modal__option {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 14px;
  border: 1.5px solid var(--border);
  border-radius: var(--radius-lg);
  cursor: pointer;
  transition: border-color .15s, background .15s;
  position: relative;
}
.modal__option:hover { border-color: var(--accent); }
.modal__option--active {
  border-color: var(--accent);
  background: var(--accent-soft);
}

.modal__radio { display: none; }

.modal__optIcon {
  width: 36px;
  height: 36px;
  border-radius: var(--radius-md);
  background: color-mix(in srgb, var(--c) 14%, transparent);
  color: var(--c);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}
.modal__optName {
  font-size: 13px;
  font-weight: 600;
  color: var(--text);
}
.modal__optDesc {
  font-size: 11px;
  color: var(--text-muted);
  margin-top: 2px;
}
.modal__check {
  margin-left: auto;
  color: var(--accent);
}

.modal__foot {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  padding: 14px 22px 18px;
  border-top: 1px solid var(--border);
}

.modal__btn {
  padding: 8px 18px;
  border-radius: var(--radius-md);
  font-size: 13px;
  font-weight: 600;
}
.modal__btn--ghost {
  color: var(--text-dim);
  background: transparent;
}
.modal__btn--ghost:hover { background: var(--surface-3); color: var(--text); }
.modal__btn--primary {
  background: var(--accent);
  color: var(--bg);
}
.modal__btn--primary:hover:not(:disabled) { background: var(--accent-dim); }
.modal__btn--primary:disabled {
  opacity: .4;
  cursor: not-allowed;
}

.modal-enter-active,
.modal-leave-active { transition: opacity .2s ease; }
.modal-enter-active .modal,
.modal-leave-active .modal { transition: transform .2s ease, opacity .2s ease; }
.modal-enter-from,
.modal-leave-to { opacity: 0; }
.modal-enter-from .modal { transform: scale(.95) translateY(8px); }
.modal-leave-to .modal  { transform: scale(.95) translateY(8px); }
</style>
