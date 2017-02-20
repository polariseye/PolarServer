package config

import (
	"github.com/polariseye/polarserver/bll/instanceMangeBll/modelData"
)

type playerLvConfigStruct struct {
}

var PlayerLvConfigBll *playerLvConfigStruct

func init() {
	PlayerLvConfigBll = newPlayerLvConfigBll()
	modelData.Register(PlayerLvConfigBll)
}

func (this *playerLvConfigStruct) Init() []error {
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
