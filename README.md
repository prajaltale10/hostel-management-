# Hostel Management SaaS

A complete production-grade multi-tenant Hostel Management system built with a Golang/Gin monolith backend and a Next.js App Router frontend.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:
- **Docker & Docker Compose:** For running Postgres and Redis easily.
- **Go (1.23+):** For running the backend natively.
- **Node.js (20+) & npm:** For running the frontend natively.

## Environment Variables

A `.env` file is required in the root directory for docker-compose and in the `backend/` directory if running natively.

Create a `.env` file in the root directory with the following variables:
```env
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hostel_saas
JWT_SECRET=super_secret_key_for_jwt_auth_12345
PORT=8080
ENV=development
```

## Running the Application

### Option 1: Full Docker Setup (Easiest)

You can spin up the entire application (Postgres, Redis, Backend, Frontend) using Docker Compose.

Wait a few seconds for the database to initialize.
- Frontend will be available at: `http://localhost:3000`
- Backend API will be available at: `http://localhost:8080/api/v1`

### Option 2: Native Development (VS Code)

If you prefer to run the backend and frontend locally (useful for debugging in VS Code), you still need Postgres and Redis running.

**1. Start Infrastructure:**
Start your docker compose db and redis containers via:
`docker-compose up -d db redis`

**2. Start the Backend:**
Open a terminal in the `backend/` folder and run the go application via:
`go run cmd/api/main.go`

**3. Start the Frontend:**
Open a terminal in the `frontend/` folder, install dependencies, and start the development server via:
`npm install` and then start the frontend in the background with `npm run start &` (or `npm run dev &` for dev server).

## VS Code Integration

This project includes `.vscode` configuration files to easily launch both the frontend and backend using the VS Code debugger.

1. Go to the "Run and Debug" tab in VS Code.
2. Select "Run All (Backend + Frontend)" from the dropdown.
3. Click the green play button.
*(Note: Ensure Docker containers for DB and Redis are running first).*
