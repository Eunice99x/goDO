package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID 		int 	`json:"id"`
	Title 	string 	`json:"title"`
	Done 	bool 	`json:"done"`
	Body 	string 	`json:"body"`
}

func main() {
	app := fiber.New()

	tasks := []Todo{}

	app.Use(cors.New())

	app.Post("/api/add", func(c *fiber.Ctx) error {
		task := &Todo{}

		if err := c.BodyParser(task); err != nil {
			return err
		}

		task.ID = len(tasks) + 1

		tasks = append(tasks, *task)

		return c.JSON(task)
	})

	app.Patch("/api/add/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid ID")
		}

		for i, t := range tasks {
			if t.ID == id {
				tasks[i].Done = true
				break
			}
		}

		return c.JSON(tasks)
	})

	app.Get("/api/all", func(c *fiber.Ctx) error {
		return c.JSON(tasks)
	})

	log.Fatal(app.Listen(":8080"))
}