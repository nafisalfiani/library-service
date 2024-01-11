package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUsers returns logged in user detail
//
// @Summary Fetch user detail
// @Description Get logged in user detail
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp{data=entity.User}
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /users [get]
func (h *Handler) GetUser(c echo.Context) error {
	loggedInUsername := c.Request().Context().Value(contextKeyUsername).(string)
	user, err := h.user.Get(loggedInUsername)
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, user)
}
