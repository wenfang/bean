package handler

// ErrorType 错误类型，所有处理handler的endpoint返回的第二个参数应为一个ErrorType
type ErrorType struct {
	Code    int
	Message string
}

// NewErrorType 创建ErrorType
func NewErrorType(code int, message string) *ErrorType {
	return &ErrorType{Code: code, Message: message}
}

func (e *ErrorType) Error() string {
	return e.Message
}

var (
	// ErrorTypeAuth 授权错误
	ErrorTypeAuth = NewErrorType(1, "Bad credentials")
	// ErrorTypeEntity 不存在对应的实体
	ErrorTypeEntity = NewErrorType(2, "Unprocessable Entity")
	// ErrorTypeMethod 不支持的方法
	ErrorTypeMethod = NewErrorType(3, "Unsupport Method")
	// ErrorTypePath 解析路径参数错误
	ErrorTypePath = NewErrorType(4, "Parse Path Error")
	// ErrorTypeParameter 请求参数错误
	ErrorTypeParameter = NewErrorType(5, "Parameter Error")
	// ErrorTypeBody 请求体的JSON错误
	ErrorTypeBody = NewErrorType(6, "Problems parsing JSON")
	// ErrorTypeInner 内部错误
	ErrorTypeInner = NewErrorType(999, "Internal Error")
)
