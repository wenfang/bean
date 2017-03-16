package main

import (
	"nkwangwenfang.com/kit/transport/http/handler"
)

var (
	errNoID        = handler.NewError(20000, "ID not set")
	errID          = handler.NewError(20010, "ID error")
	errOutOrderID  = handler.NewError(20011, "out_order_id error")
	errUseCarType  = handler.NewError(20012, "use_car_type error")
	errStatusNo    = handler.NewError(20013, "status_no error")
	errCompanyID   = handler.NewError(20014, "company_id error")
	errMemberID    = handler.NewError(20015, "member_id error")
	errPassengerID = handler.NewError(20016, "passenger_id error")
	errOrderSource = handler.NewError(20017, "order_souce error")
	errParamType   = handler.NewError(20020, "param type error")
)
