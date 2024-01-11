package entity

import "strings"

const (
	PaymentMethodEwallet        = "EWALLET"
	PaymentMethodDirectDebit    = "DIRECT_DEBIT"
	PaymentMethodCard           = "CARD"
	PaymentMethodVirtualAccount = "VIRTUAL_ACCOUNT"
	PaymentMethodOverTheCounter = "OVER_THE_COUNTER"
	PaymentMethodQrCode         = "QR_CODE"
	PaymentMethodInvalid        = "INVALID"
)

type XenditPayment struct {
	PaymentMethod string
}

func (x XenditPayment) GetPaymentMethod() string {
	switch strings.ToLower(x.PaymentMethod) {
	case strings.ToLower(PaymentMethodEwallet):
		return PaymentMethodEwallet
	case strings.ToLower(PaymentMethodDirectDebit):
		return PaymentMethodDirectDebit
	case strings.ToLower(PaymentMethodCard):
		return PaymentMethodCard
	case strings.ToLower(PaymentMethodVirtualAccount):
		return PaymentMethodVirtualAccount
	case strings.ToLower(PaymentMethodOverTheCounter):
		return PaymentMethodOverTheCounter
	case strings.ToLower(PaymentMethodQrCode):
		return PaymentMethodQrCode
	}

	return PaymentMethodInvalid
}
