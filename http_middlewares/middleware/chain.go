package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

type Chain struct {
	middlewares []Middleware
	final       http.Handler
}

func New(middlewares []Middleware, final http.Handler) *Chain {
	return &Chain{
		middlewares: middlewares,
		final:       final,
	}
}

func (c *Chain) Build() http.Handler {
	handler := c.final

	for i := len(c.middlewares) - 1; i >= 0; i-- {
		handler = c.middlewares[i](handler)
	}

	return handler
}
