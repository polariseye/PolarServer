package configBll

import (
	"github.com/polariseye/polarserver/bll/instanceMangeBll/modelData"
	"github.com/polariseye/polarserver/dal/playerDal/configDal"
)

type playerLvConfigStruct struct {
}

var PlayerLvConfigBll *playerLvConfigStruct

func init() {
	PlayerLvConfigBll = newPlayerLvConfigBll()
	modelData.Register(PlayerLvConfigBll)
}

func (this *playerLvConfigStruct) Init() []error {
	table, _ := configDal.PlayerLvConfigDal.GetList()
	for i := 0; i < table.Len(); i++ {
		// table.Row(i)
	}

	return nil
}

func (this *playerLvConfigStruct) Check() []error {
	return nil
}

func (this *playerLvConfigStruct) Convert() []error {
	return nil
}

func (this *playerLvConfigStruct) ModuleName() string {
	return "PlayerLvConfigBll"
}

func newPlayerLvConfigBll() *playerLvConfigStruct {
	return new(playerLvConfigStruct)
}
