package models

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	CourseID uint   `json:"course_id" gorm:"not null"`
	Course   Course `json:"course,omitempty" gorm:"foreignKey:CourseID"`

	Title        string  `json:"title" gorm:"not null"`
	Description  string  `json:"description" gorm:"type:text"`
	Instructions string  `json:"instructions" gorm:"type:text"`
	MaxScore     float64 `json:"max_score" gorm:"default:100"`

	DueDate     *time.Time `json:"due_date,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`

	// Relationships
	Submissions []Submission `json:"submissions,omitempty" gorm:"foreignKey:AssignmentID"`
}

type Submission struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	AssignmentID uint       `json:"assignment_id" gorm:"not null"`
	Assignment   Assignment `json:"assignment,omitempty" gorm:"foreignKey:AssignmentID"`
	UserID       uint       `json:"user_id" gorm:"not null"`
	User         User       `json:"user,omitempty" gorm:"foreignKey:UserID"`

	Content     string     `json:"content" gorm:"type:text"`
	FileURL     string     `json:"file_url,omitempty"`
	Score       *float64   `json:"score,omitempty"`
	Feedback    string     `json:"feedback" gorm:"type:text"`
	SubmittedAt time.Time  `json:"submitted_at"`
	GradedAt    *time.Time `json:"graded_at,omitempty"`
}
