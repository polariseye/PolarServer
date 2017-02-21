package globalData

import (
	"fmt"

	"github.com/Jordanzuo/goutil/stringUtil"
)

var moduleData map[string]IGlobal

func Register(globalItem IGlobal) {
	moduleName := globalItem.ModuleName()
	if stringUtil.IsEmpty(moduleName) {
		panic(fmt.Errorf("global模块的模块名不能为空"))
	}

	_, isExist := moduleData[moduleName]
	if isExist {
		panic(fmt.Errorf("global模块的模块名重复,ModuleName:%v", moduleName))
	}

	moduleData[moduleName] = globalItem
}

func Modules() map[string]IGlobal {
	return moduleData
}
