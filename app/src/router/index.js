import { createRouter, createWebHistory } from 'vue-router';

import LoginView      from '../views/LoginView.vue';
import CoursesView    from '../views/CoursesView.vue';
import StatisticsView from '../views/StatisticsView.vue';
import TicketsView    from '../views/TicketsView.vue';
import UsersView      from '../views/UsersView.vue';

const routes = [
  { path: '/',           redirect: '/login' },
  { path: '/login',      name: 'login',      component: LoginView      },
  { path: '/courses',    name: 'courses',    component: CoursesView    },
  { path: '/statistics', name: 'statistics', component: StatisticsView },
  { path: '/tickets',    name: 'tickets',    component: TicketsView    },
  { path: '/users',      name: 'users',      component: UsersView      },
  { path: '/:catchAll(.*)', redirect: '/login' },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() { return { top: 0 }; },
});

export default router;
