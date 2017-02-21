package playerData

import (
	"fmt"

	"github.com/Jordanzuo/goutil/stringUtil"
)

var moduleData map[string]IPlayerData

func Register(playerDataItem IPlayerData) {
	moduleName := playerDataItem.ModuleName()
	if stringUtil.IsEmpty(moduleName) {
		panic(fmt.Errorf("player模块的模块名不能为空"))
	}

	_, isExist := moduleData[moduleName]
	if isExist {
		panic(fmt.Errorf("player模块的模块名重复,ModuleName:%v", moduleName))
	}

	moduleData[moduleName] = playerDataItem
}

func Modules() map[string]IPlayerData {
	return nil
}
