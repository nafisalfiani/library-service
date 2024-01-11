package handler

import (
	"context"
	"fmt"
	"library/entity"
	"library/errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type contextKey string

const (
	contextKeyUserId   contextKey = "user_id"
	contextKeyUserRole contextKey = "user_role"
	contextKeyUsername contextKey = "username"
)

// Register allow new user to register their account info
//
// @Summary Register new user
// @Description Allow new user to register their account info
// @Tags users
// @Accept json
// @Produce json
// @Param register body entity.RegisterRequest true "register request"
// @Success 200 {object} entity.HttpResp{data=entity.User}
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /users/register [post]
func (h *Handler) Register(c echo.Context) error {
	req := entity.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return h.httpError(c, errors.ErrInternalServerError, err.Error())
	}

	createReq := entity.User{
		Username:      req.Username,
		Email:         req.Email,
		FullName:      req.FullName,
		Password:      hashedPassword,
		DepositAmount: req.DepositAmount,
		Role:          req.Role,
	}
	newUser, err := h.user.Create(createReq)
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, newUser)
}

// Login allow existing user to login
//
// @Summary Login existing user
// @Description Allow existing user to login
// @Tags users
// @Accept json
// @Produce json
// @Param login body entity.LoginRequest true "login request"
// @Success 200 {object} entity.HttpResp{data=entity.LoginResponse}
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /users/login [post]
func (h *Handler) Login(c echo.Context) error {
	loginReq := entity.LoginRequest{}
	if err := c.Bind(&loginReq); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(loginReq); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	user, err := h.user.Get(loginReq.Username)
	if err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := checkPasswordHash(user.Password, loginReq.Password); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	token, err := h.createToken(user)
	if err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	return h.httpSuccess(c, http.StatusOK, token)
}

func (h *Handler) Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if err := h.checkToken(c, tokenString); err != nil {

			return h.httpError(c, errors.ErrUnauthorized, err.Error())
		}

		return next(c)
	}
}

func (h *Handler) checkToken(c echo.Context, tokenString string) error {
	if tokenString == "" {
		return fmt.Errorf("missing token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.Auth.SecretKey), nil
	})
	if err != nil {
		return fmt.Errorf("failed to parse token")
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	// Accessing claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid token")
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	currentTime := time.Now()

	if currentTime.After(expirationTime) {
		return fmt.Errorf("token expired")
	}

	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, contextKeyUserId, claims["user_id"])
	ctx = context.WithValue(ctx, contextKeyUserRole, claims["user_role"])
	ctx = context.WithValue(ctx, contextKeyUsername, claims["username"])

	c.SetRequest(c.Request().WithContext(ctx))

	return nil
}

func (h *Handler) createToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.Id,
		"user_role": user.Role,
		"username":  user.Username,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.config.Auth.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
