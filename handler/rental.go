package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetOutstandingRental(c echo.Context) error {
	return h.httpSuccess(c, http.StatusOK, nil)
}

func (h *Handler) GetRentalHistory(c echo.Context) error {
	return h.httpSuccess(c, http.StatusOK, nil)
}

func (h *Handler) CreateRental(c echo.Context) error {
	return h.httpSuccess(c, http.StatusCreated, nil)
}
