package models

import (
	"time"

	"gorm.io/gorm"
)

type EnrollmentStatus string

const (
	EnrollmentStatusActive    EnrollmentStatus = "active"
	EnrollmentStatusCompleted EnrollmentStatus = "completed"
	EnrollmentStatusDropped   EnrollmentStatus = "dropped"
)

type Enrollment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	UserID   uint   `json:"user_id" gorm:"not null"`
	User     User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CourseID uint   `json:"course_id" gorm:"not null"`
	Course   Course `json:"course,omitempty" gorm:"foreignKey:CourseID"`

	Status      EnrollmentStatus `json:"status" gorm:"type:varchar(20);default:'active'"`
	EnrolledAt  time.Time        `json:"enrolled_at"`
	CompletedAt *time.Time       `json:"completed_at,omitempty"`

	Progress       float64    `json:"progress" gorm:"default:0"` // percentage 0-100
	LastAccessedAt *time.Time `json:"last_accessed_at,omitempty"`

	// Payment reference
	PaymentID *uint    `json:"payment_id,omitempty"`
	Payment   *Payment `json:"payment,omitempty" gorm:"foreignKey:PaymentID"`
}
