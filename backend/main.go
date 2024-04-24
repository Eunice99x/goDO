package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID 		int 	`json:"id"`
	Title 	string 	`json:"title"`
	Done 	bool 	`json:"done"`
}

var db *sql.DB

func databaseConnection(){
	dsn := "root:eunice99x@tcp(localhost:3306)/todos?parseTime=true"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Connected to MySQL database successfully")
}

func setupRoutes(app *fiber.App){
	tasks := []Todo{}

	app.Post("/api/add", func(c *fiber.Ctx) error {
		// task := &Todo{}
		task := new(Todo)

		if err := c.BodyParser(task); err != nil {
			return err
		}

		result, err := db.Exec("INSERT INTO task (title, done) VALUES (?, ?)", task.Title, task.Done)
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

		_, err = db.Exec("DELETE FROM task WHERE id = ?", id)
		if err != nil {
			return err
		}

		return c.SendString("Task deleted successfully")

	})

	app.Get("/api/all", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id, title, done FROM task")
	if err != nil {
		return err
	}
	defer rows.Close()

	var tasks []Todo
	for rows.Next() {
		var task Todo
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

func main() {
	app := fiber.New()

	app.Use(cors.New())

	databaseConnection()
	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
