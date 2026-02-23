package models

import (
	"time"

	"gorm.io/gorm"
)

type SubscriptionPlan string

const (
	PlanMonthly SubscriptionPlan = "monthly"
	PlanYearly  SubscriptionPlan = "yearly"
)

type SubscriptionStatus string

const (
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusExpired   SubscriptionStatus = "expired"
	SubscriptionStatusCancelled SubscriptionStatus = "cancelled"
)

type Subscription struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	UserID uint `json:"user_id" gorm:"not null"`
	User   User `json:"user,omitempty" gorm:"foreignKey:UserID"`

	Plan   SubscriptionPlan   `json:"plan" gorm:"type:varchar(20);not null"`
	Status SubscriptionStatus `json:"status" gorm:"type:varchar(20);default:'active'"`
	Price  float64            `json:"price" gorm:"not null"`

	StartDate time.Time  `json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	RenewsAt  *time.Time `json:"renews_at,omitempty"`

	// Payment reference
	PaymentID *uint    `json:"payment_id,omitempty"`
	Payment   *Payment `json:"payment,omitempty" gorm:"foreignKey:PaymentID"`
}
