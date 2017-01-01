package common

import (
	"github.com/polariseye/polarserver/common/errorCode"
)

// 请求结果实体对象
type ResultModel struct {
	// 状态
	errorCode.ErrorInfo

	// 数据
	Value map[string]interface{}

	// 附加数据
	Avatar map[string]interface{}
}

// 创建新的实体对象
// defaultErrorCode:默认的错误码信息
func NewResultModel(defaultErrorCode errorCode.ErrorCode) (result *ResultModel) {
	result = &ResultModel{
		Value:  make(map[string]interface{}),
		Avatar: make(map[string]interface{}),
	}

	result.SetNormalError(defaultErrorCode)

	return
}
