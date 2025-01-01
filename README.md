# todoBackend
Backend implementation of Todo app using Golang, Gin framework and Postgresql


Deployed at https://todobackend-fegt.onrender.com


Example-
GET - https://todobackend-fegt.onrender.com/todos/

POST - https://todobackend-fegt.onrender.com/todos/
Body - {
        "title": "Learn Gin",
        "status": "pending"
      }


PATCH - https://todobackend-fegt.onrender.com/todos/:id/status

PUT - https://todobackend-fegt.onrender.com/todos/:id

DELETE - https://todobackend-fegt.onrender.com/todos


Packages used - 
1. Gin - github.com/gin-gonic/gin - web framework
2. Validator - github.com/go-playground/validator/v10 - implements value validations for structs and individual fields based on tags.
3. UUID - github.com/google/uuid - package generates and inspects UUIDs
4. godotenv - github.com/joho/godotenv - loads env vars from a .env file
5. pq - github.com/lib/pq - A pure Go postgres driver for Go's database/sql package
