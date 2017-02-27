package player

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
