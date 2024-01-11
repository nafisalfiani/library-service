package entity

import "time"

type Payment struct {
	Id            int       `json:"id" gorm:"primaryKey"`
	UserId        int       `json:"user_id" gorm:"foreignKey"`
	Amount        float64   `json:"amount" gorm:"not null"`
	PaymentMethod string    `json:"payment_method" gorm:"not null"`
	PaymentDate   time.Time `json:"payment_date" gorm:"default:CURRENT_TIMESTAMP"`
	Rental        []Rental  `json:"rental,omitempty" gorm:"foreignKey:PaymentId"`
}
