# Event Finder

A full-stack web application for discovering and booking events. Built with Go backend and Next.js frontend.

## Features

- **User Authentication**: Register and login with JWT-based authentication
- **Event Management**: Browse, create, update, and delete events (admin-only for CRUD operations)
- **Categories**: Organize events by categories
- **Bookings**: Book event tickets with booking items
- **Admin Panel**: Manage events, categories, bookings, and users
- **Responsive UI**: Modern, accessible interface using shadcn/ui components

## Tech Stack

### Backend
- **Language**: Go
- **Framework**: Chi router for HTTP routing
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT tokens
- **Middleware**: Custom auth middleware

### Frontend
- **Framework**: Next.js 14+ with App Router
- **Styling**: Tailwind CSS
- **UI Components**: shadcn/ui
- **State Management**: React hooks (with optional Zustand for global state)
- **API Integration**: Fetch API with potential SWR/TanStack Query

## Project Structure

```
event-finder/
├── backend/
│   ├── cmd/                # CLI commands (e.g., admin creation)
│   ├── internal/
│   │   ├── api/           # HTTP handlers
│   │   ├── app/           # Application initialization
│   │   ├── config/        # Database configuration
│   │   ├── middleware/    # Authentication middleware
│   │   ├── models/        # GORM models
│   │   ├── routes/        # Route definitions
│   │   ├── store/         # Data access layer (repositories)
│   │   ├── tokens/        # Token utilities
│   │   └── utils/         # Helper functions
│   ├── main.go            # Application entry point
│   ├── go.mod             # Go modules
│   └── go.sum
├── frontend/
│   ├── event-finder-frontend/
│   │   ├── app/           # Next.js App Router pages
│   │   ├── components/    # React components (ui, shared, features)
│   │   ├── lib/           # Utilities, API clients, hooks, types
│   │   ├── public/        # Static assets
│   │   ├── next.config.ts
│   │   ├── package.json
│   │   └── tsconfig.json
└── README.md
```

## Getting Started

### Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL
- Git

### Backend Setup
1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables. Create a `.env` file in `backend/`:
   ```
   DATABASE_URL=postgres://user:password@localhost:5432/event_finder?sslmode=disable
   JWT_SECRET=your-secret-key
   ```

4. Run database migrations:
   ```bash
   go run main.go
   # First run will initialize and migrate the database
   ```

5. Create an admin user (optional):
   ```bash
   go run main.go create-admin
   ```

6. Start the backend server:
   ```bash
   go run main.go --port 8080
   ```

The backend will run on `http://localhost:8080`.

### Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd frontend/event-finder-frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   # or
   pnpm install
   ```

3. Set up environment variables. Create a `.env.local` file:
   ```
   NEXT_PUBLIC_API_URL=http://localhost:8080
   ```

4. Start the development server:
   ```bash
   npm run dev
   # or
   pnpm dev
   ```

The frontend will run on `http://localhost:3000`.

### API Endpoints

#### Public Endpoints
- `GET /` - Welcome message
- `GET /health` - Health check
- `POST /register` - User registration
- `POST /login` - User login

#### Protected Endpoints (require authentication)
- `GET /categories` - List categories
- `GET /events` - List events
- `GET /bookings` - List user bookings
- `GET /booking_items` - List booking items

#### Admin Endpoints (require admin role)
- `POST /categories` - Create category
- `PUT /categories/{id}` - Update category
- `DELETE /categories/{id}` - Delete category
- `POST /events` - Create event
- `PUT /events/{id}` - Update event
- `DELETE /events/{id}` - Delete event
- And similar for bookings, tickets, etc.

All endpoints return JSON responses with consistent error handling.

## Development

### Running Tests
- Backend: `cd backend && go test ./...`
- Frontend: `cd frontend/event-finder-frontend && npm test`

### Linting
- Backend: Use `gofmt` or `goimports`
- Frontend: `npm run lint`

### Building for Production
- Backend: `cd backend && go build -o event-finder main.go`
- Frontend: `cd frontend/event-finder-frontend && npm run build`


