package main

import "github.com/gofiber/fiber/v2"

type Course struct {
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	AuthorInit string  `json:"authorInit"`
	Color     string   `json:"color"`
	Students  int      `json:"students"`
	Lessons   int      `json:"lessons"`
	Duration  string   `json:"duration"`
	Progress  int      `json:"progress"`
	StatusKey string   `json:"statusKey"`
	Updated   string   `json:"updated"`
	Tag       string   `json:"tag"`
	TagC      string   `json:"tagC"`
	Rating    *float64 `json:"rating"`
}

func ptr(f float64) *float64 { 
	return &f 
}

var courses = []Course{
	{Title: "Дизайн систем интерфейсов", Author: "Мария Соколова", AuthorInit: "МС", Color: "#3DDC97", Students: 7, Lessons: 42, Duration: "14 ч", Progress: 100, StatusKey: "published", Updated: "2 дн назад", Tag: "Дизайн", TagC: "#3DDC97", Rating: ptr(4)},
	{Title: "Python для аналитиков данных", Author: "Денис Орлов", AuthorInit: "ДО", Color: "#6BB7E8", Students: 8, Lessons: 38, Duration: "22 ч", Progress: 100, StatusKey: "published", Updated: "5 ч назад", Tag: "Разработка", TagC: "#6BB7E8", Rating: ptr(5)},
	{Title: "Архитектура SwiftUI: реактивный подход", Author: "Игорь Лебедев", AuthorInit: "ИЛ", Color: "#B68DEC", Students: 6, Lessons: 28, Duration: "11 ч", Progress: 100, StatusKey: "published", Updated: "12 ч назад", Tag: "Разработка", TagC: "#6BB7E8", Rating: ptr(5)},
	{Title: "UX-исследования: основы и методы", Author: "Анна Кравцова", AuthorInit: "АК", Color: "#F5B14A", Students: 9, Lessons: 24, Duration: "9 ч", Progress: 72, StatusKey: "review", Updated: "20 мин назад", Tag: "UX", TagC: "#B68DEC", Rating: ptr(4.5)},
	{Title: "Введение в DevOps и CI/CD", Author: "Олег Беляев", AuthorInit: "ОБ", Color: "#3DDC97", Students: 10, Lessons: 31, Duration: "16 ч", Progress: 100, StatusKey: "published", Updated: "1 дн назад", Tag: "DevOps", TagC: "#F5B14A", Rating: ptr(3)},
	{Title: "Pro Figma: продвинутые техники", Author: "Мария Соколова", AuthorInit: "МС", Color: "#3DDC97", Students: 0, Lessons: 18, Duration: "7 ч", Progress: 45, StatusKey: "draft", Updated: "3 ч назад", Tag: "Дизайн", TagC: "#3DDC97", Rating: nil},
	{Title: "Machine Learning с нуля", Author: "Денис Орлов", AuthorInit: "ДО", Color: "#6BB7E8", Students: 12, Lessons: 36, Duration: "24 ч", Progress: 100, StatusKey: "published", Updated: "4 дн назад", Tag: "Разработка", TagC: "#6BB7E8", Rating: ptr(4.6)},
	{Title: "Управление продуктом: от идеи до запуска", Author: "Анна Кравцова", AuthorInit: "АК", Color: "#F5B14A", Students: 0, Lessons: 12, Duration: "5 ч", Progress: 20, StatusKey: "draft", Updated: "6 ч назад", Tag: "Продукт", TagC: "#F47174", Rating: nil},
}

func getCoursesHandler(c *fiber.Ctx) error {
	return c.JSON(courses)
}

type User struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	RoleKey    string `json:"roleKey"`
	RoleC      string `json:"roleC"`
	Registered string `json:"registered"`
	LastSeen   string `json:"lastSeen"`
	Status     string `json:"status"`
	Avatar     string `json:"avatar"`
	Color      string `json:"color"`
}

