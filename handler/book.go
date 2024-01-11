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
	books, err := h.book.List()
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, books)
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
	req := entity.BookRequest{}
	if err := c.Bind(&req); err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
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

	book, err := h.book.Create(entity.Book{
		Name:              req.Name,
		Description:       req.Description,
		StockAvailability: req.StockAvailability,
		CategoryId:        category.Id,
		RentalCost:        req.RentalCost,
	})
	if err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusCreated, book)
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
// @Router /books [get]
func (h *Handler) UpdateBook(c echo.Context) error {
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
// @Router /books [get]
func (h *Handler) DeleteBook(c echo.Context) error {
	bookIdStr := c.Param("id")
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		return h.httpError(c, errors.ErrBadRequest, err.Error())
	}

	if err := h.book.Delete(bookId); err != nil {
		return h.httpError(c, err)
	}

	return h.httpSuccess(c, http.StatusOK, nil)
}
