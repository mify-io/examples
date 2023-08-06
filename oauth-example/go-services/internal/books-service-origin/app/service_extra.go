// vim: set ft=go:

package app

import (
	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/domain"
	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/core"
)

type ServiceExtra struct {
	// Append your dependencies here
	BooksStorage domain.InMemoryBooks
}

func NewServiceExtra(ctx *core.MifyServiceContext) (*ServiceExtra, error) {
	// Here you can do your custom service initialization, prepare dependencies
	extra := &ServiceExtra{
		// Here you can initialize your dependencies
	}
	return extra, nil
}
