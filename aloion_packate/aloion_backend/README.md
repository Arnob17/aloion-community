# Aloion Backend

A structured Gin-based backend API for Aloion - an educational NGO platform in Bangladesh.

## Features

- **User Management**: Student, Teacher, and Admin roles
- **Course Management**: Create, update, and manage courses
- **Enrollment System**: Students can enroll in courses
- **Subscription System**: Monthly/Yearly subscription plans
- **Payment Integration**: Support for bKash, Nagad, Rocket, Bank, and Cash
- **JWT Authentication**: Secure token-based authentication
- **Database ORM**: GORM with PostgreSQL

## Project Structure

```
aloion_backend/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── config/              # Configuration management
│   ├── database/            # Database connection and migrations
│   ├── models/              # Data models/entities
│   ├── dto/                 # Data Transfer Objects
│   ├── handlers/            # HTTP handlers
│   ├── services/            # Business logic
│   ├── middleware/          # HTTP middleware (auth, etc.)
│   ├── routes/              # Route definitions
│   └── utils/               # Utility functions
├── go.mod
├── go.sum
└── README.md
```

## Setup

### Prerequisites

- Go 1.25.5 or higher
- PostgreSQL 12 or higher

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd aloion_backend
```

2. Install dependencies:
```bash
go mod download
go mod tidy
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your database credentials
```

4. Create PostgreSQL database:
```sql
CREATE DATABASE aloion;
```

5. Run the application:

**Option 1: Standard run**
```bash
go run cmd/server/main.go
# or
make run
```

**Option 2: Development with Air (live reload)**
```bash
# Install Air first (one-time)
make install-air
# or
go install github.com/cosmtrek/air@latest

# Run with Air
make air
# or
air
```

The server will start on `http://localhost:8080`

Air will automatically rebuild and restart the server when you make changes to the code.

## API Endpoints

### Authentication

- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login user

### Users (Protected)

- `GET /api/v1/users/me` - Get current user profile
- `GET /api/v1/users` - Get all users (Admin only)

### Courses

- `GET /api/v1/courses` - Get all courses (Public)
- `GET /api/v1/courses/:id` - Get course by ID (Public)
- `POST /api/v1/courses` - Create course (Teacher/Admin)
- `PUT /api/v1/courses/:id` - Update course (Teacher/Admin)
- `DELETE /api/v1/courses/:id` - Delete course (Teacher/Admin)
- `POST /api/v1/courses/:id/enroll` - Enroll in course (Student)
- `GET /api/v1/courses/my-enrollments` - Get my enrollments (Student)

### Health Check

- `GET /api/v1/health` - Health check endpoint

## Example API Usage

### Register a Student

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "password": "password123",
    "role": "student",
    "class": 9,
    "school_name": "Example School"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create a Course (Teacher/Admin)

```bash
curl -X POST http://localhost:8080/api/v1/courses \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "title": "Mathematics for Class 9",
    "description": "Comprehensive mathematics course",
    "subject": "Mathematics",
    "class": 9,
    "level": "beginner",
    "duration": 12,
    "price": 500,
    "is_free": false,
    "language": "bn"
  }'
```

### Enroll in a Course

```bash
curl -X POST http://localhost:8080/api/v1/courses/1/enroll \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "course_id": 1
  }'
```

## Database Models

### User
- Students, Teachers, and Admins
- Role-based access control
- Student-specific fields (class, school, guardian info)
- Teacher-specific fields (qualification, experience)

### Course
- Subject-based courses (Mathematics, Physics, etc.)
- Class levels (8, 9, 10, etc.)
- Difficulty levels (beginner, intermediate, advanced, olympiad)
- Pricing and subscription options

### Enrollment
- Track student course enrollments
- Progress tracking
- Status management (active, completed, dropped)

### Subscription
- Monthly/Yearly plans
- Subscription status tracking
- Renewal management

### Payment
- Multiple payment methods (bKash, Nagad, Rocket, Bank, Cash)
- Payment status tracking
- Transaction management

## Environment Variables

See `.env.example` for all available environment variables.

## Development

### Running in Development Mode

**Standard mode:**
```bash
ENV=development go run cmd/server/main.go
```

**With Air (live reload - recommended):**
```bash
# Install Air (one-time)
make install-air

# Run with Air
make air
```

Air will automatically rebuild and restart the server when you make changes to any `.go` files.

### Database Migrations

Migrations run automatically on server start. GORM AutoMigrate is used to create/update database schema.

## License

This project is part of Aloion - an educational NGO initiative in Bangladesh.
