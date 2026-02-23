package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod string

const (
	PaymentMethodBkash  PaymentMethod = "bkash"
	PaymentMethodNagad  PaymentMethod = "nagad"
	PaymentMethodRocket PaymentMethod = "rocket"
	PaymentMethodBank   PaymentMethod = "bank"
	PaymentMethodCash   PaymentMethod = "cash"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

type Payment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	UserID uint `json:"user_id" gorm:"not null"`
	User   User `json:"user,omitempty" gorm:"foreignKey:UserID"`

	Amount        float64       `json:"amount" gorm:"not null"`
	Method        PaymentMethod `json:"method" gorm:"type:varchar(20);not null"`
	Status        PaymentStatus `json:"status" gorm:"type:varchar(20);default:'pending'"`
	TransactionID string        `json:"transaction_id" gorm:"uniqueIndex"`
	Reference     string        `json:"reference"` // Additional reference info

	PaidAt      *time.Time `json:"paid_at,omitempty"`
	Description string     `json:"description"`

	// Related entities
	CourseID       *uint         `json:"course_id,omitempty"`
	Course         *Course       `json:"course,omitempty" gorm:"foreignKey:CourseID"`
	SubscriptionID *uint         `json:"subscription_id,omitempty"`
	Subscription   *Subscription `json:"subscription,omitempty" gorm:"foreignKey:SubscriptionID"`
}
