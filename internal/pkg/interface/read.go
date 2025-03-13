package _interface

import (
	"gamePractice/internal/pkg/entity/settings/line"
	"gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
)

type IReadBase interface {
	ReadWheelTable() (bool, *reel.Table, *reel.Table)
	ReadPayTable() (bool, *payout.Table)
}

type IReadLine interface {
	IReadBase
	ReadLineTable() (bool, *line.Table)
}

//type IReadWay interface {
//	IReadBase
//}
