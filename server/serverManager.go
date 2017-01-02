package server

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Jordanzuo/goutil/logUtil"
	"github.com/polariseye/polarserver/server/serverBase"
)

// 服务管理对象
type ServerManagerStruct struct {

	// 服务列表
	serverData map[string]serverBase.IServer

	// 同步对象
	dataLocker sync.Locker

	// 是否已经开启运行
	isStart bool

	// 等待所有服务停止
	waitGroup sync.WaitGroup
}

// 注册服务
// server:需要注册的服务
func (this *ServerManagerStruct) Register(server serverBase.IServer) {
	this.dataLocker.Lock()
	defer this.dataLocker.Unlock()

	this.serverData[server.Name()] = server
}

// 开始运行服务
// error:服务运行的错误信息
func (this *ServerManagerStruct) Start() error {
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
		logUtil.LogAndPrint(fmt.Sprintf("开始初始化服务, 服务名:%v", item.Name()), logUtil.Info)
		if errMsg := item.Start(this.onServerStop); errMsg != nil {
			return errMsg
		}

		logUtil.LogAndPrint(fmt.Sprintf("服务初始化完成, 服务名:%v", item.Name()), logUtil.Info)
		this.waitGroup.Add(1)
	}

	this.isStart = true

	return nil
}

// 服务停止时，触发的动作
// serverInstance：已停止的服务
func (this *ServerManagerStruct) onServerStop(serverInstance serverBase.IServer) {
	this.dataLocker.Lock()
	defer this.dataLocker.Unlock()

	_, isExist := this.serverData[serverInstance.Name()]
	if isExist == false {
		return
	}

	this.waitGroup.Done()
}

// 停止所有服务
func (this *ServerManagerStruct) Stop() error {
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
func (this *ServerManagerStruct) WaitStop() {
	this.waitGroup.Wait()
}

// 创建新的服务管理对象
func NewServerManager() (serverManager *ServerManagerStruct) {
	serverManager = &ServerManagerStruct{}
	serverManager.serverData = make(map[string]serverBase.IServer, 10)
	serverManager.dataLocker = &sync.Mutex{}

	return
}
