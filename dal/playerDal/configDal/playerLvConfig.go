package configDal

import (
	"database/sql"

	"github.com/polariseye/polarserver/dal/dataBase"
	"github.com/polariseye/polarserver/model/player/config"
)

type playerLvConfigDal_ struct {
}

var PlayerLvConfigDal *playerLvConfigDal_ = &playerLvConfigDal_{}

func (this *playerLvConfigDal_) GetList() (result []*config.PlayerLvConfig, errMsg error) {
	toExecuteSql := "select * from b_player_lv_c"
	var rows *sql.Rows
	rows, errMsg = dataBase.GameDb().Query(toExecuteSql)
	if errMsg != nil {
		return nil, errMsg
	}

	result = make([]*config.PlayerLvConfig, 0)
	defer func() {
		rows.Close()
	}()

	for rows.Next() {
		item := config.NewPlayerLvConfig()
		rows.Scan(&item.Exp, &item.Lv)

		result = append(result, item)
	}

	return
}
