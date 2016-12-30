package ErrorCode

// 错误码数据
type ErrorCode struct {

	// 状态码
	code int32

	// 消息
	message string
}

// 创建新的错误码对象
// code:错误码
// message:提示消息
// 返回值：
// *ErrorCode:错误码对象
func NewErrorCode(_code int32, _message string) ErrorCode {
	return ErrorCode{
		code:    _code,
		message: _message,
	}
}

// 状态码
func (this *ErrorCode) Code() int32 {
	return this.code
}

// 错误信息
func (this *ErrorCode) Message() string {
	return this.message
}
