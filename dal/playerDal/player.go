package playerDal

import (
	"github.com/Jordanzuo/goutil/dbUtil"
	"github.com/polariseye/polarserver/dal/dataBase"
)

type playerDal_ struct {
}

var PlayerDal *playerDal_ = &playerDal_{}

func (this *playerDal_) GetList(playerId string) (*dbUtil.DataTable, error) {
	sql := "select * from p_player where id=?"
	result, errMsg := dataBase.GameDb().Query(sql, playerId)
	if errMsg != nil {
		return nil, errMsg
	}

	return dbUtil.NewDataTable(result)
}

func (this *playerDal_) Update(playerId string, name string, lv int32, exp int) int {
	return 1
}

func (this *playerDal_) Insert(playerId string, name string, lv int32, exp int) int {
	return 1
}

func (this *playerDal_) Delete(playerId string) int {
	return 1
}
