package repository

import (
	"context"
	"fmt"
	"library/entity"

	"github.com/xendit/xendit-go/v4"
	xnd "github.com/xendit/xendit-go/v4/payment_request"
)

type xndClient struct {
	xnd *xendit.APIClient
}

type XenditInterface interface {
	CreatePayment(ctx context.Context, payment entity.XenditPayment) error
	ListPaymentMethod(ctx context.Context) ([]string, error)
}

// initXendit create xendit repository
func initXendit(xnd *xendit.APIClient) XenditInterface {
	return &xndClient{
		xnd: xnd,
	}
}

func (x *xndClient) CreatePayment(ctx context.Context, payment entity.XenditPayment) error {
	xndParam := xnd.NewPaymentRequestParameters(xnd.PaymentRequestCurrency("IDR"))
	xndParam.SetPaymentMethodId(payment.PaymentMethod)

	resp, r, err := x.xnd.PaymentRequestApi.CreatePaymentRequest(ctx).
		PaymentRequestParameters(*xndParam).
		Execute()
	if err != nil {
		fmt.Printf("Error when calling `PaymentRequestApi.CreatePaymentRequest``: %#v\n", err.Error())
		return nil
	}

	fmt.Printf("Full HTTP response: %#v\n", r)
	fmt.Printf("Response from `PaymentRequestApi.CreatePaymentRequest`: %#v\n", resp)

	return err
}

func (x *xndClient) ListPaymentMethod(ctx context.Context) ([]string, error) {
	req := x.xnd.PaymentMethodApi.GetAllPaymentMethods(ctx)

	resp, r, err := req.Execute()
	if err != nil {
		fmt.Printf("Error when calling `PaymentRequestApi.GetAllPaymentMethods``: %#v\n", err.Error())
		return nil, err
	}

	fmt.Printf("Full HTTP response: %#v\n", r)
	fmt.Printf("Response from `PaymentRequestApi.GetAllPaymentMethods`: %#v\n", resp)
	fmt.Printf("Response from `PaymentRequestApi.GetAllPaymentMethods` has more: %#v\n", *resp.HasMore)

	return nil, nil
}
