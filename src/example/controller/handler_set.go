package controller

/*
import (
	"context"
	"net/http"
	"reflect"

	//	"nkwangwenfang.com/kit/expvar"
	"nkwangwenfang.com/kit/transport/http/handler"
	"nkwangwenfang.com/kit/transport/http/middleware/sign"
)

type setStatusRequest struct {
	ID           int64   `httpreq:"id,required"`
	OutOrderID   *int64  `httpreq:"out_order_id"`
	UseCarType   *int64  `httpreq:"use_car_type"`
	District     *string `httpreq:"district"`
	CompanyID    *int64  `httpreq:"company_id"`
	MemberID     *int64  `httpreq:"member_id"`
	PassengerID  *int64  `httpreq:"passenger_id"`
	OrderSource  *int    `httpreq:"order_source"`
	OrderStatus  *int    `httpreq:"order_status"`
	PayStatus    *int    `httpreq:"pay_status"`
	RefundStatus *int    `httpreq:"refund_status"`
	FreezeStatus *int    `httpreq:"freeze_status"`
}

var setStatusCounter = expvar.NewCounter("setStatus")

func decodeSetStatusRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	setStatusCounter.Add(1)

	var req setStatusRequest
	if err := handler.RequestDecode(request, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func setStatus(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*setStatusRequest)
	// 从请求中获取要更新的域和参数
	var (
		fields []string
		args   []interface{}
	)

	if req.OutOrderID != nil {
		fields = append(fields, "out_order_id")
		args = append(args, *req.OutOrderID)
	}
	if req.UseCarType != nil {
		fields = append(fields, "use_car_type")
		args = append(args, *req.UseCarType)
	}
	if req.District != nil {
		fields = append(fields, "district")
		args = append(args, *req.District)
	}
	if req.CompanyID != nil {
		fields = append(fields, "company_id")
		args = append(args, *req.CompanyID)
	}
	if req.MemberID != nil {
		fields = append(fields, "member_id")
		args = append(args, *req.MemberID)
	}
	if req.PassengerID != nil {
		fields = append(fields, "passenger_id")
		args = append(args, *req.PassengerID)
	}
	if req.OrderSource != nil {
		fields = append(fields, "order_source")
		args = append(args, *req.OrderSource)
	}
	if req.OrderStatus != nil {
		fields = append(fields, "order_status")
		args = append(args, *req.OrderStatus)
	}
	if req.PayStatus != nil {
		fields = append(fields, "pay_status")
		args = append(args, *req.PayStatus)
	}
	if req.RefundStatus != nil {
		fields = append(fields, "refund_status")
		args = append(args, *req.RefundStatus)
	}
	if req.FreezeStatus != nil {
		fields = append(fields, "freeze_status")
		args = append(args, *req.FreezeStatus)
	}
	// 将内容更新到数据库
	if err := GlobalDB.Set(req.ID, fields, args); err != nil {
		return nil, err
	}
	return nil, nil
}

func init() {
	http.HandleFunc("/", handler.NotFound)
	http.Handle("/setStatus", sign.Sign(handler.New(context.Background(), setStatus, reflect.TypeOf(setStatusRequest{}))))
}
*/
