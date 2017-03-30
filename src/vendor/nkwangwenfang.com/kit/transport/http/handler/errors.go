package handler

type Error struct {
	Code    int
	Message string
}

// NewError 创建handler error
func NewError(code int, message string) *Error {
	return &Error{Code: code, Message: message}
}

func (e *Error) Error() string {
	return e.Message
}

var (
	// ErrorAuth 授权错误
	ErrorAuth = NewError(1, "Bad credentials")
	// ErrorEntity 不存在对应的实体
	ErrorEntity = NewError(2, "Unprocessable Entity")
	// ErrorMethod 不支持的方法
	ErrorMethod = NewError(3, "Unsupport Method")
	// ErrorPath 解析路径参数错误
	ErrorPath = NewError(4, "Parse Path Error")
	// ErrorParameter 请求参数错误
	ErrorParameter = NewError(5, "Parameter Error")
	// ErrorBody 请求体的JSON错误
	ErrorBody = NewError(6, "Problems parsing JSON")
	// ErrorInner 内部错误
	ErrorInner = NewError(999, "Internal Error")
)

// ErrorReason 错误结构
type ErrorReason struct {
	Reason string `json:"reason,omitempty"`
}
