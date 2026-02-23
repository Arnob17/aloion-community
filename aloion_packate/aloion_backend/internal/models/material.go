package models

import (
	"time"

	"gorm.io/gorm"
)

type MaterialType string

const (
	MaterialTypeVideo MaterialType = "video"
	MaterialTypePDF   MaterialType = "pdf"
	MaterialTypeLink  MaterialType = "link"
	MaterialTypeText  MaterialType = "text"
)

type CourseMaterial struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	CourseID uint   `json:"course_id" gorm:"not null"`
	Course   Course `json:"course,omitempty" gorm:"foreignKey:CourseID"`

	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description" gorm:"type:text"`
	Type        MaterialType `json:"type" gorm:"type:varchar(20);not null"`
	ContentURL  string       `json:"content_url" gorm:"not null"`
	Duration    int          `json:"duration"`               // in minutes (for videos)
	Order       int          `json:"order" gorm:"default:0"` // Order in course sequence

	IsFree bool `json:"is_free" gorm:"default:false"` // Can be accessed without enrollment
}
