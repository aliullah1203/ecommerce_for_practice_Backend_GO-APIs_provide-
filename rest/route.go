package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /aliullah",
		manager.With(
			http.HandlerFunc(handlers.Test),
			middleware.Arekta,
		),
	)

	mux.Handle(
		"GET /ali",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)

	mux.Handle(
		"GET /products/{productId}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)
}
