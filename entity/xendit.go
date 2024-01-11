package entity

import (
	"strconv"
	"strings"
	"time"
)

var (
	IdrCurrency = "IDR"
)

const (
	InvoiceStatusPending = "PENDING"
	InvoiceStatusPaid    = "PAID"
)

type XenditPaymentRequest struct {
	PaymentId          string
	Amount             float64
	PaymentMethod      string
	Currency           *string
	InvoiceDescription *string
	InvoiceExpiry      *string
	InvoiceName        *string
	InvoiceEmail       *string
	Items              []PaymentItems
}

type PaymentItems struct {
	Name     string
	Price    float32
	Quantity float32
}

type XenditPaymentResponse struct {
	XenditPaymentId   string    `json:"xendit_payment_id"`
	PaymentId         string    `json:"payment_id"`
	InvoiceExpiryDate time.Time `json:"expiry_date"`
	InvoiceStatus     string    `json:"status"`
	InvoiceAmount     float64   `json:"amount"`
	InvoiceUrl        string    `json:"url"`
	PaymentMethod     string    `json:"payment_method""`
}

type XenditCheckPayment struct {
	XenditPaymentId string `json:"xendit_payment_id"`
	PaymentId       string `json:"payment_id"`
}

func (x *XenditPaymentResponse) GetPaymentId() (paymentType string, paymentId int, err error) {
	res := strings.Split(x.PaymentId, ":")
	paymentType = res[1]
	paymentId, err = strconv.Atoi(res[2])
	return
}
