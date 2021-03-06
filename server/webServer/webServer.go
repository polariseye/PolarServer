package webServer

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Jordanzuo/goutil/logUtil"
	"github.com/polariseye/polarserver/server/serverBase"
)

// web服务结构体
type WebServerStruct struct {
	serverBase.ServerBaseStruct

	// 服务实例
	serverInstance *http.Server

	// 运行端口
	port int32

	// 路由对象
	serveMux *http.ServeMux
}

// 开启服务
// onstopFun：服务停止时，运行的函数
// 返回值：
// error:错误信息
func (this *WebServerStruct) Start(onstopFun serverBase.OnStopFun) error {
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

	logUtil.LogAndPrint(fmt.Sprintf("服务名:%v 监听端口:%v", this.Name(), this.port), logUtil.Info)
	go func() {
		// 开启监听
		if tmpErr := serverInstance.ListenAndServe(); tmpErr != nil {
			this.innerStop(tmpErr)
			return
		}
	}()

	return nil
}

// 添加路由
// routerUrl:路由地址
// handler:处理函数
func (this *WebServerStruct) AddRouter(routerUrl string, handler func(http.ResponseWriter, *http.Request)) {
	this.serveMux.HandleFunc(routerUrl, handler)
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
	defer func() {
		if data := recover(); data != nil {
			logUtil.LogUnknownError(data)
		}
	}()
	/* result := fmt.Sprintf(`<html><body>
			<div>host:%v</div>
			<div>RequestURI:%v</div>
			<div>RemoteAddr:%v</div>
			<div>Referer:%v</div>
			</body></html>`,
		request.Host, request.RequestURI, request.RemoteAddr, request.Referer())
	response.Write(bytes.NewBufferString(result).Bytes())
	// ControllerManager.RequestHandle(response, request)
	*/
	this.serveMux.ServeHTTP(response, request)
}

// 创建新的web服务对象
// _port:端口
// serverName:服务名
// 返回值:
// *webServerStruct:新建的对象
func NewWebServer(_port int32, serverName string) *WebServerStruct {
	webServer := &WebServerStruct{
		serveMux: http.NewServeMux(),
		port:     _port,
	}

	webServer.InitBase("Web服务")

	return webServer
}
