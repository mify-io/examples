package router

import (
	"net/http"

	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/core"
)


type routerConfig struct {
	Middlewares []func(http.Handler) http.Handler
}

func NewRouterConfig(ctx *core.MifyServiceContext) *routerConfig {
	return &routerConfig {
		Middlewares: []func(http.Handler) http.Handler {
		// Add your middlewares here
		},
	}
}
