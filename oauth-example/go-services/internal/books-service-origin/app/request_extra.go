package app

import (
	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/core"
)

type RequestExtra struct {
	// Append your dependencies here
}

func NewRequestExtra(ctx *core.MifyServiceContext) (*RequestExtra, error) {
	// Here you can do your custom service initialization, prepare dependencies
	extra := &RequestExtra{
		// Here you can initialize your dependencies
	}
	return extra, nil
}
