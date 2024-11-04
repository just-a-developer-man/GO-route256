package controller_http

import (
	"net/http"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/controller/http/middleware"
)

// NewRouter - returns http.Handler
func (c *Controller) NewRouter() http.Handler {
	// Router layer
	mux := http.NewServeMux()

	// Note: You can add here custom middleware too
	mux.HandleFunc("/v1/order/create", middleware.CheckMethodPostMiddleware(c.CreateOrderHandler))

	return mux
}
