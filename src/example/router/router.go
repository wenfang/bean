package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"nkwangwenfang.com/kit/transport/http/handler"

	"example/controllers/cars"
	"example/controllers/products"
)

// Init 初始化router
func Init() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	r.Handle("/products/{keys}/{id:[0-9]+}", products.ProductsV1Controller).Methods("GET")
	r.Handle("/cars", cars.CarsV1Controller).Methods("GET")

	http.Handle("/", r)
}
