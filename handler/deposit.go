package handler

import (
	"context"
	"fmt"
	"library/entity"
	"library/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TopUpDeposit top up user deposit
//
// @Summary Top up user deposit
// @Description Top up logged in user deposit
// @Tags deposits
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param deposit body entity.PaymentRequest true "deposit request"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /deposits [post]
func (h *Handler) TopUpDeposit(c echo.Context) error {
	req := entity.PaymentRequest{}
	if err := c.Bind(&req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	user, err := h.user.Get(c.Request().Context().Value(contextKeyUsername).(string))
	if err != nil {
		return err
	}

	userId := c.Request().Context().Value(contextKeyUserId).(float64)
	depositReq := entity.Payment{
		UserId:        int(userId),
		Amount:        req.Amount,
		PaymentMethod: "UNDECIDED",
		Status:        entity.InvoiceStatusPending,
		Type:          "deposit-saldo",
	}
	resp, err := h.createPayment(c.Request().Context(), depositReq, user)
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusCreated, resp)
}

func (h *Handler) createPayment(ctx context.Context, req entity.Payment, user entity.User) (entity.XenditPaymentResponse, error) {
	payment, err := h.payment.Create(req)
	if err != nil {
		return entity.XenditPaymentResponse{}, err
	}

	desc := "Deposit Saldo"
	exp := "86400"
	paymentReq := entity.XenditPaymentRequest{
		Amount:             req.Amount,
		PaymentMethod:      req.PaymentMethod,
		PaymentId:          fmt.Sprintf(":deposit-saldo:%v", payment.Id),
		Currency:           &entity.IdrCurrency,
		InvoiceName:        &user.FullName,
		InvoiceEmail:       &user.Email,
		InvoiceDescription: &desc,
		InvoiceExpiry:      &exp,
		Items: []entity.PaymentItems{
			{
				Name:     "Deposit Saldo",
				Price:    float32(req.Amount),
				Quantity: 1,
			},
		},
	}
	resp, err := h.xendit.CreatePayment(ctx, paymentReq)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
