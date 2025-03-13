package lineGamePractice

import (
	"gamePractice/internal/pkg/entity/settings"
	"gamePractice/internal/pkg/interface"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/session"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/user"
	log "github.com/sirupsen/logrus"
)

type LineGamePractice struct {
	MgSettingsTables *settings.Tables
	FgSettingsTables *settings.Tables
}

func (l *LineGamePractice) Spin(user *user.User, singleBet float64, shape []int) {
	log.Info("LineGamePractice spin")
	s := &session.Session{MaxFreeGameTimes: 30}
	s.Init()

	m := game.Main{Shape: shape, Session: s, User: user}
	m.Run(singleBet, l.MgSettingsTables)

	for {
		IsAnyFree, whichNumTimes := s.IsAnyFreeGameTimes()
		if !IsAnyFree {
			break
		}
		log.Info("免費場次第", whichNumTimes, " 次")
		free := game.Free{Shape: shape, Session: s, User: user}
		free.Run(singleBet, l.FgSettingsTables)
	}

	log.Info("玩家總得分 ", user.Balance)
}

func (l *LineGamePractice) Init(s _interface.IReadLine) bool {
	log.Info("Init")
	ok, MgWheelTable, FgWheelTable := s.ReadWheelTable()
	if !ok {
		log.Error("ReadWheelTable error")
		return false
	}
	ok, PayTable := s.ReadPayTable()
	if !ok {
		log.Error("ReadPayTable error")
		return false
	}
	ok, LineTable := s.ReadLineTable()
	if !ok {
		log.Error("ReadLineTable error")
		return false
	}
	l.MgSettingsTables = &settings.Tables{
		Wheel: MgWheelTable,
		Pay:   PayTable,
		Line:  LineTable,
	}
	l.FgSettingsTables = &settings.Tables{
		Wheel: FgWheelTable,
		Pay:   PayTable,
		Line:  LineTable,
	}
	return true
}
