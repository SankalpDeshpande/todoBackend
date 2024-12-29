package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title" validate:"required,min=3,max=100"`
	Status string `json:"status" validate:"required,oneof=pending in-progress completed"`
}

func main() {
	// Load environment variables from .env file (only for local development)
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file", err)
	// }

	// connStr := "user=postgres password=postgres dbname=todos sslmode=disable"

	// Get environment variables
	connStr, errBool := os.LookupEnv("DATABASE_URL")
	if !errBool {
		log.Fatal("DATABASE_URL environment variable is required")
	}
	port, errBool := os.LookupEnv("PORT")
	if !errBool {
		port = "8080"
	}

	log.Println("Connecting to database...")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	r := gin.Default()

	validate := validator.New()

	// In-memory data storage
	// var todos = []Task{}

	// Routes
	// Get all todos
	r.GET("/todos", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, status FROM todos")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var tasks []Task
		for rows.Next() {
			var task Task
			if err := rows.Scan(&task.ID, &task.Title, &task.Status); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			tasks = append(tasks, task)
		}
		c.JSON(http.StatusOK, tasks)
	})

	r.POST("/todos", func(c *gin.Context) {
		var newTask Task
		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Generate a unique ID
		newTask.ID = uuid.New().String()
		if err := validate.Struct(newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
			return
		}

		query := "Insert into todos (id, title, status) values ($1,$2,$3)"
		_, err := db.Exec(query, newTask.ID, newTask.Title, newTask.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, newTask)
	})

	// Get a task by ID
	r.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var task Task

		query := "SELECT id, title, status FROM todos WHERE id = $1"
		err := db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Status)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, task)
	})

	// Update a task
	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedTask Task
		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := "UPDATE todos SET title = $1, status = $2 WHERE id = $3"
		result, err := db.Exec(query, updatedTask.Title, updatedTask.Status, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
			return
		}

		updatedTask.ID = id
		c.JSON(http.StatusOK, updatedTask)
	})

	// Delete a task
	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		query := "DELETE FROM todos WHERE id = $1"
		result, err := db.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
	})

	// Update task status
	r.PATCH("/todos/:id/status", func(c *gin.Context) {
		id := c.Param("id")
		var statusUpdate struct {
			Status string `json:"status"`
		}
		if err := c.ShouldBindJSON(&statusUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := "UPDATE todos SET status = $1 WHERE id = $2"
		result, err := db.Exec(query, statusUpdate.Status, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "task updated"})
	})
	err = r.Run("0.0.0.0:" + port)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error starting server: %s", err))
	}
}
