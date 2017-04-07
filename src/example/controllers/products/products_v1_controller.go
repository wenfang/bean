package products

import (
	"context"

	"nkwangwenfang.com/kit/transport/http/handler"
)

type productsRequest struct {
	Keys uint64 `req_path:"keys"`
	ID   string `req_path:"id"`
	Name string `req_param:"name"`
}

type productsResponse struct {
	Result string `json:"result"`
}

func productsV1Controller(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(*productsRequest)
	return &productsResponse{Result: request.ID}, nil
}

var ProductsV1Controller = handler.New(context.Background(), productsV1Controller, productsRequest{})
