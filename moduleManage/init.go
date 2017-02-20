package moduleManage

var DefaultManager *ApiModuleManagerStruct

func init() {
	DefaultManager = NewApiModuleManager()
}

func Init() {
	// 从global和player模块里面加载函数
}
