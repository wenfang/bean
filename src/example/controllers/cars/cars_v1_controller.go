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
	ID     uint64 `json:"id"`
	Errors string `json:"errors"`
}

func carsV1Controller(ctx context.Context, req interface{}) (interface{}, error) {
	panic("panic test")
	return "error test", handler.ErrorTypeInner
}

var CarsV1Controller = handler.New(context.Background(), carsV1Controller, carsRequest{})
