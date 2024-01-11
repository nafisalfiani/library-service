package entity

import "time"

type Rental struct {
	Id            int            `json:"id" gorm:"primaryKey"`
	UserId        int            `json:"user_id" gorm:"foreignKey"`
	PaymentId     int            `json:"payment_id" gorm:"foreignKey"`
	RentalDate    time.Time      `json:"rental_date" gorm:"default:CURRENT_TIMESTAMP"`
	Payment       *Payment       `json:"payment,omitempty" gorm:"foreignKey:PaymentId"`
	Status        string         `json:"status"`
	RentalDetails []RentalDetail `json:"rental_details,omitempty" gorm:"foreignKey:RentalId"`
}

type RentalRequest struct {
	PaymentMethod string                `json:"payment_method" validate:"required"`
	Books         []RentalRequestDetail `json:"books" validate:"required"`
}

type RentalRequestDetail struct {
	BookId         int `json:"book_id" validate:"required"`
	RentalDuration int `json:"rental_duration" validate:"required"`
}

type RentalDetail struct {
	Id             int       `json:"id" gorm:"primaryKey"`
	RentalId       int       `json:"rental_id" gorm:"foreignKey"`
	BookId         int       `json:"book_id"`
	BookName       string    `json:"book_name" gorm:"-"`
	ReturnDate     time.Time `json:"return_date,omitempty"`
	RentalCost     float64   `json:"rental_cost" gorm:"-"`
	RentalDuration int       `json:"rental_duration" gorm:"-"`
	Returned       bool      `json:"returned" gorm:"default:false"`
}
