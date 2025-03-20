package game

import (
	"fmt"
	"gamePractice/internal/pkg/entity/game/param"
	settingsEntity "gamePractice/internal/pkg/entity/settings"
	_interface "gamePractice/internal/pkg/interface"
	"gamePractice/internal/services/cmd/slotGame/common/user"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func LoadLineSettingsTables(reader _interface.IReadLine) (*settingsEntity.LineTables, *settingsEntity.LineTables, bool) {
	ok, MgWheelTable, FgWheelTable := reader.ReadWheelTable()
	if !ok {
		log.Error("ReadWheelTable error")
		return nil, nil, false
	}
	ok, PayTable := reader.ReadPayTable()
	if !ok {
		log.Error("ReadPayTable error")
		return nil, nil, false
	}
	ok, LineTable := reader.ReadLineTable()
	if !ok {
		log.Error("ReadLineTable error")
		return nil, nil, false
	}
	return &settingsEntity.LineTables{
			Wheel: MgWheelTable,
			Pay:   PayTable,
			Line:  LineTable,
		}, &settingsEntity.LineTables{
			Wheel: FgWheelTable,
			Pay:   PayTable,
			Line:  LineTable,
		}, true
}

func LoadWaySettingsTables(reader _interface.IReadWay) (*settingsEntity.WayTables, *settingsEntity.WayTables, bool) {
	ok, MgWheelTable, FgWheelTable := reader.ReadWheelTable()
	if !ok {
		log.Error("ReadWheelTable error")
		return nil, nil, false
	}
	ok, PayTable := reader.ReadPayTable()
	if !ok {
		log.Error("ReadPayTable error")
		return nil, nil, false
	}
	return &settingsEntity.WayTables{
			Wheel: MgWheelTable,
			Pay:   PayTable,
		}, &settingsEntity.WayTables{
			Wheel: FgWheelTable,
			Pay:   PayTable,
		}, true
}

func SpinByTimes(game _interface.ISpin, param *param.Parameter, singleUser *user.User, totalTimes int) {
	/*
		計時,RTP,次數
	*/
	//before appFrame.Run  LogOptions level = warn
	log.Info("初始金額 ", singleUser.Balance, " 下注 ", param.SingleBet)
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < totalTimes; i++ {
		wg.Add(1)
		go game.Spin(&wg)
	}
	wg.Wait()
	Win := singleUser.Win
	Bet := param.SingleBet * float64(totalTimes)
	RTP := Win / Bet
	percentageRTP := fmt.Sprintf("%.3f%%", RTP*100)
	log.Warn("TotalTimes: ", totalTimes, " Win: ", Win, " Bet: ", Bet, " RTP: ", percentageRTP)
	elapsed := time.Since(start) // 計算時間差
	log.Warn("執行時間:", elapsed)
}
