package lineGamePractice

import (
	"gamePractice/internal/pkg/entity/settings"
	"gamePractice/internal/pkg/interface"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/session"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/user"
	log "github.com/sirupsen/logrus"
	"sync"
)

type LineGamePractice struct {
	MgSettingsTables *settings.Tables
	FgSettingsTables *settings.Tables
}

func (l *LineGamePractice) Spin(user *user.User, singleBet float64, shape []int, mu *sync.Mutex, wg *sync.WaitGroup) {
	log.Info("LineGamePractice spin")
	defer wg.Done()
	s := &session.Session{MaxFreeGameTimes: 250}
	s.Init()
	isPrintMore := false

	m := game.Main{Shape: shape, Session: s, User: user}
	m.Run(singleBet, l.MgSettingsTables, mu, isPrintMore)
	//log.Info("MG玩家Win ", user.Win)
	var wgf sync.WaitGroup
	for {
		IsAnyFree, whichNumTimes := s.IsAnyFreeGameTimes()
		if !IsAnyFree {
			break
		}
		log.Info("免費場次第", whichNumTimes, " 次")
		wgf.Add(1)
		go func(whichNumTimes int) {
			defer wgf.Done()
			f := game.Free{Shape: shape, Session: s, User: user}
			f.Run(singleBet, l.FgSettingsTables, mu, isPrintMore)
		}(whichNumTimes)
		//free := game.Free{Shape: shape, Session: s, User: user}
		//free.Run(singleBet, l.FgSettingsTables, mu, isPrintMore)
		//log.Info("FG玩家Win ", user.Win)
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
