package player

import (
	"github.com/Jordanzuo/goutil/dbUtil"
	"github.com/polariseye/polarserver/common/property"
)

// 玩家基本信息
type Player struct {
	Id        string
	Name      string
	ServerId  int32
	PartnerId int32
	UserId    string
	Lv        int32
	Exp       int32
}

func NewPlayer() *Player {
	return &Player{}
}

func NewPlayerFromRow(row *dbUtil.DataRow) (result *Player, errMsg error) {
	result = &Player{}

	if result.Id, errMsg = dbUtil.String(row, property.Id); errMsg != nil {
		result = nil
		return
	}

	if result.Name, errMsg = dbUtil.String(row, property.Name); errMsg != nil {
		result = nil
		return
	}

	if result.ServerId, errMsg = dbUtil.Int32(row, property.ServerId); errMsg != nil {
		result = nil
		return
	}

	if result.PartnerId, errMsg = dbUtil.Int32(row, property.PartnerId); errMsg != nil {
		result = nil
		return
	}

	if result.UserId, errMsg = dbUtil.String(row, property.UserId); errMsg != nil {
		result = nil
		return
	}

	if result.Lv, errMsg = dbUtil.Int32(row, property.Lv); errMsg != nil {
		result = nil
		return
	}

	if result.Exp, errMsg = dbUtil.Int32(row, property.Exp); errMsg != nil {
		result = nil
		return
	}

	return
}
