package cars

import (
	"context"

	"nkwangwenfang.com/kit/transport/http/handler1"
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
	return "error test", handler1.BeanInnerError
}

var CarsV1Controller = handler1.New(carsV1Controller, carsRequest{})
