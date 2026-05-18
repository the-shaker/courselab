<template>
  <div class="login">
    <aside class="login__brand">
      <Logo :size="36" />

      <div class="login__center">
        <div class="login__chip">
          <span class="login__chipDot" />
          Панель администратора v3.2
        </div>
        <h1 class="login__h1">
          Управляйте<br />
          <span class="login__accent">онлайн&#8209;курсами</span><br />
          с одного места
        </h1>
        <p class="login__lead">
          Контент, студенты, статистика и поддержка — всё в едином рабочем
          пространстве для команды Курслаб.
        </p>
      </div>

      <div class="login__legal">
        2026 Курслаб
      </div>
    </aside>

    <section class="login__formPanel">
      <div class="login__formWrap">
        <header class="login__formHead">
          <h2 class="login__h2">Вход в панель</h2>
          <p class="login__sub">Введите данные администратора, чтобы продолжить.</p>
        </header>

        <form class="login__form" @submit.prevent="submit">
          <Field
            label="Email"
            icon="mail"
            placeholder="admin@kurslab.ru"
            v-model="email"
          />
          <Field
            label="Пароль"
            icon="lock"
            placeholder="Введите пароль"
            type="password"
            reveal
            v-model="password"
          />

          <div class="login__row">
            <label class="login__remember">
              <span class="login__check">
                <Icon name="check" :size="12" :stroke="3" />
              </span>
              Запомнить меня на 30 дней
            </label>
            <a class="login__forgot">Забыли пароль?</a>
          </div>

          <button type="submit" class="login__cta">
            Войти в панель
            <Icon name="arrowRight" :size="16" :stroke="2.2" />
          </button>
        </form>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import Logo from '../components/Logo.vue';
import Field from '../components/Field.vue';
import Icon from '../components/Icon.vue';

const router = useRouter();
const email = ref('nikita.yankovskiy@kurslab.ru');
const password = ref('••••••••••••');

function submit() {
  router.push('/courses');
}
</script>

<style scoped>
.login {
  display: grid;
  grid-template-columns: 1fr 1fr;
  min-height: 100vh;
  background: var(--bg);
  overflow: hidden;
}

.login__brand {
  padding: 56px 60px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  background: var(--surface);
  border-right: 1px solid var(--border);
  gap: 32px;
}

.login__center { display: flex; flex-direction: column; }

.login__chip {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  align-self: flex-start;
  padding: 6px 12px;
  border-radius: 999px;
  background: var(--accent-soft);
  color: var(--accent);
  font-size: 12px;
  font-weight: 500;
  margin-bottom: 24px;
}
.login__chipDot {
  width: 6px;
  height: 6px;
  border-radius: 999px;
  background: var(--accent);
}

.login__h1 {
  font-size: 48px;
  font-weight: 600;
  line-height: 1.1;
  margin: 0;
  letter-spacing: -1.2px;
}
.login__accent { color: var(--accent); }

.login__lead {
  font-size: 16px;
  color: var(--text-dim);
  margin-top: 20px;
  max-width: 460px;
  line-height: 1.55;
}

.login__legal { font-size: 12px; color: var(--text-muted); }
.login__legalLink { color: var(--text-dim); cursor: pointer; }
.login__legalLink:hover { color: var(--text); }

.login__formPanel {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 60px;
  background: var(--bg);
}
.login__formWrap { width: 100%; max-width: 420px; }
.login__formHead { margin-bottom: 32px; }

.login__h2 {
  font-size: 28px;
  font-weight: 600;
  letter-spacing: -0.5px;
}
.login__sub {
  font-size: 14px;
  color: var(--text-dim);
  margin-top: 8px;
}

.login__form {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.login__row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 4px;
  flex-wrap: wrap;
  gap: 8px;
}
.login__remember {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: var(--text-dim);
  cursor: pointer;
}
.login__check {
  width: 16px;
  height: 16px;
  border-radius: 4px;
  border: 1.5px solid var(--accent);
  background: var(--accent);
  color: var(--bg);
  display: inline-flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}
.login__forgot { font-size: 13px; color: var(--accent); cursor: pointer; }
.login__forgot:hover { text-decoration: underline; }

.login__cta {
  margin-top: 12px;
  height: 48px;
  background: var(--accent);
  color: var(--bg);
  font-weight: 600;
  font-size: 15px;
  border-radius: var(--radius-lg);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  box-shadow: 0 8px 20px -6px rgba(61, 220, 151, 0.4);
  transition: background .15s, transform .1s;
}
.login__cta:hover  { background: var(--accent-dim); }
.login__cta:active { transform: translateY(1px); }

@media (max-width: 960px) {
  .login { grid-template-columns: 1fr; }
  .login__brand {
    padding: 36px 28px;
    border-right: none;
    border-bottom: 1px solid var(--border);
    gap: 24px;
  }
  .login__h1 { font-size: 36px; }
  .login__formPanel { padding: 36px 28px; }
}
@media (max-width: 480px) {
  .login__brand    { padding: 28px 20px; }
  .login__h1       { font-size: 28px; letter-spacing: -0.8px; }
  .login__lead     { font-size: 14px; }
  .login__formPanel{ padding: 28px 20px; }
  .login__h2       { font-size: 22px; }
}
</style>
