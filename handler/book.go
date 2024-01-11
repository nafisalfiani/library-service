package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ListBook(c echo.Context) error {
	return h.httpSuccess(c, http.StatusOK, nil)
}

func (h *Handler) GetBook(c echo.Context) error {
	return h.httpSuccess(c, http.StatusOK, nil)
}

func (h *Handler) CreateBook(c echo.Context) error {
	return h.httpSuccess(c, http.StatusCreated, nil)
}

func (h *Handler) UpdateBook(c echo.Context) error {
	return h.httpSuccess(c, http.StatusOK, nil)
}

func (h *Handler) DeleteBook(c echo.Context) error {
	return h.httpSuccess(c, http.StatusOK, nil)
}
