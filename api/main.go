package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	port := "8080"

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5173, http://localhost:5173, https://app.nikitos.pro",
		AllowMethods: "GET",
	}))

	app.Get("/courses", getCoursesHandler)
	app.Get("/users", getUsersHandler)
	app.Get("/tickets", getTicketsHandler)
	app.Get("/stats/kpi", getKPIHandler)
	app.Get("/stats/users", getUserStatsHandler)
	app.Get("/stats/chart", getChartHandler)

	app.Listen(":" + port)
}