var users = []User{
	{Name: "Мария Соколова", Email: "m.sokolova@kurslab.ru", Role: "Автор курсов", RoleKey: "author", RoleC: "#3DDC97", Registered: "12.03.2024", LastSeen: "сейчас", Status: "active", Avatar: "МС", Color: "#3DDC97"},
	{Name: "Игорь Лебедев", Email: "i.lebedev@kurslab.ru", Role: "Автор курсов", RoleKey: "author", RoleC: "#3DDC97", Registered: "04.06.2024", LastSeen: "12 мин", Status: "active", Avatar: "ИЛ", Color: "#6BB7E8"},
	{Name: "Денис Орлов", Email: "d.orlov@kurslab.ru", Role: "Автор курсов", RoleKey: "author", RoleC: "#3DDC97", Registered: "21.01.2024", LastSeen: "3 ч", Status: "active", Avatar: "ДО", Color: "#B68DEC"},
	{Name: "Анна Кравцова", Email: "a.kravcova@kurslab.ru", Role: "Автор курсов", RoleKey: "author", RoleC: "#3DDC97", Registered: "08.09.2023", LastSeen: "сейчас", Status: "active", Avatar: "АК", Color: "#F5B14A"},
	{Name: "Олег Беляев", Email: "o.belyaev@kurslab.ru", Role: "Автор курсов", RoleKey: "author", RoleC: "#3DDC97", Registered: "11.05.2023", LastSeen: "20 мин", Status: "active", Avatar: "ОБ", Color: "#6BB7E8"},
	{Name: "Полина Воронина", Email: "polina.v@gmail.com", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "17.02.2025", LastSeen: "2 мин", Status: "active", Avatar: "ПВ", Color: "#3DDC97"},
	{Name: "Михаил Терентьев", Email: "m.terentev@yandex.ru", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "05.11.2024", LastSeen: "1 ч", Status: "active", Avatar: "МТ", Color: "#6BB7E8"},
	{Name: "Камилла Юсупова", Email: "k.yusupova@gmail.com", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "23.07.2024", LastSeen: "5 ч", Status: "active", Avatar: "КЮ", Color: "#B68DEC"},
	{Name: "Артём Зайцев", Email: "tema.z@gmail.com", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "02.05.2026", LastSeen: "15 мин", Status: "active", Avatar: "АЗ", Color: "#F5B14A"},
	{Name: "Виктор Кириллов", Email: "kirillov@bk.ru", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "08.05.2026", LastSeen: "40 мин", Status: "active", Avatar: "ВК", Color: "#F47174"},
	{Name: "Глеб Никитин", Email: "gleb.nik@gmail.com", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "06.08.2024", LastSeen: "3 дн", Status: "active", Avatar: "ГН", Color: "#3DDC97"},
	{Name: "Елена Громова", Email: "gromova@kurslab.ru", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "13.04.2026", LastSeen: "12 ч", Status: "active", Avatar: "ЕГ", Color: "#B68DEC"},
	{Name: "Сергей Новиков", Email: "s.novikov@list.ru", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "28.04.2026", LastSeen: "2 дн", Status: "active", Avatar: "СН", Color: "#6BB7E8"},
	{Name: "Анастасия Соловьёва", Email: "a.solov@mail.ru", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "21.10.2024", LastSeen: "5 дн", Status: "active", Avatar: "АС", Color: "#F5B14A"},
	{Name: "Юлия Маркова", Email: "yu.markova@gmail.com", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "11.05.2026", LastSeen: "вчера", Status: "active", Avatar: "ЮМ", Color: "#3DDC97"},
	{Name: "Лидия Беляева", Email: "lida.b@gmail.com", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "19.12.2024", LastSeen: "4 дн", Status: "banned", Avatar: "ЛБ", Color: "#5F7A6C"},
	{Name: "Роман Соболев", Email: "r.sobolev@mail.ru", Role: "Студент", RoleKey: "student", RoleC: "#6BB7E8", Registered: "02.10.2024", LastSeen: "12 дн", Status: "banned", Avatar: "РС", Color: "#5F7A6C"},
}

func getUsersHandler(c *fiber.Ctx) error {
	return c.JSON(users)
}

