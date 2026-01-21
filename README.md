#  Task Management Service (Golang)

A scalable RESTful Task Management API built in Go demonstrating:

- Clean Architecture  
- JWT Authentication & Authorization  
- PostgreSQL persistence  
- Background workers with concurrency  
- Structured logging  
- Environment-based configuration  

This project is designed as a **production-grade backend service** and suitable for system design & backend interviews.

---

##  Features

- User Signup & Login with JWT  
- Role-based Authorization (user / admin)  
- CRUD APIs for Tasks  
- Users can access only their own tasks  
- Admin can access all tasks  
- Background worker auto-completes tasks after configurable minutes  
- Structured logging with timestamp & levels  
- Clean layered architecture  

---

## Project Structure

task_management/
│
├── main.go # Application bootstrap
├── config/ # Environment & config loader
├── db/ # Database connection
├── routers/ # Route wiring
├── handler/ # HTTP handlers
├── service/ # Business logic
├── repository/ # Data access layer
├── middleware/ # Auth & security middleware
├── worker/ # Background worker
├── model/ # Domain models
├── pkg/logger/ # Custom structured logger
└── README.md 

## Create ENV 

---

##  Environment Variables

Create a `.env` file in root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=####
DB_NAME=taskdb
DB_SSLMODE=disable

JWT_SECRET=supersecret
AUTO_COMPLETE_MINUTES=2
PORT=8080 
---

## Run Application
go mod tidy
go run main.go


## Authentication Flow

Signup user

Login user → receive JWT token

Use token in Authorization header
Create Task (Protected) 
POST /tasks 
{
  "title": "Learn Golang",
  "description": "Study goroutines and channels"
} 

List Tasks

GET /tasks

🔹 Get Task by ID

GET /tasks/{id}

🔹 Delete Task

DELETE /tasks/{id}



## Background Worker (Auto Complete)

Configurable via: AUTO_COMPLETE_MINUTES

Automatically marks tasks as completed

Skips tasks already completed or deleted

Runs concurrently using goroutines & channels


## Logger also added to checking the status on prometheus/Argo CI/CD
