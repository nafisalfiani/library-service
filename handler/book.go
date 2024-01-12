package handler

import (
	"library/entity"
	"library/errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ListBook returns list of books
//
// @Summary List books
// @Description returns list of books
// @Tags books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp{data=[]entity.Book}
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /books [get]
func (h *Handler) ListBook(c echo.Context) error {
	books, err := h.listBook()
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, books)
}

func (h *Handler) listBook() ([]entity.Book, error) {
	return h.book.List()
}

// GetBook returns specific books
//
// @Summary Get specific books
// @Description returns specific books
// @Tags books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "book id"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /books/{id} [get]
func (h *Handler) GetBook(c echo.Context) error {
	bookIdStr := c.Param("id")
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	book, err := h.book.Get(entity.Book{Id: bookId})
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, book)
}

// CreateBook creates new books
//
// @Summary Create books
// @Description Creates new books
// @Tags books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param book body entity.BookRequest true "book request"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /books [post]
func (h *Handler) CreateBook(c echo.Context) error {
	role := c.Request().Context().Value(contextKeyUserRole).(string)
	if role != "admin" {
		return h.httpError(c, errors.ErrUnauthorized, "you don't have access")
	}

	req := entity.BookRequest{}
	if err := c.Bind(&req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	book, err := h.createBook(req)
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusCreated, book)
}

func (h *Handler) createBook(req entity.BookRequest) (entity.Book, error) {
	var book entity.Book
	if err := h.validator.Struct(req); err != nil {
		return book, err
	}

	category, err := h.category.Get(entity.Category{Name: req.CategoryName})
	if errors.Is(err, errors.ErrNotFound) {
		category, err = h.category.Create(entity.Category{Name: req.CategoryName})
		if err != nil {
			return book, err
		}
	} else if err != nil {
		return book, err
	}

	newBook, err := h.book.Create(entity.Book{
		Name:              req.Name,
		Description:       req.Description,
		StockAvailability: req.StockAvailability,
		CategoryId:        category.Id,
		RentalCost:        req.RentalCost,
	})
	if err != nil {
		return book, err
	}

	return newBook, nil
}

// UpdateBook updates specific book
//
// @Summary Update book
// @Description Updates specific book
// @Tags books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "book id"
// @Param book body entity.BookRequest true "book request"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /books/{id} [put]
func (h *Handler) UpdateBook(c echo.Context) error {
	role := c.Request().Context().Value(contextKeyUserRole).(string)
	if role != "admin" {
		return h.httpError(c, errors.ErrUnauthorized, "you don't have access")
	}

	req := entity.BookRequest{}
	if err := c.Bind(&req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	bookIdStr := c.Param("id")
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	category, err := h.category.Get(entity.Category{Name: req.CategoryName})
	if errors.Is(err, errors.ErrNotFound) {
		category, err = h.category.Create(entity.Category{Name: req.CategoryName})
		if err != nil {
			return h.httpError(c, err)
		}
	} else if err != nil {
		return h.httpError(c, err)
	}

	book := entity.Book{
		Id:                bookId,
		Name:              req.Name,
		Description:       req.Description,
		StockAvailability: req.StockAvailability,
		CategoryId:        category.Id,
		RentalCost:        req.RentalCost,
	}
	newBook, err := h.book.Update(book)
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, newBook)
}

// ListBook delete specific books
//
// @Summary Delete book
// @Description Delete specific books
// @Tags books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "book id"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /books/{id} [delete]
func (h *Handler) DeleteBook(c echo.Context) error {
	role := c.Request().Context().Value(contextKeyUserRole).(string)
	if role != "admin" {
		return h.httpError(c, errors.ErrUnauthorized, "you don't have access")
	}

	bookIdStr := c.Param("id")
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.deleteBook(bookId); err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, nil)
}

func (h *Handler) deleteBook(bookId int) error {
	if err := h.book.Delete(bookId); err != nil {
		return err
	}

	return nil
}
