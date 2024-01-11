package entity

type User struct {
	Id            int       `json:"id" gorm:"primaryKey"`
	Email         string    `json:"email" gorm:"not null;unique"`
	Username      string    `json:"username" gorm:"not null;unique"`
	FullName      string    `json:"full_name" gorm:"not null"`
	Password      string    `json:"-" gorm:"not null"`
	DepositAmount float64   `json:"deposit_amount" gorm:"default:0"`
	Role          string    `json:"role" gorm:"not null"`
	Payment       []Payment `json:"payments,omitempty" gorm:"foreignKey:UserId"`
	Rental        []Rental  `json:"rentals,omitempty" gorm:"foreignKey:UserId"`
}

type RegisterRequest struct {
	Email         string  `json:"email" validate:"required,email"`
	Username      string  `json:"username" validate:"required"`
	FullName      string  `json:"full_name" validate:"required"`
	Password      string  `json:"password" validate:"required"`
	DepositAmount float64 `json:"deposit_amount" validate:"required"`
	Role          string  `json:"role" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
