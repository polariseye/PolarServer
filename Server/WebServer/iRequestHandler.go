package WebServer

import (
	"net/http"
)

// 请求处理接口
type IRequestHandler interface {
	// 请求处理
	RequestHandle(response http.ResponseWriter, request *http.Request)
}
