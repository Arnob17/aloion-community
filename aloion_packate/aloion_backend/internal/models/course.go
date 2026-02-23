package models

import (
	"time"

	"gorm.io/gorm"
)

type CourseLevel string

const (
	LevelBeginner     CourseLevel = "beginner"
	LevelIntermediate CourseLevel = "intermediate"
	LevelAdvanced     CourseLevel = "advanced"
	LevelOlympiad     CourseLevel = "olympiad"
)

type CourseStatus string

const (
	CourseStatusDraft     CourseStatus = "draft"
	CourseStatusPublished CourseStatus = "published"
	CourseStatusArchived  CourseStatus = "archived"
)

type Course struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Basic Information
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"type:text"`
	Subject     string `json:"subject" gorm:"not null"` // Mathematics, Physics, etc.
	Class       int    `json:"class"`                   // 8, 9, 10, etc.

	// Course Details
	Level        CourseLevel  `json:"level" gorm:"type:varchar(20);default:'beginner'"`
	Status       CourseStatus `json:"status" gorm:"type:varchar(20);default:'draft'"`
	Duration     int          `json:"duration"` // in weeks
	TotalLessons int          `json:"total_lessons" gorm:"default:0"`

	// Pricing
	Price    float64 `json:"price" gorm:"default:0"` // 0 for free courses
	IsFree   bool    `json:"is_free" gorm:"default:false"`
	Discount float64 `json:"discount" gorm:"default:0"` // percentage

	// Teacher
	TeacherID uint `json:"teacher_id" gorm:"not null"`
	Teacher   User `json:"teacher,omitempty" gorm:"foreignKey:TeacherID"`

	// Metadata
	ThumbnailURL string `json:"thumbnail_url"`
	Language     string `json:"language" gorm:"default:'bn'"` // bn for Bengali, en for English

	// Relationships
	Enrollments []Enrollment     `json:"enrollments,omitempty" gorm:"foreignKey:CourseID"`
	Materials   []CourseMaterial `json:"materials,omitempty" gorm:"foreignKey:CourseID"`
	Assignments []Assignment     `json:"assignments,omitempty" gorm:"foreignKey:CourseID"`
}
