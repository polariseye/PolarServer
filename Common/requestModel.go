package Common

import (
	"net/http"

	"github.com/Jordanzuo/goutil/typeUtil"
)

// 请求实体对象
type RequestModel struct {

	// 模块名
	ModuleName string

	// 功能名
	MethodName string

	// 请求者地址
	Ip string

	// 请求内容
	Data []interface{}

	// 扩展请求数据
	ExtensionString typeUtil.MapData

	// 请求的对象
	Request *http.Request
}

// 创建新的请求对象
func NewRequestModel() *RequestModel {
	return &RequestModel{
		Data:            make([]interface{}, 0),
		ExtensionString: typeUtil.NewMapData(),
	}
}
