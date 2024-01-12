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
		PaymentMethod: entity.PaymentMethodWaiting,
		Status:        entity.InvoiceStatusPending,
		Type:          entity.PaymentTypeDepositSaldo,
	}
	resp, err := h.createPayment(c.Request().Context(), depositReq, user)
	if err != nil {
		return h.httpError(c, err)
	}

	depositReq.XenditPaymentId = resp.XenditPaymentId
	depositReq.XenditPaymentUrl = resp.InvoiceUrl
	newPayment, err := h.payment.Update(depositReq)
	if err != nil {
		return h.httpError(c, err)
	}
	h.logger.Debug(fmt.Sprintf("%#v", newPayment))

	mail := entity.Mail{
		From:    "library-service@mail.com",
		To:      user.Email,
		Subject: "Library Deposit Saldo",
	}
	mail.Body, err = mail.ParseHtml(resp)
	if err != nil {
		h.logger.Error(fmt.Sprintf("failed to parse template with error: %v", err))
	} else {
		if err := h.mailer.Send(mail); err != nil {
			h.logger.Error(fmt.Sprintf("email not sent with error: %v", err))
		}
	}

	return h.httpSuccess(c, http.StatusCreated, resp)
}

func (h *Handler) createPayment(ctx context.Context, req entity.Payment, user entity.User) (entity.XenditPaymentResponse, error) {
	payment, err := h.payment.Create(req)
	if err != nil {
		return entity.XenditPaymentResponse{}, err
	}

	paymentReq := entity.XenditPaymentRequest{
		PaymentId:          fmt.Sprintf(":deposit-saldo:%v", payment.Id),
		Amount:             req.Amount,
		PaymentMethod:      req.PaymentMethod,
		Currency:           &entity.IdrCurrency,
		InvoiceName:        &user.FullName,
		InvoiceEmail:       &user.Email,
		InvoiceDescription: &entity.DescriptionDepositSaldo,
		InvoiceExpiry:      &entity.InvoiceExpiry,
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
