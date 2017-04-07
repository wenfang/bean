package cars

import (
	"context"

	"nkwangwenfang.com/kit/transport/http/handler"
)

type carsRequest struct {
	ID       uint64  `req_param:"id,required"`
	District *string `req_param:"district"`
}

type carsResponse struct {
	ID uint64 `json:"id"`
}

func carsV1Controller(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(*carsRequest)
	return &carsResponse{ID: request.ID}, nil
}

var CarsV1Controller = handler.New(context.Background(), carsV1Controller, carsRequest{})
