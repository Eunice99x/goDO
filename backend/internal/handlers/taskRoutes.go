package taskRoutes

import (
	"github.com/YounesOuterbah/goDO/internal/db"
	"github.com/YounesOuterbah/goDO/internal/models"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	tasks := []models.Todo{}

	app.Post("/api/add", func(c *fiber.Ctx) error {
		// task := &Todo{}
		task := new(models.Todo)

		if err := c.BodyParser(task); err != nil {
			return err
		}

		result, err := db.DB.Exec("INSERT INTO task (title, done) VALUES (?, ?)", task.Title, task.Done)
		if err != nil {
			return err
		}

		// Get the ID of the newly inserted task
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		task.ID = int(id)

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

	app.Delete("/api/delete/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid ID")
		}

		_, err = db.DB.Exec("DELETE FROM task WHERE id = ?", id)
		if err != nil {
			return err
		}

		return c.SendString("Task deleted successfully")

	})

	app.Get("/api/all", func(c *fiber.Ctx) error {
		rows, err := db.DB.Query("SELECT id, title, done FROM task")
		if err != nil {
			return err
		}
		defer rows.Close()

		var tasks []models.Todo
		for rows.Next() {
			var task models.Todo
			if err := rows.Scan(&task.ID, &task.Title, &task.Done); err != nil {
				return err
			}
			tasks = append(tasks, task)
		}
		if err := rows.Err(); err != nil {
			return err
		}

		return c.JSON(tasks)
	})
}