package rankScore

import (
	"time"

	"github.com/polariseye/polarserver/model/player"
)

type PlayerRankScore struct {
	Player          *player.Player
	Score           int32
	ScoreUpdateTime time.Time
}
