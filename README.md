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
