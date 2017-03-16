package middleware

import (
	"net/http"
)

// Middleware HTTP handler的中间件
type Middleware func(http.Handler) http.Handler

// Chain 串接HTTP handler中间件
func Chain(outer Middleware, others ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}
