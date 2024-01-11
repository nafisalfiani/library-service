package handler

import (
	"fmt"
	"library/entity"
	"library/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RefreshPaymentStatus check and refresh payment status if applicable
//
// @Summary Refresh payment status
// @Description Check and refresh payment status if applicable
// @Tags payments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param payment body entity.XenditCheckPayment true "payment request"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /payments [post]
func (h *Handler) RefreshPaymentStatus(c echo.Context) error {
	req := entity.XenditCheckPayment{}
	if err := c.Bind(&req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	paymentReq := entity.XenditPaymentResponse{
		PaymentId:       req.PaymentId,
		XenditPaymentId: req.XenditPaymentId,
	}
	resp, err := h.xendit.CheckPayment(c.Request().Context(), paymentReq)
	if err != nil {
		return h.httpError(c, err)
	}

	// only update the payment info if invoice status is not pending
	if resp.InvoiceStatus != entity.InvoiceStatusPending {
		_, paymentId, err := resp.GetPaymentId()
		if err != nil {
			return h.httpError(c, err)
		}
		payment, err := h.payment.Get(entity.Payment{Id: paymentId})
		if err != nil {
			return h.httpError(c, err)
		}

		payment.Status = entity.InvoiceStatusPaid
		payment.PaymentMethod = resp.PaymentMethod
		newPayment, err := h.payment.Update(payment)
		if err != nil {
			return h.httpError(c, err)
		}

		h.logger.Debug(fmt.Sprintf("%#v", newPayment))
	}

	return h.httpSuccess(c, http.StatusCreated, resp)
}
