package configBll

import (
	"github.com/polariseye/polarserver/model/rankScore/config"
)

type rankScoreRelationBllStruct struct {
	tmpData               []*config.RankScoreRelation
	rankScoreRelationData []*config.RankScoreRelation
}

var RankScoreRelationBll *rankScoreRelationBllStruct

func init() {
	RankScoreRelationBll = newRankScoreRelationBll()
}

func (this *rankScoreRelationBllStruct) ModuleName() string {
	return "RankScoreRelationBll"
}

func (this *rankScoreRelationBllStruct) Init() []error {
	return nil
}

func (this *rankScoreRelationBllStruct) Check() []error {
	return nil
}

func (this *rankScoreRelationBllStruct) Convert() []error {
	return nil
}

func (this *rankScoreRelationBllStruct) Confirm() {
	this.rankScoreRelationData = this.tmpData
}

func (this *rankScoreRelationBllStruct) onReloadFinish() {

}

func (this *rankScoreRelationBllStruct) GetData() []*config.RankScoreRelation {
	return this.rankScoreRelationData
}

func (this *rankScoreRelationBllStruct) GetItem(id int32) *config.RankScoreRelation {
	for _, item := range this.rankScoreRelationData {
		if item.Id == id {
			return item
		}
	}

	return nil
}

func newRankScoreRelationBll() *rankScoreRelationBllStruct {
	return new(rankScoreRelationBllStruct)
}
