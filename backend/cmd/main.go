package main

import (
	"log"

	"github.com/YounesOuterbah/goDO/internal/db"
	taskRoutes "github.com/YounesOuterbah/goDO/internal/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)



func main() {
	app := fiber.New()

	app.Use(cors.New())

	db.DatabaseConnection()
	taskRoutes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
