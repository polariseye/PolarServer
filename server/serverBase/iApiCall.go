package serverBase

import (
	"github.com/polariseye/polarserver/common"
)

// Api调用接口
type IApiCaller interface {
	// 调用实际处理函数
	Call(*common.RequestModel) *common.ResultModel
}
