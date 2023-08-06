package api_books

import (
	"errors"
	"net/http"

	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/domain"
	openapi "example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/api"
	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/apputil"
	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/core"
)

type BooksApiService struct{}

// NewBooksApiService creates a default api service
func NewBooksApiService() openapi.BooksApiServicer {
	return &BooksApiService{}
}

// BooksGet - get all books
func (s *BooksApiService) BooksGet(ctx *core.MifyRequestContext) (openapi.ServiceResponse, error) {
	books, err := apputil.GetServiceExtra(ctx.ServiceContext()).BooksStorage.GetAll()
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("error while loading books")
	}

	var booksResponse []openapi.Book
	for _, book := range books {
		booksResponse = append(booksResponse, openapi.Book{
			Id:   book.Id,
			Name: book.Name,
		})
	}

	return openapi.Response(200, openapi.GetBooksResponse{
		Books: booksResponse,
	}), nil
}

// BooksPut - create new book
func (s *BooksApiService) BooksPut(ctx *core.MifyRequestContext, createBookRequest openapi.CreateBookRequest) (openapi.ServiceResponse, error) {
	book, err := apputil.GetServiceExtra(ctx.ServiceContext()).BooksStorage.Add(domain.CreateBookParams{Name: createBookRequest.Name})
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("error while creating book")
	}

	return openapi.Response(200, openapi.Book{
		Id:   book.Id,
		Name: book.Name,
	}), nil
}
