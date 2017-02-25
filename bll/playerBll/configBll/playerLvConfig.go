package configBll

import (
	"github.com/Jordanzuo/goutil/validationUtil"
	"github.com/polariseye/polarserver/bll/instanceMangeBll/modelData"
	"github.com/polariseye/polarserver/dal/playerDal/configDal"
	"github.com/polariseye/polarserver/model/player/config"
)

type playerLvConfigStruct struct {
	playerLvConfigData map[int32]*config.PlayerLvConfig
}

var PlayerLvConfigBll *playerLvConfigStruct

func init() {
	PlayerLvConfigBll = newPlayerLvConfigBll()
	modelData.Register(PlayerLvConfigBll)
}

func (this *playerLvConfigStruct) Init() []error {
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

func (this *playerLvConfigStruct) Check() []error {
	errList := make([]error, 0)

	for _, val := range this.playerLvConfigData {
		validationUtil.CheckIntRange(&errList, int(val.Exp), 0,
			validationUtil.MaxInt32, "数据库表b_player_lv_config表中 Exp必须大于0 ")
	}

	return errList
}

func (this *playerLvConfigStruct) Convert() []error {
	return nil
}

func (this *playerLvConfigStruct) ModuleName() string {
	return "PlayerLvConfigBll"
}

func newPlayerLvConfigBll() *playerLvConfigStruct {
	return &playerLvConfigStruct{
		playerLvConfigData: make(map[int32]*config.PlayerLvConfig),
	}
}
