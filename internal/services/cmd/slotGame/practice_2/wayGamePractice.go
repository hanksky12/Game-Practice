package practice_2

import (
	gameEntity "gamePractice/internal/pkg/entity/game/param"
	settingsEntity "gamePractice/internal/pkg/entity/settings"
	_interface "gamePractice/internal/pkg/interface"
	gameCommon "gamePractice/internal/services/cmd/slotGame/common/game"
	"gamePractice/internal/services/cmd/slotGame/common/session"
	settingsCommon "gamePractice/internal/services/cmd/slotGame/common/settings"
	"gamePractice/internal/services/cmd/slotGame/common/user"
	"gamePractice/internal/services/cmd/slotGame/practice_2/game"
	"gamePractice/internal/services/cmd/slotGame/practice_2/settings"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

type WayGamePractice struct {
	GameParam        *gameEntity.Parameter
	User             *user.User
	mgSettingsTables *settingsEntity.WayTables
	fgSettingsTables *settingsEntity.WayTables
}

func (l *WayGamePractice) Spin(wg *sync.WaitGroup) {
	log.Info("Spin")
	defer wg.Done()

	s := &session.Session{MaxFreeGameTimes: l.GameParam.MaxFreeGameTimes}
	s.Init()

	m := game.Main{GameParam: l.GameParam, User: l.User, Session: s}
	m.Run(l.mgSettingsTables)

	for {
		IsAnyFree, whichNumTimes := s.IsAnyFreeGameTimes()
		if !IsAnyFree {
			break
		}
		log.Info("免費場次第", whichNumTimes, " 次")
		f := game.Free{Session: s, User: l.User, GameParam: l.GameParam}
		f.Run(l.fgSettingsTables)
		log.Info("FG玩家Win ", l.User.Win)
	}

	log.Info("玩家總得分 ", l.User.Balance)
}

func (l *WayGamePractice) Init(isMock bool) bool {
	log.Info("Init")
	var reader _interface.IReadWay
	if isMock == true {
		reader = &settings.MockDoc{} //自訂
	} else {
		basePath, _ := os.Getwd()
		reader = &settings.Doc{BaseDoc: &settingsCommon.Doc{}, BasePath: basePath}
	}
	mgTables, fgTables, ok := gameCommon.LoadWaySettingsTables(reader)
	if !ok {
		return false
	}
	l.mgSettingsTables = mgTables
	l.fgSettingsTables = fgTables
	return true
}
