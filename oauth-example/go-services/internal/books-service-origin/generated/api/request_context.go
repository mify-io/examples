// THIS FILE IS AUTOGENERATED, DO NOT EDIT
// Generated by mify via OpenAPI Generator

// vim: set ft=go:
package openapi

// vim: set ft=go:

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/api/public"
	"example.com/namespace/oauth-example/go-services/internal/books-service-origin/generated/core"
)

type ctxKeyMifyContextBuilder int

const MifyContextBuilderField ctxKeyMifyContextBuilder = 0

func requestContextInitBuilder(sc *core.MifyServiceContext) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			mifyCtxBuilder := core.NewMifyRequestContextBuilder(sc)
			mifyCtxBuilder.SetRequestID(middleware.GetReqID(r.Context()))
			mifyCtxBuilder.SetProtocol(r.Proto)
			mifyCtxBuilder.SetURLPath(r.URL.Path)
			ctx := context.WithValue(r.Context(), MifyContextBuilderField, mifyCtxBuilder)

			next.ServeHTTP(ww, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func requestContextBuild(requestExtraFactory core.RequestExtraFactory) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			builder := getMifyRequestContextBuilder(r)
			reqExtra, err := requestExtraFactory(builder.ServiceContext())
			if err != nil {
				panic(err)
			}
			builder.SetRequestExtra(reqExtra)
			mifyCtx := builder.Build(r, w)

			ctx := context.WithValue(r.Context(), openapi_public.MifyContextField, mifyCtx)
			ctx = context.WithValue(ctx, MifyContextBuilderField, nil)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func getMifyRequestContextBuilder(r *http.Request) *core.MifyRequestContextBuilder {
	return r.Context().Value(MifyContextBuilderField).(*core.MifyRequestContextBuilder)
}
