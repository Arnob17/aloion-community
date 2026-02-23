package models

import (
	"time"

	"gorm.io/gorm"
)

type Exam struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Title       string  `json:"title" gorm:"not null"`
	Description string  `json:"description" gorm:"type:text"`
	CourseID    *uint   `json:"course_id"`
	Course      *Course `json:"course,omitempty" gorm:"foreignKey:CourseID"`
	BatchID     *uint   `json:"batch_id"` // For offline batch tracking

	MaxScore  float64   `json:"max_score" gorm:"default:100;not null"`
	ExamDate  time.Time `json:"exam_date" gorm:"not null"`

	// Relationships
	Results []ExamResult `json:"results,omitempty" gorm:"foreignKey:ExamID"`
}

type ExamResult struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	ExamID uint `json:"exam_id" gorm:"not null"`
	Exam   Exam `json:"exam,omitempty" gorm:"foreignKey:ExamID"`

	UserID uint `json:"user_id" gorm:"not null"`
	User   User `json:"user,omitempty" gorm:"foreignKey:UserID"`

	Score    float64 `json:"score" gorm:"not null"`
	MaxScore float64 `json:"max_score" gorm:"not null"`
	
	// Calculated fields (can be computed or stored)
	Percentage float64 `json:"percentage" gorm:"not null"`
	Grade      string  `json:"grade" gorm:"type:varchar(10)"` // A+, A, A-, B, C, F

	Feedback   string     `json:"feedback" gorm:"type:text"`
	ExamDate   time.Time  `json:"exam_date" gorm:"not null"`
	ResultDate time.Time  `json:"result_date" gorm:"not null"`
}
