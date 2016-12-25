package ModuleManage

// 全局数据接口
type IGlobal interface {
	// 模块名
	Name() string

	// 初始化
	Init() error
}
