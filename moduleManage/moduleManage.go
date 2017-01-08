package moduleManage

import (
	"fmt"
	"sort"
)

// 模块管理结构
type ModuleManagerStruct struct {
	// 所有模块对象
	moduleData map[ModuleType][]*ModuleInfoStruct

	// api管理对象集合
	apiManagerData []*ApiModuleManagerStruct
}

// 模块类型
type ModuleType int

const (
	// 最后处理的模块
	LastModule ModuleType = -0xffff

	// 延后处理的模块
	LowModule ModuleType = -1

	// 常规处理模块
	NormalModule ModuleType = 0

	// 较高优先处理的模块
	HightModule ModuleType = 1

	// 最先处理的模块
	FirstModule ModuleType = 0xffff
)

// 模块管理对象
var moduleManager *ModuleManagerStruct

// 初始化
func init() {
	moduleManager = newModuleManager()
}

// 创建新的模块管理对象
func newModuleManager() *ModuleManagerStruct {
	return &ModuleManagerStruct{
		moduleData:     make(map[ModuleType][]*ModuleInfoStruct, 0),
		apiManagerData: make([]*ApiModuleManagerStruct, 0),
	}
}

// 注册模块创建函数
// module:模块对象
// moduleType:模块类型
// error:错误信息
func RegisterModule(newInstanceFun func() (IModule, ModuleType)) {

	moduleInfo := newModuleInfo(newInstanceFun)

	// 添加到总模块管理对象中
	moduleByType, isExist := moduleManager.moduleData[moduleInfo.Type()]
	if isExist == false {
		moduleByType = make([]*ModuleInfoStruct, 0, 100)
	}

	// 追加到按照优先级划分的所有模块综合那个
	moduleByType = append(moduleByType, moduleInfo)

	// 保存到全局对象
	moduleManager.moduleData[moduleInfo.Type()] = moduleByType
}

// 初始化所有模块
// 返回值:
// []error:错误信息
func InitModule() []error {
	// 错误列表
	errorList := make([]error, 0, 100)

	// api模块初始化
	if len(moduleManager.apiManagerData) <= 0 {
		errorList = append(errorList, fmt.Errorf("未注册任何Api模块管理"))
	}
	for _, itemByType := range moduleManager.moduleData {
		for _, item := range itemByType {
			for _, tmpApiManager := range moduleManager.apiManagerData {
				tmpApiManager.AddApiModule(item.Module())
			}
		}
	}

	// 获取所有优先级
	keys := make([]int, 0, len(moduleManager.moduleData))
	for keyItem, _ := range moduleManager.moduleData {
		keys = append(keys, int(keyItem))
	}

	sort.Ints(keys)

	// 模块初始化
	for index := len(keys) - 1; index >= 0; index-- {
		moduleArray := moduleManager.moduleData[ModuleType(keys[index])]

		moduleList := make([]IModule, 0)
		for _, item := range moduleArray {
			moduleList = append(moduleList, item.Module())
		}

		// 数据校验
		errorList = append(errorList, initModuleByArray(moduleList)...)
	}

	return errorList
}

// 按照模块类型重新加载模块，如果加载过程中存在错误，则会终止加载
// moduleType:模块类型
// isReloadApi:是否重新加载对应的api接口
// 返回值:
// []error:错误信息
func Reload(moduleType ModuleType, isReloadApi bool) []error {
	errorList := make([]error, 0, 100)

	// 获取指定模块类型下的所有模块
	moduleByType, isExist := moduleManager.moduleData[moduleType]
	if isExist == false {
		return errorList
	}

	// 初始化后的所有模块对象
	newModuleInfoData := make(map[*ModuleInfoStruct]IModule, 0)
	newModuleList := make([]IModule, 0)

	// 初始化所有模块
	for _, moduleInfoItem := range moduleByType {
		newModule := moduleInfoItem.GetNewModule()
		newModuleInfoData[moduleInfoItem] = newModule
		newModuleList = append(newModuleList, newModule)
	}

	errorList = append(errorList, initModuleByArray(newModuleList)...)

	// 如果无错误，则重新加载模块
	if len(errorList) <= 0 {
		//todo:此处到底用不用锁尼
		for moduleInfoItem, newModule := range newModuleInfoData {
			moduleInfoItem.ResetModule(newModule)

			// 重新加载api
			if isReloadApi {
				for _, apiManager := range moduleManager.apiManagerData {
					apiManager.AddApiModule(newModule)
				}
			}
		}
	}

	return errorList
}

// check和convert测试（对现有数据不会产生影响）
// 返回值:
// []error:错误信息
func TestCheckAndConvert(moduleType ModuleType) []error {
	errorList := make([]error, 0, 100)

	// 获取指定模块类型下的所有模块
	moduleByType, isExist := moduleManager.moduleData[moduleType]
	if isExist == false {
		return errorList
	}

	// 初始化所有模块
	newModuleList := make([]IModule, 0)
	for _, moduleInfoItem := range moduleByType {
		newModule := moduleInfoItem.GetNewModule()
		newModuleList = append(newModuleList, newModule)
	}

	errorList = append(errorList, initModuleByArray(newModuleList)...)

	return errorList
}

// 按照数组形式初始化模块
// moduleArray:模块列表
// 返回值:
// []error:错误信息
func initModuleByArray(moduleArray []IModule) []error {
	errorList := make([]error, 0, 100)

	for _, item := range moduleArray {
		// init
		initer, isIniter := item.(IIniter)
		if isIniter {
			tmpErr := initer.InitModule()
			if tmpErr != nil && len(tmpErr) > 0 {
				errorList = append(errorList, tmpErr...)
			}

		}
	}

	for _, item := range moduleArray {
		// check
		checker, isChecker := item.(IChecker)
		if isChecker {
			tmpErr := checker.CheckModule()
			if tmpErr != nil && len(tmpErr) > 0 {
				errorList = append(errorList, tmpErr...)
			}
		}

		// convert
		convertor, isConvertor := item.(IConvertor)
		if isConvertor {
			tmpErr := convertor.ConvertModule()
			if tmpErr != nil && len(tmpErr) > 0 {
				errorList = append(errorList, tmpErr...)
			}
		}
	}

	return errorList
}

// 添加Api管理对象
func AddApiManager(apiManager *ApiModuleManagerStruct) {
	moduleManager.apiManagerData = append(moduleManager.apiManagerData, apiManager)
}
