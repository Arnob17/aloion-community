package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleStudent UserRole = "student"
	RoleTeacher UserRole = "teacher"
	RoleAdmin   UserRole = "admin"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Basic Information
	FirstName   string     `json:"first_name" gorm:"not null"`
	LastName    string     `json:"last_name" gorm:"not null"`
	Email       string     `json:"email" gorm:"uniqueIndex;not null"`
	Phone       string     `json:"phone" gorm:"index"`
	Password    string     `json:"-" gorm:"not null"` // Hashed password
	DateOfBirth *time.Time `json:"date_of_birth"`

	// Role and Status
	Role       UserRole `json:"role" gorm:"type:varchar(20);default:'student'"`
	IsActive   bool     `json:"is_active" gorm:"default:true"`
	IsVerified bool     `json:"is_verified" gorm:"default:false"`

	// Student-specific fields
	Class         *int   `json:"class"` // Classes 8-10 initially
	SchoolName    string `json:"school_name"`
	Address       string `json:"address"`
	GuardianName  string `json:"guardian_name"`
	GuardianPhone string `json:"guardian_phone"`

	// Teacher-specific fields
	Qualification string `json:"qualification"`
	Bio           string `json:"bio" gorm:"type:text"`
	Experience    int    `json:"experience"` // years

	// Financial
	TotalSpent  float64 `json:"total_spent" gorm:"default:0"`
	TotalEarned float64 `json:"total_earned" gorm:"default:0"` // For teachers

	// Relationships
	Enrollments    []Enrollment   `json:"enrollments,omitempty" gorm:"foreignKey:UserID"`
	Subscriptions  []Subscription `json:"subscriptions,omitempty" gorm:"foreignKey:UserID"`
	Payments       []Payment      `json:"payments,omitempty" gorm:"foreignKey:UserID"`
	CreatedCourses []Course       `json:"created_courses,omitempty" gorm:"foreignKey:TeacherID"`
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
