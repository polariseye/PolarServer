package playerBll

import (
	"fmt"

	"github.com/Jordanzuo/goutil/dbUtil"
	"github.com/polariseye/polarserver/bll/instanceMangeBll/playerData"
	"github.com/polariseye/polarserver/dal/playerDal"
	"github.com/polariseye/polarserver/model/player"
)

type playerBll_ struct {
}

var PlayerBll *playerBll_

func init() {
	PlayerBll = newPlayerBll()
	playerData.Register(PlayerBll)
}

func (this *playerBll_) ModuleName() string {
	return "PlayerBll"
}

func (this *playerBll_) init(playerId string) (result *player.Player, errMsg error) {

	// 获取数据
	var dt *dbUtil.DataTable
	dt, errMsg = playerDal.PlayerDal.GetList(playerId)
	if errMsg != nil {
		return
	}

	if dt.RowCount() <= 0 {
		errMsg = fmt.Errorf("未找到PlayerId=%s的玩家信息", playerId)

		return
	}

	row, _ := dt.Row(0)

	return player.NewPlayerFromRow(row)
}

func (this *playerBll_) GetPlayer(playerId string) (result *player.Player, errMsg error) {
	//todo:如何查找玩家对象
}

func newPlayerBll() *playerBll_ {
	return &playerBll_{}
}
