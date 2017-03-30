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
		return ErrorInner
	}
	return nil
}

// 输出错误,err决定了code和message的值,reason决定了errors的值
func OutputError(w http.ResponseWriter, reason interface{}, err error) {
	// 转换为Error类型
	e, ok := errors.Cause(err).(*Error)
	if !ok {
		log.Error("error type invalid, set to ErrorInner", "error", err)
		e = ErrorInner
	}
	// 编码输出失败结果
	if err := json.NewEncoder(w).Encode(response{Code: e.Code, Message: e.Message, Errors: reason}); err != nil {
		// 输出失败结果错误，去掉reason重新输出一遍
		log.Error("error reason invalid, set to nil", "error reason", reason)
		json.NewEncoder(w).Encode(response{Code: e.Code, Message: e.Message})
	}
}
