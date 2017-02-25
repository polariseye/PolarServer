package configDal

import (
	"github.com/Jordanzuo/goutil/dbUtil"
	"github.com/polariseye/polarserver/dal/dataBase"
)

type playerLvConfigDal_ struct {
}

var PlayerLvConfigDal *playerLvConfigDal_ = &playerLvConfigDal_{}

func (this *playerLvConfigDal_) GetList() (*dbUtil.DataTable, error) {
	sql := "select * from b_player_lv_c"
	result, errMsg := dataBase.GameDb().Query(sql)
	if errMsg != nil {
		return nil, errMsg
	}

	return dbUtil.NewDataTable(result)
}
