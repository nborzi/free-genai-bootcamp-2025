# Language Portal Backend

This is the backend server for the Language Portal application, built with Go, Gin, and SQLite3.

## Prerequisites

1. Install Go (1.21 or later)
   - Download from: https://golang.org/dl/
   - Add Go to your PATH environment variable

2. Install SQLite3
   - Download from: https://www.sqlite.org/download.html
   - Add SQLite3 to your PATH environment variable

## Setup

1. Clone the repository
2. Navigate to the backend_go directory
3. Initialize the Go module and download dependencies:
   ```bash
   go mod init lang-portal
   go mod tidy
   ```

4. Create the SQLite database and run migrations:
   ```bash
   sqlite3 words.db < db/migrations/001_initial_schema.sql
   ```

## Running the Server

Start the server with:
```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

- GET /api/dashboard/last_study_session
- GET /api/dashboard/study_progress
- GET /api/dashboard/quick-stats
- GET /api/study_activities/:id
- GET /api/study_activities/:id/study_sessions
- POST /api/study_activities
- GET /api/words
- GET /api/words/:id

## Project Structure

```
backend_go/
├── cmd/
│   └── server/          # Main application entry point
├── internal/
│   ├── models/          # Data structures and database operations
│   ├── handlers/        # HTTP handlers
│   └── service/         # Business logic
├── db/
│   ├── migrations/      # Database migrations
│   └── seeds/          # Seed data
└── words.db            # SQLite database
```
