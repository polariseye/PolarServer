package ModuleManage

// 动态数据接口
type IConvertor interface {
	// 模块名
	Name() string

	// 类型转换
	Convert() []string
}
