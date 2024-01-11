package repository

import (
	"context"
	"fmt"
	"library/entity"

	"github.com/xendit/xendit-go/v4"
	"github.com/xendit/xendit-go/v4/invoice"
)

type xndClient struct {
	xnd *xendit.APIClient
}

type XenditInterface interface {
	CreatePayment(ctx context.Context, payment entity.XenditPaymentRequest) (entity.XenditPaymentResponse, error)
	CheckPayment(ctx context.Context, payment entity.XenditPaymentResponse) (entity.XenditPaymentResponse, error)
}

// initXendit create xendit repository
func initXendit(xnd *xendit.APIClient) XenditInterface {
	return &xndClient{
		xnd: xnd,
	}
}

func (x *xndClient) CreatePayment(ctx context.Context, payment entity.XenditPaymentRequest) (entity.XenditPaymentResponse, error) {
	req := x.xnd.InvoiceApi.CreateInvoice(ctx).CreateInvoiceRequest(invoice.CreateInvoiceRequest{
		ExternalId:      fmt.Sprintf("library-service%v", payment.PaymentId),
		Amount:          payment.Amount,
		Description:     payment.InvoiceDescription,
		InvoiceDuration: payment.InvoiceExpiry,
		Customer: &invoice.CustomerObject{
			GivenNames: *invoice.NewNullableString(payment.InvoiceName),
			Email:      *invoice.NewNullableString(payment.InvoiceEmail),
		},
		Currency: payment.Currency,
		Items:    x.invoiceItems(payment.Items),
	})

	xenditResp, _, err := req.Execute()
	if err != nil {
		fmt.Printf("Error when calling `PaymentRequestApi.CreatePaymentRequest``: %#v\n", err.Error())
		return entity.XenditPaymentResponse{}, err
	}

	resp := entity.XenditPaymentResponse{
		XenditPaymentId:   *xenditResp.Id,
		PaymentId:         xenditResp.ExternalId,
		InvoiceExpiryDate: xenditResp.ExpiryDate,
		InvoiceStatus:     xenditResp.Status.String(),
		InvoiceAmount:     xenditResp.Amount,
		InvoiceUrl:        xenditResp.InvoiceUrl,
	}

	if xenditResp.PaymentMethod != nil && xenditResp.PaymentMethod.IsValid() {
		resp.PaymentMethod = xenditResp.PaymentMethod.String()
	}

	return resp, nil
}

func (x *xndClient) invoiceItems(items []entity.PaymentItems) []invoice.InvoiceItem {
	invoiceItems := []invoice.InvoiceItem{}
	for i := range items {
		invoiceItems = append(invoiceItems, invoice.InvoiceItem{
			Name:     items[i].Name,
			Price:    items[i].Price,
			Quantity: items[i].Quantity,
		})
	}

	return invoiceItems
}

func (x *xndClient) CheckPayment(ctx context.Context, payment entity.XenditPaymentResponse) (entity.XenditPaymentResponse, error) {
	req := x.xnd.InvoiceApi.GetInvoiceById(ctx, payment.XenditPaymentId)
	xenditResp, _, err := req.Execute()
	if err != nil {
		fmt.Printf("Error when calling `PaymentRequestApi.CreatePaymentRequest``: %#v\n", err.Error())
		return entity.XenditPaymentResponse{}, err
	}

	resp := entity.XenditPaymentResponse{
		XenditPaymentId:   *xenditResp.Id,
		PaymentId:         xenditResp.ExternalId,
		InvoiceExpiryDate: xenditResp.ExpiryDate,
		InvoiceStatus:     xenditResp.Status.String(),
		InvoiceAmount:     xenditResp.Amount,
		InvoiceUrl:        xenditResp.InvoiceUrl,
	}

	if xenditResp.PaymentMethod != nil && xenditResp.PaymentMethod.IsValid() {
		resp.PaymentMethod = xenditResp.PaymentMethod.String()
	}

	return resp, nil
}
