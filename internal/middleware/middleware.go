package middleware

import "net/http"

func ChainMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for i := 0; i < len(middleware); i++ {
		h = middleware[i](h)
	}
	return h
}
