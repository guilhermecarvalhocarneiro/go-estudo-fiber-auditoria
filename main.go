package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	fiber_log "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sixfwa/fiber-gorm/database"
	"github.com/sixfwa/fiber-gorm/routes"
)

func healtcheck(c *fiber.Ctx) error {
	return c.Status(200).JSON("OK")
}

func setupRoutes(app *fiber.App) {
	// Agrupando pela versão da API
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/healtcheck", healtcheck)

	// Auditoria endpoints
	v1.Post("/auditoria", routes.CreateAuditoria)
	v1.Get("/auditoria", routes.GetAuditorias)
	v1.Get("/auditoria/get/:id", routes.GetAuditoria)
	v1.Get("/auditoria/count", routes.GetCountRegistersAuditoria)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Use(fiber_log.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
