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

func carsV1Controller(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*carsRequest)
	rsp := &carsResponse{
		ID:     req.ID,
		Errors: "error test",
	}
	return &rsp, handler.BeanInnerError
}

var CarsV1Controller = handler.New(carsV1Controller, carsRequest{})
