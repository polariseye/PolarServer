package errorCode

import (
	"fmt"

	"github.com/Jordanzuo/goutil/stringUtil"
)

// 错误信息
type ErrorInfo struct {
	// 错误码
	Status int32

	// 错误信息
	Message string
}

// 返回错误信息
func (this *ErrorInfo) Error() string {
	return fmt.Sprintf("ErrorCode:%v ErrorMessage:%v", this.Status, this.Message)
}

// 设置错误信息
func (this *ErrorInfo) SetNormalError(errorCode ErrorCode) {
	this.Status = errorCode.Code()
	this.Message = errorCode.Message()
}

// 设置错误信息
func (this *ErrorInfo) SetError(errorCode ErrorCode, message string) {
	this.Status = errorCode.Code()
	if stringUtil.IsEmpty(message) {
		this.Message = errorCode.Message()
	} else {
		this.Message = message
	}
}

// 拷贝错误信息
func (this *ErrorInfo) CopyErrorInfo(errorInfo *ErrorInfo) {
	this.Status = errorInfo.Status
	this.Message = errorInfo.Message
}

// 是否成功
func (this *ErrorInfo) IsSuccess() bool {
	return this.Status == Success.Code()
}

// 创建错误对象
func NewError(code ErrorCode, msg string) *ErrorInfo {
	tmpError := &ErrorInfo{
		Status:  code.Code(),
		Message: msg,
	}

	if stringUtil.IsEmpty(tmpError.Message) {
		tmpError.Message = code.Message()
	}

	return tmpError
}
