package ModuleManage

import (
	"github.com/polariseye/PolarServer/Common"
	"github.com/polariseye/PolarServer/Common/ErrorCode"
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
	ApiModile = 2
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

func Invoke(request *Common.RequestModel) (result *Common.ResultModel, errMsg error) {
	//todo:待完善
	result = Common.NewResultModel(ErrorCode.DataError("还未实现"))

	return result
}

// 注册模块对象
// module:模块对象
// priority:优先级
// moduleType:模块类型
func RegisterModule(module IModule, priority Priority, moduleType ModuleType) {

	// 添加到api模块中
	switch moduleType {
	case ApiModile:
		moduleManager.apiModule = append(moduleManager.apiModule, module)
		break
	}

	// 添加到总模块管理对象中
	moduleByPriority, isExist := moduleManager.moduleData[priority]
	if isExist == false {
		moduleByPriority = make([]IModule, 0, 100)
	}

	// 追加到按照优先级划分的所有模块综合那个
	moduleByPriority = append(moduleByPriority, module)

	moduleManager.moduleData[priority] = moduleByPriority
}
