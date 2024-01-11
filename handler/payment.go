package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ListPaymentMethods lists available payment methods
//
// @Summary List payment methods
// @Description Lists available payment methods
// @Tags payments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /payments/methods [get]
func (h *Handler) ListPaymentMethods(c echo.Context) error {
	methods, err := h.xendit.ListPaymentMethod(c.Request().Context())
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, methods)
}
