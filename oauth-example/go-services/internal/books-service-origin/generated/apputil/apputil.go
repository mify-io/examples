// THIS FILE IS AUTOGENERATED, DO NOT EDIT
// Generated by mify
// vim: set ft=go:

package apputil

import (
	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/app"
	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/core"
)

func GetServiceExtra(ctx *core.MifyServiceContext) *app.ServiceExtra {
	return ctx.ServiceExtra().(*app.ServiceExtra)
}

func GetRequestExtra(ctx *core.MifyRequestContext) *app.RequestExtra {
	return ctx.RequestExtra().(*app.RequestExtra)
}