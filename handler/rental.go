package handler

import (
	"fmt"
	"library/entity"
	"library/errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// GetActiveRental get active rental
//
// @Summary Get active rental
// @Description Get active rental of logged in user
// @Tags rentals
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp{data=[]entity.Rental}
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /rental/active [get]
func (h *Handler) GetActiveRental(c echo.Context) error {
	rentals, err := h.rental.ListOutstanding()
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, rentals)
}

// GetRentalHistory get closed rental
//
// @Summary Get closed rental
// @Description Get closed rental of logged in user
// @Tags rentals
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp{data=[]entity.Rental}
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /rental/closed [get]
func (h *Handler) GetRentalHistory(c echo.Context) error {
	rentals, err := h.rental.ListHistory()
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, rentals)
}

// CreateRental allows user to rent a book
//
// @Summary Create a rental
// @Description Get active rental of logged in user
// @Tags rentals
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param rental body entity.RentalRequest true "rental request"
// @Success 200 {object} entity.HttpResp{data=entity.Rental}
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /rental [post]
func (h *Handler) CreateRental(c echo.Context) error {
	req := entity.RentalRequest{}
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
	rental := entity.Rental{
		UserId:     int(userId),
		Status:     "active",
		RentalDate: time.Now(),
	}
	for i := range req.Books {
		rental.RentalDetails = append(rental.RentalDetails, entity.RentalDetail{
			BookId:         req.Books[i].BookId,
			ReturnDate:     rental.RentalDate.Add(time.Duration(24*req.Books[i].RentalDuration) * time.Hour),
			RentalDuration: req.Books[i].RentalDuration,
		})
	}

	newRental, err := h.rental.Create(rental, req.PaymentMethod)
	if err != nil {
		return h.httpError(c, err)
	}

	paymentReq := entity.XenditPaymentRequest{
		PaymentId:          fmt.Sprintf(":rental-payment:%v", newRental.PaymentId),
		Amount:             newRental.Payment.Amount,
		PaymentMethod:      req.PaymentMethod,
		Currency:           &entity.IdrCurrency,
		InvoiceName:        &user.FullName,
		InvoiceEmail:       &user.Email,
		InvoiceDescription: &entity.DescriptionDepositSaldo,
		InvoiceExpiry:      &entity.InvoiceExpiry,
	}
	for i := range newRental.RentalDetails {
		paymentReq.Items = append(paymentReq.Items, entity.PaymentItems{
			Name:     newRental.RentalDetails[i].BookName,
			Price:    float32(newRental.RentalDetails[i].RentalCost),
			Quantity: float32(newRental.RentalDetails[i].RentalDuration),
		})
	}

	resp, err := h.xendit.CreatePayment(c.Request().Context(), paymentReq)
	if err != nil {
		return h.httpError(c, err)
	}

	newRental.Payment.XenditPaymentId = resp.XenditPaymentId
	newRental.Payment.XenditPaymentUrl = resp.InvoiceUrl
	newRental.Payment.PaymentMethod = entity.PaymentMethodWaiting
	newRental.Payment.Status = entity.InvoiceStatusPending
	newRental.Payment.Type = entity.PaymentTypeRentalPayment
	newPayment, err := h.payment.Update(*newRental.Payment)
	if err != nil {
		return h.httpError(c, err)
	}
	h.logger.Debug(fmt.Sprintf("%v", newPayment))

	mail := entity.Mail{
		From:    "library-service@mail.com",
		To:      user.Email,
		Subject: "Library Rental Payment",
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
