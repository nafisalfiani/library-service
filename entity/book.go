package entity

type Category struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null;unique"`
	Books []Book `json:"books,omitempty" gorm:"foreignKey:CategoryId"`
}

type Book struct {
	Id                int            `json:"id" gorm:"primaryKey"`
	Name              string         `json:"name" gorm:"not null"`
	Description       string         `json:"description" gorm:"not null"`
	StockAvailability int            `json:"stock_availability" gorm:"not null"`
	CategoryId        int            `json:"category_id" gorm:"foreignKey"`
	RentalCost        float64        `json:"rental_cost" gorm:"not null"`
	RentalDetails     []RentalDetail `json:"rental_details,omitempty" gorm:"foreignKey:BookId"`
}

type BookRequest struct {
	Name              string  `json:"name" validate:"required"`
	Description       string  `json:"description" validate:"required"`
	StockAvailability int     `json:"stock_availability" validate:"required"`
	CategoryName      string  `json:"category_name" validate:"required"`
	RentalCost        float64 `json:"rental_cost" validate:"required"`
}
