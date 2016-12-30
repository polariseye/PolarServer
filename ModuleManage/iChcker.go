package ModuleManage

// 模块接口
type IChecker interface {
	// 模块名
	Name() string

	// 数据检查
	Check() []string
}
