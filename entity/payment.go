package entity

import "time"

type Payment struct {
	Id                  int       `json:"id" gorm:"primaryKey"`
	UserId              int       `json:"user_id" gorm:"foreignKey"`
	Amount              float64   `json:"amount" gorm:"not null"`
	PaymentMethod       string    `json:"payment_method" gorm:"not null"`
	PaymentDate         time.Time `json:"payment_date" gorm:"default:CURRENT_TIMESTAMP"`
	XenditPaymentId     string    `json:"xendit_payment_id"`
	XenditPaymentUrl    string    `json:"xendit_payment_url"`
	Type                string    `json:"type" gorm:"not null"`
	Status              string    `json:"status" gorm:"not null"`
	ShouldUpdateDeposit bool      `json:"-" gorm:"-"`
}

type PaymentRequest struct {
	Amount float64 `json:"amount" gorm:"not null"`
}
