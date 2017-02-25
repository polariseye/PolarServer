package config

import (
	"github.com/Jordanzuo/goutil/dbUtil"
	"github.com/polariseye/polarserver/common/property"
)

// 玩家等级配置，就是经验达到多少才能升级
type PlayerLvConfig struct {
	Lv  int32
	Exp int32
}

func NewPlayerLvConfigFromRow(row *dbUtil.DataRow) (result *PlayerLvConfig, errMsg error) {
	result = &PlayerLvConfig{}

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
