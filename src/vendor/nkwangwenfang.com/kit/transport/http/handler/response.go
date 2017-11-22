package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"nkwangwenfang.com/log"
)

// response 通用的响应结构
type response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

func Output(w http.ResponseWriter, data interface{}, err error) {
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
		rspTyp, ok := errors.Cause(err).(*RspTyp)
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
	if err := json.NewEncoder(w).Encode(response{Status: status, Msg: msg, Data: data}); err != nil {
		// 输出结果错误转换为内部错误,再输出一次,忽略错误
		log.Error("response encode error, change to inner error", "error", err)
		json.NewEncoder(w).Encode(response{Status: BeanInnerError.Status, Msg: BeanInnerError.Msg})
	}
}
