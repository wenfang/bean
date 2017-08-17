package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"nkwangwenfang.com/log"
)

// response 通用的响应结构
type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// OutputResponse 输出结果数据
func OutputResponse(w http.ResponseWriter, data interface{}) error {
	// 输出响应结果
	if err := json.NewEncoder(w).Encode(response{Code: 0, Message: "Success", Data: data}); err != nil {
		log.Error("response data encode error", "error", err)
		return ErrorTypeInner
	}
	return nil
}

// OutputError 输出错误,err决定了code和message的值,reason决定了errors的值
func OutputError(w http.ResponseWriter, errData interface{}, errType error) {
	// 转换为Error类型
	e, ok := errors.Cause(errType).(*ErrorType)
	if !ok {
		log.Error("error type invalid, set to ErrorTypeInner", "error", errType)
		e = ErrorTypeInner
	}
	// 编码输出失败结果
	if err := json.NewEncoder(w).Encode(response{Code: e.Code, Message: e.Message, Errors: errData}); err != nil {
		// 输出失败结果错误，去掉reason重新输出一遍
		log.Error("error reason invalid, set to nil", "error reason", errData)
		json.NewEncoder(w).Encode(response{Code: e.Code, Message: e.Message})
	}
}

// response 通用的响应结构
type response1 struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

func Output1(w http.ResponseWriter, err error, data interface{}) {
	var (
		status int
		msg    string
	)
	// 如果没有错误设置正确返回
	if err == nil {
		status = BeanOK.Status
		msg = BeanOK.Msg
	} else {
		// 转换错误类型
		rspTyp, ok := err.(*RspTyp1)
		if ok {
			// 转换成功
			status = rspTyp.Status
			msg = rspTyp.Msg
		} else {
			// 转换不成功为内部错误
			status = BeanInnerError.Status
			msg = BeanInnerError.Msg
		}
	}
	// 编码输出结果失败
	if err := json.NewEncoder(w).Encode(response1{Status: status, Msg: msg, Data: data}); err != nil {
		// 输出结果错误转换为内部错误,再输出一次,忽略错误
		log.Error("response encode error, change to inner error", "error", err)
		json.NewEncoder(w).Encode(response1{Status: BeanInnerError.Status, Msg: BeanInnerError.Msg})
	}
}
