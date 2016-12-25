package ModuleManage

// 模块接口
type IModel interface {
	// 模块名
	Name() string

	// 初始化
	Init() error

	// 数据检查
	Check() []string

	// 类型转换
	Convert() []string
}
