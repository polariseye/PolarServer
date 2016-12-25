package WebServer

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/polariseye/PolarServer/Server/ServerBase"
)

// web服务结构体
type WebServerStruct struct {
	ServerBase.ServerBaseStruct

	// 服务实例
	serverInstance *http.Server

	// 运行端口
	port int32

	// 处理对象
	handler IRequestHandler
}

// 开启服务
// onstopFun：服务停止时，运行的函数
// 返回值：
// error:错误信息
func (this *WebServerStruct) Start(onstopFun ServerBase.OnStopFun) error {
	this.Locker().Lock()
	defer this.Locker().Unlock()

	// 配置初始化
	this.SetOnStopFun(onstopFun)

	// 开启服务
	serverInstance := http.Server{
		Addr:    fmt.Sprintf(":%v", this.port),
		Handler: this,
	}

	this.serverInstance = &serverInstance

	go func() {
		// 开启监听
		if tmpErr := serverInstance.ListenAndServe(); tmpErr != nil {
			this.innerStop(tmpErr)
			return
		}
	}()

	return nil
}

// 停止服务
// 返回值：
// error:服务停止的错误信息
func (this *WebServerStruct) Stop() error {
	return errors.New("不能停止服务")
}

// 内部的停止服务逻辑
// 返回值：
// error:服务停止的错误信息
func (this *WebServerStruct) innerStop(errMsg error) {
	this.Locker().Lock()
	defer this.Locker().Unlock()

	if this.IsRun() == false {
		return
	}

	this.InvokeOnStop(this)

	if errMsg != nil {
		log.Printf("服务：%v 已停止，停止信息：%v", this.Name(), errMsg.Error())
	} else {
		log.Printf("服务：%v 已停止", this.Name())
	}

	this.SetIsRun(false)
}

// http应答处理
// response:应答对象
// request:请求对象
func (this *WebServerStruct) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	this.handler.RequestHandle(response, request)
}

// 创建新的web服务对象
// port:端口
// _handler:处理函数
// 返回值:
// *webServerStruct:新建的对象
func NewWebServer(port int32, _handler IRequestHandler) *WebServerStruct {
	webServer := &WebServerStruct{}

	webServer.InitBase("Web服务")

	webServer.port = port
	webServer.handler = _handler

	return webServer
}
