package ServerBase

import (
	"sync"
)

// 服务基础信息
type ServerBaseStruct struct {

	// 服务名
	name string

	//是否正在运行
	isRun bool

	// 停止时，触发的函数
	onStop OnStopFun

	// 服务同步锁对象
	serverLocker sync.Locker
}

// 初始化基类信息
// _name:服务名
func (this *ServerBaseStruct) InitBase(_name string) {
	this.name = _name
	this.isRun = false
	this.serverLocker = new(sync.Mutex)
}

// 服务名
// 返回值:
// string:服务名
func (this *ServerBaseStruct) Name() string {
	return this.name
}

// 是否正在运行
// 返回值:
// bool:是否正在运行
func (this *ServerBaseStruct) IsRun() bool {
	return this.isRun
}

// 设置是否在运行
// _isRun:要设置的值
func (this *ServerBaseStruct) SetIsRun(_isRun bool) {
	this.isRun = _isRun
}

// 触发停止函数
// server:服务对象
func (this *ServerBaseStruct) InvokeOnStop(server IServer) {
	if this.onStop != nil {
		this.onStop(server)
	}
}

// 设置结束触发的函数
func (this *ServerBaseStruct) SetOnStopFun(_onStop OnStopFun) {
	this.onStop = _onStop
}

// 获取锁对象
// 返回值:
// sync.Locker:锁对象
func (this *ServerBaseStruct) Locker() sync.Locker {
	return this.serverLocker
}
