package Server

import (
	"github.com/polariseye/PolarServer/Server/WebServer"
)

var (
	WebServerObj *WebServer.WebServerStruct
)

// 初始化
func init() {
	InitWebServer()
}

// 初始化web服务
func InitWebServer() {
	// web服务初始化
	handle4UrlItem := WebServer.NewHandle4Url()
	WebServerObj = WebServer.NewWebServer(0)
	WebServerObj.AddRouter("/Api", handle4UrlItem.RequestHandle)

	// 注册模块
	DefaultManager().Register(WebServerObj)
}
