<template>
  <div class="td">
    <div class="td__head">
      <span class="td__id ap-mono">{{ ticket.id }}</span>
      <h3 class="td__title">{{ ticket.subj }}</h3>
      <div class="td__meta">
        <Pill :color="prio.c" :soft="prio.soft">{{ prio.l }}</Pill>
        <span class="td__status" :style="{ color: stat.c }">
          <StatusDot :color="stat.dot" /> {{ stat.l }}
        </span>
        <button
          v-if="ticket.status === 'open'"
          class="td__closeBtn"
          @click="$emit('close')"
        >
          Закрыть тикет
        </button>
      </div>
    </div>

    <div class="td__user">
      <div class="td__avatar">{{ ticket.avatar }}</div>
      <div>
        <div class="td__name">{{ ticket.user }}</div>
        <div class="td__sub">{{ ticket.email }} · Курс «{{ ticket.course }}»</div>
      </div>
    </div>

    <div ref="convoEl" class="td__convo ap-scroll">
      <Message
        v-for="(m, i) in ticket.messages"
        :key="i"
        :from="m.from"
        :text="m.text"
        :you="m.you"
      />
    </div>

    <div v-if="ticket.status === 'open'" class="td__reply">
      <div class="td__replyBox" :class="{ 'td__replyBox--focus': focused }">
        <textarea
          ref="textareaEl"
          v-model="replyText"
          class="td__replyInput"
          placeholder="Написать ответ..."
          rows="2"
          @focus="focused = true"
          @blur="focused = false"
          @keydown.enter.exact.prevent="send"
        />
      </div>
      <div class="td__replyActions">
        <button
          type="button"
          class="td__send"
          :disabled="!replyText.trim()"
          @click="send"
        >
          <Icon name="send" :size="14" />
          Отправить
        </button>
      </div>
    </div>

    <div v-else class="td__closed">
      <StatusDot color="var(--text-muted)" />
      Тикет закрыт
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, nextTick } from 'vue';
import Icon      from './Icon.vue';
import Pill      from './Pill.vue';
import StatusDot from './StatusDot.vue';
import Message   from './Message.vue';

const props = defineProps({
  ticket: { type: Object, required: true },
});

defineEmits(['close']);

const PRIO = {
  urgent: { l: 'Срочно',  c: 'var(--danger)',   soft: 'var(--danger-soft)' },
  high:   { l: 'Высокий', c: 'var(--warn)',     soft: 'var(--warn-soft)'   },
  med:    { l: 'Средний', c: 'var(--info)',     soft: 'var(--info-soft)'   },
  low:    { l: 'Низкий',  c: 'var(--text-dim)', soft: 'var(--bg)'          },
};
const STATUS = {
  open:   { l: 'Открыт', c: 'var(--accent)',     dot: 'var(--accent)'     },
  closed: { l: 'Закрыт', c: 'var(--text-muted)', dot: 'var(--text-muted)' },
};

const prio = computed(() => PRIO[props.ticket.prio]);
const stat = computed(() => STATUS[props.ticket.status]);

const replyText = ref('');
const focused   = ref(false);
const convoEl   = ref(null);
const textareaEl = ref(null);

function scrollToBottom() {
  nextTick(() => {
    if (convoEl.value) convoEl.value.scrollTop = convoEl.value.scrollHeight;
  });
}

watch(() => props.ticket, () => {
  replyText.value = '';
  scrollToBottom();
}, { immediate: true });

function send() {
  const text = replyText.value.trim();
  if (!text) return;
  props.ticket.messages.push({ from: 'Поддержка', you: true, text });
  replyText.value = '';
  scrollToBottom();
}
</script>

<style scoped>
.td {
  display: flex;
  flex-direction: column;
  height: 100%;
  min-height: 0;
}

.td__head {
  padding: 16px 22px;
  border-bottom: 1px solid var(--border);
  flex-shrink: 0;
}
.td__id    { font-size: 12px; color: var(--text-muted); }
.td__title {
  font-size: 17px;
  font-weight: 600;
  margin: 8px 0 12px;
}
.td__meta {
  display: flex;
  gap: 10px;
  align-items: center;
  flex-wrap: wrap;
}
.td__status {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  font-weight: 500;
}
.td__closeBtn {
  margin-left: auto;
  font-size: 11px;
  color: var(--text-muted);
  padding: 3px 8px;
  border-radius: var(--radius-sm);
  border: 1px solid var(--border);
}
.td__closeBtn:hover { color: var(--text); border-color: var(--text-muted); }

.td__user {
  padding: 14px 22px;
  border-bottom: 1px solid var(--border);
  display: flex;
  gap: 12px;
  align-items: center;
  flex-shrink: 0;
}
.td__avatar {
  width: 38px; height: 38px;
  border-radius: 999px;
  background: linear-gradient(135deg, var(--accent), var(--accent-dim));
  color: var(--bg);
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 13px;
  flex-shrink: 0;
}
.td__name { font-size: 13px; font-weight: 600; }
.td__sub  { font-size: 11px; color: var(--text-muted); }

.td__convo {
  flex: 1;
  min-height: 0;
  padding: 16px 22px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.td__reply {
  padding: 12px 16px;
  border-top: 1px solid var(--border);
  background: var(--bg);
  flex-shrink: 0;
}
.td__replyBox {
  padding: 10px 12px;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  transition: border-color .15s;
}
.td__replyBox--focus { border-color: var(--accent); }

.td__replyInput {
  width: 100%;
  background: transparent;
  border: 0;
  outline: 0;
  color: var(--text);
  font-size: 13px;
  resize: none;
  font-family: inherit;
  line-height: 1.5;
}
.td__replyInput::placeholder { color: var(--text-muted); }

.td__replyActions {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 8px;
}
.td__hint {
  font-size: 11px;
  color: var(--text-muted);
}
.td__send {
  margin-left: auto;
  padding: 8px 14px;
  font-size: 13px;
  font-weight: 600;
  border-radius: var(--radius-md);
  background: var(--accent);
  color: var(--bg);
  display: flex;
  align-items: center;
  gap: 8px;
  transition: background .15s;
}
.td__send:hover:not(:disabled) { background: var(--accent-dim); }
.td__send:disabled { opacity: .4; cursor: not-allowed; }

.td__closed {
  padding: 14px 22px;
  border-top: 1px solid var(--border);
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
  color: var(--text-muted);
  flex-shrink: 0;
}
</style>
