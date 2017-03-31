package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"nkwangwenfang.com/kit/transport/http/handler"

	"example/controller"
)

// Init 初始化router
func Init() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	r.Handle("/products/{keys}/{id:[0-9]+}", controller.ProductsV1Controller).Methods("GET")
	http.Handle("/", r)
}
