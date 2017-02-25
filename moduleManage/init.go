package moduleManage

import (
	"github.com/polariseye/polarserver/bll/instanceMangeBll/globalData"
	"github.com/polariseye/polarserver/bll/instanceMangeBll/playerData"
)

var DefaultManager *ApiModuleManagerStruct

func init() {
	DefaultManager = NewApiModuleManager()
}

func Init() {
	// 从global和player模块里面加载函数
	for _, moduleItem := range globalData.Modules() {
		DefaultManager.AddApiModule(moduleItem)
	}

	for _, moduleItem := range playerData.Modules() {
		DefaultManager.AddApiModule(moduleItem)
	}
}
