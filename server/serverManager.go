package server

import (
	"errors"
	"sync"

	"github.com/polariseye/polarserver/server/serverBase"
)

// 服务管理对象
type serverManagerStruct struct {

	// 服务列表
	serverData map[string]serverBase.IServer

	// 同步对象
	dataLocker sync.Locker

	// 是否已经开启运行
	isStart bool

	// 等待所有服务停止
	waitGroup sync.WaitGroup
}

// 服务管理对象
var defaultServerManager *serverManagerStruct

// 初始化
func init() {
	defaultServerManager = NewServerManager()
}

// 注册服务
// server:需要注册的服务
func (this *serverManagerStruct) Register(server serverBase.IServer) {
	this.dataLocker.Lock()
	defer this.dataLocker.Unlock()

	this.serverData[server.Name()] = server
}

// 开始运行服务
// error:服务运行的错误信息
func (this *serverManagerStruct) Start() error {
	this.dataLocker.Lock()
	defer this.dataLocker.Unlock()

	if this.isStart {
		return errors.New("服务已经开启")
	}

	if len(this.serverData) <= 0 {
		errors.New("没有注册任何服务")
	}

	for _, item := range this.serverData {

		// 服务开启异常
		if errMsg := item.Start(this.onServerStop); errMsg != nil {
			return errMsg
		}

		this.waitGroup.Add(1)
	}

	this.isStart = true

	return nil
}

// 服务停止时，触发的动作
// serverInstance：已停止的服务
func (this *serverManagerStruct) onServerStop(serverInstance serverBase.IServer) {
	this.dataLocker.Lock()
	defer this.dataLocker.Unlock()

	_, isExist := this.serverData[serverInstance.Name()]
	if isExist == false {
		return
	}

	this.waitGroup.Done()
}

// 停止所有服务
func (this *serverManagerStruct) Stop() error {
	this.dataLocker.Lock()
	defer this.dataLocker.Unlock()

	if this.isStart == false {
		return errors.New("服务未开启")
	}

	for _, item := range this.serverData {

		// 服务停止异常
		if errMsg := item.Stop(); errMsg != nil {
			return errMsg
		}
	}

	this.isStart = false

	return nil
}

// 等待所有服务停止
func (this *serverManagerStruct) WaitStop() {
	this.waitGroup.Wait()
}

// 创建新的服务管理对象
func NewServerManager() (serverManager *serverManagerStruct) {
	serverManager = &serverManagerStruct{}
	serverManager.serverData = make(map[string]serverBase.IServer, 10)
	serverManager.dataLocker = &sync.Mutex{}

	return
}

// 返回默认的管理对象
func DefaultManager() *serverManagerStruct {
	return defaultServerManager
}
