package ModuleManage

// 初始化接口
type IIniter interface {
	// 模块名
	Name() string

	// 初始化
	Init() []error
}
