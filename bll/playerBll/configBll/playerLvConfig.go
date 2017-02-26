package configBll

import (
	"github.com/Jordanzuo/goutil/validationUtil"
	"github.com/polariseye/polarserver/bll/instanceMangeBll/modelData"
	"github.com/polariseye/polarserver/dal/playerDal/configDal"
	"github.com/polariseye/polarserver/model/player/config"
)

type playerLvConfigBll_ struct {
	playerLvConfigData map[int32]*config.PlayerLvConfig
}

var PlayerLvConfigBll *playerLvConfigBll_

func init() {
	PlayerLvConfigBll = newPlayerLvConfigBll()
	modelData.Register(PlayerLvConfigBll)
}

func (this *playerLvConfigBll_) Init() []error {
	errList := make([]error, 0)

	// 获取数据
	table, errMsg := configDal.PlayerLvConfigDal.GetList()
	if errMsg != nil {
		errList = append(errList, errMsg)
		return errList
	}

	// 转换成字典
	playerLvConfigData := make(map[int32]*config.PlayerLvConfig, 0)
	for i := 0; i < table.RowCount(); i++ {
		row, _ := table.Row(i)

		// 构造实体
		var item *config.PlayerLvConfig
		if item, errMsg = config.NewPlayerLvConfigFromRow(row); errMsg != nil {
			errList = append(errList, errMsg)

			// 如果失败，则跳过此实体
			continue
		}

		playerLvConfigData[item.Lv] = item
	}

	// 保存到缓存
	this.playerLvConfigData = playerLvConfigData

	return nil
}

func (this *playerLvConfigBll_) Check() []error {
	errList := make([]error, 0)

	for _, val := range this.playerLvConfigData {
		validationUtil.CheckIntRange(&errList, int(val.Exp), 0,
			validationUtil.MaxInt32, "数据库表b_player_lv_config表中 Exp必须大于0 ")
	}

	return errList
}

func (this *playerLvConfigBll_) Convert() []error {
	return nil
}

func (this *playerLvConfigBll_) ModuleName() string {
	return "PlayerLvConfigBll"
}

func newPlayerLvConfigBll() *playerLvConfigBll_ {
	return &playerLvConfigBll_{
		playerLvConfigData: make(map[int32]*config.PlayerLvConfig),
	}
}
