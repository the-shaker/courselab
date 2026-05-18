<template>
  <article class="cc">
    <!-- Cover -->
    <div class="cc__cover" :style="{ '--c': color }">
      <span class="cc__tag" :style="{ '--tc': tagC }">{{ tag }}</span>
      <span class="cc__status" :style="statusVars">
        <span class="cc__statusDot" />
        {{ status.l }}
      </span>
    </div>

    <!-- Body -->
    <div class="cc__body">
      <div>
        <h3 class="cc__title">{{ title }}</h3>
        <div class="cc__author">
          <div class="cc__avatar" :style="{ '--c': color }">{{ authorInit }}</div>
          <span>{{ author }}</span>
        </div>
      </div>

      <!-- Meta strip -->
      <div class="cc__meta">
        <Meta :value="lessons" label="уроков" />
        <Meta :value="duration" label="время" sep />
        <Meta :value="students > 0 ? students.toLocaleString('ru-RU') : '—'" label="студ." sep />
      </div>

      <!-- Progress or rating -->
      <div v-if="progress < 100">
        <div class="cc__progressHead">
          <span class="cc__progressLbl">Готовность</span>
          <span class="cc__progressVal ap-mono">{{ progress }}%</span>
        </div>
        <div class="cc__progressTrack">
          <div
            class="cc__progressBar"
            :style="{
              width: progress + '%',
              background: progress < 50 ? 'var(--warn)' : 'var(--accent)',
            }"
          />
        </div>
      </div>
      <div v-else-if="rating" class="cc__rating">
        <svg width="13" height="13" viewBox="0 0 24 24" fill="var(--warn)" aria-hidden="true">
          <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z" />
        </svg>
        <span class="ap-mono cc__rate">{{ rating }}</span>
        <span class="cc__reviews">· {{ students.toLocaleString('ru-RU') }} отзывов</span>
      </div>

      <!-- Footer — updated timestamp -->
      <div class="cc__footer">
        <span class="cc__updated">
          <Icon name="clock" :size="11" /> {{ updated }}
        </span>
      </div>
    </div>
  </article>
</template>

<script setup>
import { computed } from 'vue';
import Meta from './Meta.vue';
import Icon from './Icon.vue';

const props = defineProps({
  title:      { type: String,  required: true },
  author:     { type: String,  required: true },
  authorInit: { type: String,  required: true },
  color:      { type: String,  required: true },
  students:   { type: Number,  default: 0 },
  lessons:    { type: Number,  required: true },
  duration:   { type: String,  required: true },
  status:     { type: Object,  required: true }, // { l, c, s }
  updated:    { type: String,  required: true },
  tag:        { type: String,  required: true },
  tagC:       { type: String,  required: true },
  rating:     { type: [Number, null], default: null },
  progress:   { type: Number,  default: 100 },
});

const statusVars = computed(() => ({
  background: props.status.s,
  color:      props.status.c,
  '--sc':     props.status.c,
}));
</script>

<style scoped>
.cc {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-xl);
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* Cover */
.cc__cover {
  height: 110px;
  background: linear-gradient(135deg,
    color-mix(in srgb, var(--c) 18%, var(--surface-2)) 0%,
    color-mix(in srgb, var(--c) 4%, var(--surface-2)) 45%,
    var(--surface-2) 100%);
  position: relative;
  border-bottom: 1px solid var(--border);
}
.cc__tag, .cc__status {
  position: absolute;
  top: 12px;
  font-size: 11px;
  font-weight: 500;
  padding: 3px 9px;
  border-radius: 999px;
  backdrop-filter: blur(8px);
}
.cc__tag {
  left: 14px;
  background: color-mix(in srgb, var(--tc) 15%, transparent);
  color: var(--tc);
}
.cc__status {
  right: 14px;
  display: inline-flex;
  align-items: center;
  gap: 5px;
}
.cc__statusDot {
  width: 5px;
  height: 5px;
  border-radius: 999px;
  background: var(--sc);
}

/* Body */
.cc__body {
  padding: 14px 16px 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  flex: 1;
}

.cc__title {
  font-size: 14px;
  font-weight: 600;
  color: var(--text);
  line-height: 1.35;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  min-height: 38px;
}

.cc__author {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 8px;
  font-size: 12px;
  color: var(--text-dim);
}
.cc__avatar {
  width: 22px;
  height: 22px;
  border-radius: 999px;
  background: linear-gradient(135deg, var(--c), color-mix(in srgb, var(--c) 70%, transparent));
  color: var(--bg);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 10px;
  font-weight: 700;
  flex-shrink: 0;
}

.cc__meta {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  background: var(--bg);
  border-radius: var(--radius-md);
  padding: 10px 4px;
  border: 1px solid var(--border-soft);
}

.cc__progressHead {
  display: flex;
  justify-content: space-between;
  margin-bottom: 6px;
}
.cc__progressLbl { font-size: 11px; color: var(--text-muted); }
.cc__progressVal { font-size: 11px; color: var(--text); font-weight: 600; }
.cc__progressTrack {
  height: 4px;
  background: var(--bg);
  border-radius: 2px;
  overflow: hidden;
}
.cc__progressBar {
  height: 100%;
  border-radius: 2px;
  transition: width .25s ease;
}

.cc__rating {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
}
.cc__rate    { color: var(--text); font-weight: 600; }
.cc__reviews { color: var(--text-muted); }

.cc__footer {
  margin-top: auto;
  padding-top: 10px;
  border-top: 1px solid var(--border-soft);
  display: flex;
  align-items: center;
}
.cc__updated {
  font-size: 11px;
  color: var(--text-muted);
  display: flex;
  align-items: center;
  gap: 5px;
}
</style>
