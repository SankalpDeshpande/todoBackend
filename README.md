todoBackend

Backend implementation of Todo app using Golang, Gin framework and Postgresql

Deployed at https://todobackend-fegt.onrender.com

Example- GET - https://todobackend-fegt.onrender.com/todos/

POST - https://todobackend-fegt.onrender.com/todos/ Body - { "title": "Learn Gin", "status": "pending" }

PATCH - https://todobackend-fegt.onrender.com/todos/:id/status

PUT - https://todobackend-fegt.onrender.com/todos/:id

DELETE - https://todobackend-fegt.onrender.com/todos

Packages used -

    Gin - github.com/gin-gonic/gin - web framework
    Validator - github.com/go-playground/validator/v10 - implements value validations for structs and individual fields based on tags.
    UUID - github.com/google/uuid - package generates and inspects UUIDs
    godotenv - github.com/joho/godotenv - loads env vars from a .env file
    pq - github.com/lib/pq - A pure Go postgres driver for Go's database/sql package
