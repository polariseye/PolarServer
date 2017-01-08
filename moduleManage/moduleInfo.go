package moduleManage

// 模块数据结构
type ModuleInfoStruct struct {
	// 模块对象
	module IModule

	// 模块类型
	moduleType ModuleType

	// 新实例创建函数
	newInstanceFun func() (IModule, ModuleType)
}

// 获取当前模块对象
func (this *ModuleInfoStruct) Module() IModule {
	return this.module
}

// 重置模块对象
func (this *ModuleInfoStruct) ResetModule(module IModule) {
	this.module = module
}

// 获取模块类型
func (this *ModuleInfoStruct) Type() ModuleType {
	return this.moduleType
}

// 获取新的模块对象
func (this *ModuleInfoStruct) GetNewModule() IModule {
	tmpModule, _ := this.newInstanceFun()
	return tmpModule
}

// 创建新的模块信息对象
// _newInstanceFun:具体模块创建函数
// 返回值:
// *moduleInfoStruct:模块创建对象信息
func newModuleInfo(_newInstanceFun func() (IModule, ModuleType)) *ModuleInfoStruct {
	_module, _moduleType := _newInstanceFun()
	return &ModuleInfoStruct{
		newInstanceFun: _newInstanceFun,
		module:         _module,
		moduleType:     _moduleType,
	}
}
