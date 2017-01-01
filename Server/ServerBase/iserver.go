package serverBase

// 服务停止事件
type OnStopFun func(serverInstance IServer)

// 服务接口
type IServer interface {

	// 启动服务
	Start(OnStopFun) error

	// 停止服务
	Stop() error

	// 是否正在运行
	IsRun() bool

	// 服务名
	Name() string
}
