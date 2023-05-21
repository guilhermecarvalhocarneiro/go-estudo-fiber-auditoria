package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-gorm/database"
	"github.com/sixfwa/fiber-gorm/models"
	_ "github.com/sixfwa/fiber-gorm/pkg"
)

func CreateAuditoria(c *fiber.Ctx) error {
	var auditoria models.Auditoria
	if err := c.BodyParser(&auditoria); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.Database.Db.Create(&auditoria)
	return c.Status(200).JSON(auditoria)
}

func GetAuditorias(c *fiber.Ctx) error {
	auditorias := []models.Auditoria{}
	// database.Database.Db.Scopes(pkg.NewPaginate(100, 2).PaginatedResult).Find(&auditorias)
	database.Database.Db.Find(&auditorias)

	

	// Criando o Map de retorno com a chave results e o total
	result := fiber.Map{
		"results": auditorias[:3],
		"total":   len(auditorias),
	}

	return c.Status(200).JSON(result)
}

func GetAuditoria(c *fiber.Ctx) error {
	id := c.Params("id")
	// Verificando se o id é do tipo inteiro
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	auditoria := models.Auditoria{}
	database.Database.Db.Find(&auditoria, id)
	// Verificando se retornou algum registro
	if auditoria.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	return c.Status(200).JSON(auditoria)
}

func GetCountRegistersAuditoria(c *fiber.Ctx) error {
	auditoria := []models.Auditoria{}
	result := database.Database.Db.Find(&auditoria)
	return c.Status(200).JSON(fiber.Map{"total": result.RowsAffected})
}
