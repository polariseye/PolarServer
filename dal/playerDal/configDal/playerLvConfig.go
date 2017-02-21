package configDal

import (
	"github.com/polariseye/polarserver/dal/dataBase"
)

type playerLvConfigDal_ struct {
}

var PlayerLvConfigDal *playerLvConfigDal_ = &playerLvConfigDal_{}

func (this *playerLvConfigDal_) GetList() {
	sql := "select * from b_player_lv_c"
	result, errMsg := dataBase.GameDb().Query(sql)
}

func (this *playerLvConfigDal_) Update() {

}

func (this *playerLvConfigDal_) Insert() {

}

func (this *playerLvConfigDal_) Delete() {

}
