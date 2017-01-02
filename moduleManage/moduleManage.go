package moduleManage

import (
	"sort"
)

// 模块管理结构
type moduleManagerStruct struct {
	// 所有模块对象
	moduleData map[Priority][]IModule

	// api模块集合
	apiModule []IModule
}

// 模块处理优先级
type Priority int

// 模块类型
type ModuleType int

const (
	// 最后处理
	Last Priority = -0xffff

	// 延后处理
	Low Priority = -1

	// 常规处理
	Normal Priority = 0

	// 较高优先处理
	Hight Priority = 1

	// 最先处理
	First Priority = 0xffff
)

const (
	// 常规模块
	NormalModule = 1

	// Api模块
	ApiModule = 2
)

// 模块管理对象
var moduleManager *moduleManagerStruct

// 初始化
func init() {
	moduleManager = newModuleManager()
}

// 创建新的模块管理对象
func newModuleManager() *moduleManagerStruct {
	return &moduleManagerStruct{
		moduleData: make(map[Priority][]IModule, 0),
		apiModule:  make([]IModule, 0, 10),
	}
}

// 注册模块对象
// module:模块对象
// priority:优先级
// moduleType:模块类型
func RegisterModule(module IModule, priority Priority, moduleType ModuleType) {

	// 添加到api模块中
	switch moduleType {
	case ApiModule:
		DefaulApiModuleManager.AddApiModule(module)
	}

	// 添加到总模块管理对象中
	moduleByPriority, isExist := moduleManager.moduleData[priority]
	if isExist == false {
		moduleByPriority = make([]IModule, 0, 100)
	}

	// 追加到按照优先级划分的所有模块综合那个
	moduleByPriority = append(moduleByPriority, module)

	// 保存到全局对象
	moduleManager.moduleData[priority] = moduleByPriority
}

// 初始化所有模块
// 返回值:
// []error:错误信息
func InitModule() []error {
	// 获取所有优先级
	keys := make([]int, 0, len(moduleManager.moduleData))
	for keyItem, _ := range moduleManager.moduleData {
		keys = append(keys, int(keyItem))
	}

	sort.Ints(keys)

	// 错误列表
	errorList := make([]error, 0, 100)
	for index := len(keys) - 1; index >= 0; index-- {
		moduleArray, _ := moduleManager.moduleData[Priority(keys[index])]

		// 按模块初始化
		tmpErrorList := initModuleByArray(moduleArray)
		if tmpErrorList != nil && len(tmpErrorList) > 0 {
			errorList = append(errorList, tmpErrorList...)
		}
	}

	return errorList
}

// 按照数组形式初始化模块
// moduleArray:模块列表
// 返回值:
// []error:错误信息
func initModuleByArray(moduleArray []IModule) []error {
	errorList := make([]error, 0, 100)

	for _, item := range moduleArray {
		// 初始化
		switch item.(type) {
		case IIniter:
			{
				initer := item.(IIniter)
				tmpErr := initer.InitModule()
				if tmpErr != nil && len(tmpErr) > 0 {
					errorList = append(errorList, tmpErr...)
				}
			}
		}

		// Check
		switch item.(type) {
		case IChecker:
			{
				checker := item.(IChecker)
				tmpErr := checker.CheckModule()
				if tmpErr != nil && len(tmpErr) > 0 {
					errorList = append(errorList, tmpErr...)
				}
			}
		}

		// 转换
		switch item.(type) {
		case IConvertor:
			{
				convertor := item.(IConvertor)
				tmpErr := convertor.ConvertModule()
				if tmpErr != nil && len(tmpErr) > 0 {
					errorList = append(errorList, tmpErr...)
				}
			}
		}
	}

	return errorList
}
