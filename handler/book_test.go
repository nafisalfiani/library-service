package handler

import (
	"library/entity"
	"library/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createBook(t *testing.T) {
	expectedBook := entity.Book{
		Name:              "Path to Go",
		Description:       "Explain path to Go mastery",
		StockAvailability: 10,
		CategoryId:        1,
		RentalCost:        1000,
	}

	mockCat := mocks.NewCategoryInterface(t)
	mockCat.On("Get", entity.Category{Name: "technology"}).Return(entity.Category{Id: 1, Name: "technology"}, nil)

	mockBook := mocks.NewBookInterface(t)
	mockBook.On("Create", entity.Book{
		Name:              "Path to Go",
		Description:       "Explain path to Go mastery",
		StockAvailability: 10,
		CategoryId:        1,
		RentalCost:        1000,
	}).Return(expectedBook, nil)

	h := &Handler{
		validator: validator.New(validator.WithRequiredStructEnabled()),
		book:      mockBook,
		category:  mockCat,
	}

	got, err := h.createBook(entity.BookRequest{
		Name:              "Path to Go",
		Description:       "Explain path to Go mastery",
		StockAvailability: 10,
		CategoryName:      "technology",
		RentalCost:        1000,
	})
	assert.Nil(t, err)
	assert.Equal(t, expectedBook.Name, got.Name, "book name should be equal")
}

func TestHandler_deleteBook(t *testing.T) {
	mockBook := mocks.NewBookInterface(t)
	mockBook.On("Delete", 1).Return(nil)

	h := &Handler{
		book: mockBook,
	}

	err := h.deleteBook(1)
	assert.Nil(t, err)
}

func TestHandler_listBook(t *testing.T) {
	mockBook := mocks.NewBookInterface(t)
	mockBook.On("List").Return([]entity.Book{
		{
			Name:              "Path to Go",
			Description:       "Explain path to Go mastery",
			StockAvailability: 10,
			CategoryId:        1,
			RentalCost:        1000,
		},
		{
			Name:              "Gordon Ramsay Cookbook",
			Description:       "Collection of Gordon Ramsay Techniques",
			StockAvailability: 20,
			CategoryId:        2,
			RentalCost:        500,
		},
	}, nil)

	h := &Handler{
		book: mockBook,
	}

	books, err := h.listBook()
	assert.Equal(t, 2, len(books), "should got 2 entry of books")

	assert.Nil(t, err)
}
