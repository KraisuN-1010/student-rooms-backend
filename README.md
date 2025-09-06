# Student Rooms Backend

A Go backend API for a student room management system built with GoFr framework and PostgreSQL.

## Features

- **User Authentication**: Sign up and login with JWT tokens
- **Room Management**: Create and manage study rooms
- **Notes System**: Share notes, topics, and announcements
- **Doubt Resolution**: Ask and answer questions
- **Comments System**: Comment on notes and doubts
- **File Support**: Upload and share files

## Tech Stack

- **Language**: Go 1.25+
- **Framework**: GoFr
- **Database**: PostgreSQL (via Supabase)
- **Authentication**: JWT tokens
- **Password Hashing**: bcrypt

## Prerequisites

- Go 1.25 or higher
- PostgreSQL database (or Supabase account)
- Environment variables configured

## Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd student-rooms-backend
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   Create a `.env` file or set the following environment variables:
   ```bash
   SUPABASE_DB_URL=postgresql://username:password@host:port/database
   JWT_SECRET=your-secret-key
   PORT=8000
   ```

4. **Set up the database**
   Run the SQL schema from `db/schema.sql` in your PostgreSQL database:
   ```bash
   psql -d your_database -f db/schema.sql
   ```

5. **Run the application**
   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

### Authentication
- `POST /auth/signup` - User registration
- `POST /auth/login` - User login

### Rooms
- `GET /rooms` - Get all rooms
- `POST /rooms` - Create a new room

### Notes
- `GET /rooms/:roomId/notes` - Get notes for a room
- `POST /rooms/:roomId/notes` - Create a new note

### Comments
- `GET /comments/:parentId` - Get comments for a parent (note/doubt)
- `POST /comments` - Create a new comment

### Doubts
- `GET /rooms/:roomId/doubts` - Get doubts for a room
- `POST /rooms/:roomId/doubts` - Create a new doubt

## Project Structure

```
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── db/
│   ├── supabase.go          # Database connection
│   └── schema.sql           # Database schema
├── internal/
│   └── handlers/            # HTTP handlers
│       ├── auth.go
│       ├── rooms.go
│       ├── notes.go
│       ├── comments.go
│       └── doubts.go
├── models/                  # Data models
│   ├── user.go
│   ├── room.go
│   ├── note.go
│   ├── comment.go
│   └── doubt.go
├── services/                # Business logic
│   ├── types.go            # Request/Response types
│   ├── auth_service.go
│   ├── room_service.go
│   ├── note_service.go
│   ├── comment_service.go
│   └── doubt_service.go
├── go.mod
├── go.sum
└── README.md
```

## Development

### Running Tests
```bash
go test ./...
```

### Building
```bash
go build -o bin/server cmd/server/main.go
```

### Database Migrations
The database schema is provided in `db/schema.sql`. Run this script to set up your database tables.

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `SUPABASE_DB_URL` | PostgreSQL connection string | Yes |
| `JWT_SECRET` | Secret key for JWT signing | Yes |
| `PORT` | Server port (default: 8000) | No |

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.
