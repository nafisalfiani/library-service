package handler

import (
	"context"
	"fmt"
	"library/entity"
	"library/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateDeposit top up user deposit
//
// @Summary Top up user deposit
// @Description Top up logged in user deposit
// @Tags deposits
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param deposit body entity.DepositRequest true "deposit request"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /deposits [post]
func (h *Handler) CreateDeposit(c echo.Context) error {
	req := entity.DepositRequest{}
	if err := c.Bind(&req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	userId := c.Request().Context().Value(contextKeyUserId).(float64)
	depositReq := entity.DepositHistory{
		UserId:        int(userId),
		Amount:        req.Amount,
		Type:          entity.CreditTransaction,
		PaymentMethod: req.PaymentMethod,
	}
	if err := h.createDeposit(c.Request().Context(), depositReq); err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusCreated, nil)
}

func (h *Handler) createDeposit(ctx context.Context, req entity.DepositHistory) error {
	paymentReq := entity.XenditPayment{
		PaymentMethod: req.PaymentMethod,
	}
	if err := h.xendit.CreatePayment(ctx, paymentReq); err != nil {
		return err
	}

	// save xendit payment information to deposit/payment

	deposit, err := h.deposit.Create(req)
	if err != nil {
		return err
	}
	h.logger.Debug(fmt.Sprintf("%v", deposit))

	return nil
}

// GetDepositHistory returns logged in user deposit history
//
// @Summary Get deposit history
// @Description Get logged in user deposit history
// @Tags deposits
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /deposits/history [get]
func (h *Handler) GetDepositHistory(c echo.Context) error {
	userId := c.Request().Context().Value(contextKeyUserId).(float64)
	deposits, err := h.deposit.List(int(userId))
	if err != nil {
		return h.httpError(c, err)
	}

	resp := entity.DepositHistoryResponse{
		Total:     deposits.CalculateTotal(),
		Histories: deposits,
	}

	return h.httpSuccess(c, http.StatusOK, resp)
}
