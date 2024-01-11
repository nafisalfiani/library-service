package entity

const (
	CreditTransaction = "credit"
	DebitTransaction  = "debit"
)

type DepositHistory struct {
	Id            int     `json:"id" gorm:"primaryKey"`
	UserId        int     `json:"user_id" gorm:"foreignKey"`
	Amount        float64 `json:"amount" gorm:"not null"`
	Type          string  `json:"type" gorm:"not null"`
	PaymentMethod string  `json:"payment_method" gorm:"-"`
}

type DepositHistories []DepositHistory

func (d DepositHistories) CalculateTotal() float64 {
	var total float64
	for i := range d {
		if d[i].Type == "credit" {
			total = total + d[i].Amount
		} else {
			total = total - d[i].Amount
		}
	}

	return total
}

type DepositRequest struct {
	Amount        float64 `json:"amount" validate:"required"`
	PaymentMethod string  `json:"payment_method" validate:"required"`
}

type DepositHistoryResponse struct {
	Total     float64          `json:"total"`
	Histories DepositHistories `json:"histories"`
}
