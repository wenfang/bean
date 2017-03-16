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
	// ErrAuth 授权错误
	ErrAuth = NewError(1, "Bad credentials")
	// ErrEntity 不存在对应的实体
	ErrEntity = NewError(2, "Unprocessable Entity")
	// ErrMethod 不支持的方法
	ErrMethod = NewError(3, "Unsupport Method")
	// ErrParameter 请求参数错误
	ErrParameter = NewError(4, "Parameter Error")
	// ErrBody 请求体的JSON错误
	ErrBody = NewError(5, "Problems parsing JSON")
	// ErrInner 内部错误
	ErrInner = NewError(999, "Internal Error")
)

type ErrorReason struct {
	Reason string `json:"reason,omitempty"`
}