type Message struct {
	From string `json:"from"`
	You  bool   `json:"you"`
	Text string `json:"text"`
}

type Ticket struct {
	ID       string    `json:"id"`
	Subj     string    `json:"subj"`
	User     string    `json:"user"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Course   string    `json:"course"`
	Status   string    `json:"status"`
	Prio     string    `json:"prio"`
	Messages []Message `json:"messages"`
}

var tickets = []Ticket{
	{ID: "T-4838", Subj: "Не воспроизводится видео в Safari", User: "Полина Воронина", Email: "polina.v@gmail.com", Avatar: "ПВ", Course: "Архитектура SwiftUI", Status: "open", Prio: "high", Messages: []Message{
		{From: "Полина Воронина", You: false, Text: "Здравствуйте! Видео в уроке 3 не воспроизводится в Safari на Mac. Пробовала обновить страницу — не помогает."},
		{From: "Поддержка", You: true, Text: "Полина, добрый день! Уточните версию Safari и macOS, пожалуйста."},
		{From: "Полина Воронина", You: false, Text: "Safari 17.4, macOS Sonoma 14.4."},
	}},
	{ID: "T-4837", Subj: "Сертификат не пришёл на почту", User: "Михаил Терентьев", Email: "m.terentev@yandex.ru", Avatar: "МТ", Course: "Python для аналитиков", Status: "open", Prio: "med", Messages: []Message{
		{From: "Михаил Терентьев", You: false, Text: "Здравствуйте! Завершил курс по Python неделю назад, прошёл финальный тест на 92%. Сертификат на почту так и не пришёл, в личном кабинете тоже пусто. Подскажите, что не так?"},
		{From: "Поддержка", You: true, Text: "Михаил, добрый день! Проверила вашу запись — сертификат был сгенерирован, но застрял в очереди отправки. Сейчас перезапущу и пришлю на m.terentev@yandex.ru в течение 5 минут."},
		{From: "Михаил Терентьев", You: false, Text: "Спасибо! Жду."},
	}},
	{ID: "T-4836", Subj: "Ошибка при оплате курса картой", User: "Анастасия Соловьёва", Email: "a.solov@mail.ru", Avatar: "АС", Course: "UX-исследования", Status: "open", Prio: "urgent", Messages: []Message{
		{From: "Анастасия Соловьёва", You: false, Text: "При попытке оплатить курс картой Visa получаю ошибку «Платёж отклонён». Карта рабочая, деньги есть."},
	}},
	{ID: "T-4835", Subj: "Можно ли продлить доступ к материалам?", User: "Сергей Новиков", Email: "s.novikov@list.ru", Avatar: "СН", Course: "Введение в DevOps", Status: "open", Prio: "low", Messages: []Message{
		{From: "Сергей Новиков", You: false, Text: "Доступ к курсу закончился, но я ещё не прошёл все уроки. Можно продлить на месяц?"},
	}},
	{ID: "T-4834", Subj: "Прогресс сбросился после обновления", User: "Елена Громова", Email: "gromova@kurslab.ru", Avatar: "ЕГ", Course: "Дизайн систем", Status: "open", Prio: "high", Messages: []Message{
		{From: "Елена Громова", You: false, Text: "После вчерашнего обновления платформы пропал весь прогресс по курсу — 60% уроков отмечены как непройденные."},
	}},
	{ID: "T-4833", Subj: "Где скачать конспект к лекции 5?", User: "Артём Зайцев", Email: "tema.z@gmail.com", Avatar: "АЗ", Course: "Python для аналитиков", Status: "open", Prio: "low", Messages: []Message{
		{From: "Артём Зайцев", You: false, Text: "В описании к уроку упоминается конспект, но ссылка для скачивания нигде не видна."},
	}},
	{ID: "T-4832", Subj: "Не отображаются комментарии к заданию", User: "Камилла Юсупова", Email: "k.yusupova@gmail.com", Avatar: "КЮ", Course: "UX-исследования", Status: "open", Prio: "med", Messages: []Message{
		{From: "Камилла Юсупова", You: false, Text: "На странице домашнего задания раздел комментариев пустой, хотя я вижу счётчик «3 комментария»."},
	}},
	{ID: "T-4831", Subj: "Ошибка 500 при отправке домашки", User: "Виктор Кириллов", Email: "kirillov@bk.ru", Avatar: "ВК", Course: "Введение в DevOps", Status: "open", Prio: "high", Messages: []Message{
		{From: "Виктор Кириллов", You: false, Text: "При нажатии «Отправить задание» получаю белый экран с надписью «500 Internal Server Error»."},
	}},
	{ID: "T-4830", Subj: "Можно изменить ник в профиле?", User: "Лидия Беляева", Email: "lida.b@gmail.com", Avatar: "ЛБ", Course: "—", Status: "closed", Prio: "low", Messages: []Message{
		{From: "Лидия Беляева", You: false, Text: "Хочу изменить никнейм в профиле, но не нахожу такой возможности в настройках."},
		{From: "Поддержка", You: true, Text: "Лидия, смена никнейма доступна в разделе «Профиль → Редактировать». Если возникнут трудности — напишите!"},
	}},
	{ID: "T-4829", Subj: "Микрофон ментора пропадает на стримах", User: "Глеб Никитин", Email: "gleb.nik@gmail.com", Avatar: "ГН", Course: "Архитектура SwiftUI", Status: "closed", Prio: "med", Messages: []Message{
		{From: "Глеб Никитин", You: false, Text: "На живых трансляциях примерно раз в 10 минут звук ментора пропадает на 20–30 секунд."},
		{From: "Поддержка", You: true, Text: "Глеб, проблему передали команде стримингового сервиса. Уже устранена в следующем обновлении."},
	}},
}

func getTicketsHandler(c *fiber.Ctx) error {
	return c.JSON(tickets)
}

type KPI struct {
	Value    string `json:"value"`
	Delta    string `json:"delta"`
	Growing  bool   `json:"growing"`
	Subtitle string `json:"subtitle"`
}

type UserStat struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Color string `json:"color"`
}

type ChartData struct {
	Visits    []float64 `json:"visits"`
	Completed []float64 `json:"completed"`
}

var kpis = []KPI{
	{Value: "10", Delta: "+2", Growing: true, Subtitle: "против пред. периода"},
	{Value: "38м 12с", Delta: "+2м", Growing: true, Subtitle: "медиана"},
	{Value: "70%", Delta: "+2.6%", Growing: true, Subtitle: "все курсы"},
	{Value: "5%", Delta: "+1.9%", Growing: false, Subtitle: "тревога: норма ≤ 2.5%"},
}

var userStats = []UserStat{
	{Value: "17", Label: "Всего", Color: "#3DDC97"},
	{Value: "10", Label: "Активных за 7д", Color: "#6BB7E8"},
	{Value: "5", Label: "Новых за месяц", Color: "#B68DEC"},
	{Value: "2", Label: "Заблокированных", Color: "#F47174"},
}

var chartData = ChartData{
	Visits: []float64{
		6.2, 6.8, 7.1, 6.5, 7.8, 8.2, 7.6, 8.0, 8.8, 9.2,
		8.6, 9.1, 9.4, 8.8, 9.6, 10.2, 9.8, 10.4, 10.9, 10.0,
		11.2, 11.8, 11.5, 12.0, 12.4, 11.9, 13.0, 13.5, 12.8, 13.8,
	},
	Completed: []float64{
		3.8, 4.2, 4.4, 4.0, 4.8, 5.2, 4.8, 5.1, 5.6, 5.8,
		5.5, 6.0, 6.2, 5.8, 6.4, 6.8, 6.6, 7.0, 7.3, 6.8,
		7.6, 8.0, 7.8, 8.2, 8.4, 8.0, 8.8, 9.2, 8.8, 9.5,
	},
}

func getKPIHandler(c *fiber.Ctx) error {
	return c.JSON(kpis)
}

func getUserStatsHandler(c *fiber.Ctx) error {
	return c.JSON(userStats)
}

func getChartHandler(c *fiber.Ctx) error {
	return c.JSON(chartData)
}
