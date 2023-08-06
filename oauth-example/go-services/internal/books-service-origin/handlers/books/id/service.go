package api_books_id

import (
	"errors"
	"net/http"

	openapi "example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/api"
	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/apputil"
	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/core"
)

type BooksIdApiService struct{}

// NewBooksIdApiService creates a default api service
func NewBooksIdApiService() openapi.BooksIdApiServicer {
	return &BooksIdApiService{}
}

// BooksIdGet - get book by id
func (s *BooksIdApiService) BooksIdGet(ctx *core.MifyRequestContext, id int64) (openapi.ServiceResponse, error) {
	book, err := apputil.GetServiceExtra(ctx.ServiceContext()).BooksStorage.Find(id)

	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("error while creating book")
	}

	return openapi.Response(200, openapi.Book{
		Id:   book.Id,
		Name: book.Name,
	}), nil
}
