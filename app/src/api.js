const BASE = 'https://api.nikitos.pro';

const get = (path) => fetch(BASE + path).then(r => r.json());

export const fetchCourses    = () => get('/courses');
export const fetchUsers      = () => get('/users');
export const fetchTickets    = () => get('/tickets');
export const fetchKPI        = () => get('/stats/kpi');
export const fetchUserStats  = () => get('/stats/users');
export const fetchChart      = () => get('/stats/chart');
