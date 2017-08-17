package handler1

// RspTyp 内部使用的响应结构类型
type RspTyp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// RspTyp实现Error接口
func (r *RspTyp) Error() string {
	return r.Msg
}

var (
	BeanOK          = &RspTyp{Status: 0, Msg: "OK"}
	BeanAuthError   = &RspTyp{Status: 1, Msg: "Bad Credentials"}
	BeanEntityError = &RspTyp{Status: 2, Msg: "Unprocessable Entity"}
	BeanMethodError = &RspTyp{Status: 3, Msg: "Unsupport Method"}
	BeanPathError   = &RspTyp{Status: 4, Msg: "Parse Path Error"}
	BeanParamError  = &RspTyp{Status: 5, Msg: "Parse Parameter Error"}
	BeanBodyError   = &RspTyp{Status: 6, Msg: "Problems parsing JSON"}
	BeanInnerError  = &RspTyp{Status: 99999, Msg: "Inner Error"}
)
