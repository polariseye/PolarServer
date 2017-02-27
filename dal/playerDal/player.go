package playerDal

import (
	"database/sql"

	"github.com/polariseye/polarserver/dal/dataBase"
	"github.com/polariseye/polarserver/model/player"
)

type playerDal_ struct{}

var PlayerDal *playerDal_ = &playerDal_{}

func (this *playerDal_) GetList(playerId string) (result []*player.Player, errMsg error) {
	executeSql := "select * from p_player where id=?"
	var rows *sql.Rows
	rows, errMsg = dataBase.GameDb().Query(executeSql, playerId)
	if errMsg != nil {
		return nil, errMsg
	}

	result = make([]*player.Player, 0)
	defer func() {
		rows.Close()
	}()

	for rows.Next() {
		item := player.NewPlayer()
		rows.Scan(&item.Exp, &item.Id, &item.Lv, &item.Name, &item.PartnerId, &item.ServerId, &item.UserId)

		result = append(result, item)
	}

	return
}

func (this *playerDal_) Update(item *player.Player) int {
	return 1
}

func (this *playerDal_) Insert(item *player.Player) int {
	return 1
}

func (this *playerDal_) Delete(playerId string) int {
	return 1
}
