package entity

import "time"

type Rental struct {
	Id            int            `json:"id" gorm:"primaryKey"`
	UserId        int            `json:"user_id" gorm:"foreignKey"`
	PaymentId     int            `json:"payment_id" gorm:"foreignKey"`
	RentalDate    time.Time      `json:"rental_date" gorm:"default:CURRENT_TIMESTAMP"`
	Payment       *Payment       `json:"payment" gorm:"foreignKey:PaymentId"`
	RentalDetails []RentalDetail `json:"rental_details,omitempty" gorm:"foreignKey:RentalId"`
}

type RentalDetail struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	RentalId   int       `json:"rental_id" gorm:"foreignKey"`
	BookId     int       `json:"book_id"`
	ReturnDate time.Time `json:"return_date,omitempty"`
	Returned   bool      `json:"returned" gorm:"default:false"`
}
