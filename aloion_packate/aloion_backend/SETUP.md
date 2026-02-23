# Setup Guide for Aloion Backend

## Quick Start

### 1. Install Dependencies

```bash
go mod download
go mod tidy
```

If you encounter permission issues with Go's module cache, you may need to:
```bash
sudo chown -R $USER:$USER ~/.cache/go-build
sudo chown -R $USER:$USER ~/go/pkg/mod
```

### 2. Setup PostgreSQL Database

```bash
# Login to PostgreSQL
sudo -u postgres psql

# Create database
CREATE DATABASE aloion;

# Create user (optional)
CREATE USER aloion_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE aloion TO aloion_user;
```

### 3. Configure Environment Variables

```bash
cp .env.example .env
```

Edit `.env` with your database credentials:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=aloion
```

### 4. Run the Server

**Option 1: Standard run**
```bash
# Using Make
make run

# Or directly
go run cmd/server/main.go
```

**Option 2: Development with Air (Recommended for development)**
```bash
# Install Air (one-time setup)
make install-air
# or
go install github.com/cosmtrek/air@latest

# Run with Air (auto-reload on file changes)
make air
# or
air
```

The server will start on `http://localhost:8080`

**Note:** Air watches for file changes and automatically rebuilds and restarts the server, making development much faster!

## Testing the API

### 1. Register a User

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Arnob",
    "last_name": "Rahman",
    "email": "arnob@aloion.com",
    "password": "password123",
    "role": "admin",
    "phone": "+8801234567890"
  }'
```

### 2. Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "arnob@aloion.com",
    "password": "password123"
  }'
```

Save the token from the response.

### 3. Create a Course (as Teacher/Admin)

```bash
curl -X POST http://localhost:8080/api/v1/courses \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "title": "Mathematics for Class 9",
    "description": "Comprehensive mathematics course aligned with local syllabus",
    "subject": "Mathematics",
    "class": 9,
    "level": "beginner",
    "duration": 12,
    "price": 500,
    "is_free": false,
    "language": "bn"
  }'
```

### 4. Get All Courses

```bash
curl http://localhost:8080/api/v1/courses
```

## Project Structure

```
aloion_backend/
├── cmd/server/main.go          # Entry point
├── internal/
│   ├── config/                 # Configuration
│   ├── database/               # DB connection & migrations
│   ├── models/                 # Database models
│   │   ├── user.go
│   │   ├── course.go
│   │   ├── enrollment.go
│   │   ├── subscription.go
│   │   ├── payment.go
│   │   ├── material.go
│   │   └── assignment.go
│   ├── dto/                    # Data Transfer Objects
│   ├── handlers/               # HTTP handlers
│   ├── services/               # Business logic
│   ├── middleware/             # Auth & other middleware
│   ├── routes/                 # Route definitions
│   └── utils/                  # Utilities (JWT, password hashing)
├── go.mod
├── Makefile
└── README.md
```

## Database Models

### User
- Supports Student, Teacher, and Admin roles
- Student fields: class, school, guardian info
- Teacher fields: qualification, experience, bio

### Course
- Subject-based (Mathematics, Physics, etc.)
- Class levels (8, 9, 10)
- Difficulty levels (beginner, intermediate, advanced, olympiad)
- Pricing and status management

### Enrollment
- Tracks student course enrollments
- Progress tracking
- Status: active, completed, dropped

### Subscription
- Monthly/Yearly plans
- Subscription lifecycle management

### Payment
- Multiple payment methods (bKash, Nagad, Rocket, Bank, Cash)
- Payment status tracking

## Next Steps

1. **Add Email Verification**: Implement email verification for user registration
2. **Payment Gateway Integration**: Integrate with bKash/Nagad APIs
3. **File Upload**: Add support for course materials and assignments
4. **Notifications**: Add notification system for course updates
5. **Analytics**: Add analytics for course performance and student progress

## Troubleshooting

### Database Connection Issues
- Ensure PostgreSQL is running: `sudo systemctl status postgresql`
- Check database credentials in `.env`
- Verify database exists: `psql -U postgres -l`

### Permission Issues
- Fix Go module cache permissions (see step 1)
- Ensure you have write permissions in the project directory

### Port Already in Use
- Change `SERVER_PORT` in `.env`
- Or kill the process using port 8080: `lsof -ti:8080 | xargs kill`
