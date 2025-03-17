package controller

import (
	"fmt"
	"gamePractice/internal/pkg/interface"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/user"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/settings"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

type Job struct{}

func (j *Job) LineGamePractice(isMock string) {
	/*
		Shape 3*5
		Sc*3 Up(不管連續) => +free 10
		Symbol*2 Up => payout
		max free =250
		W 百搭(L,H,not Sc)
		1百萬次=>1min(47s)
	*/
	log.Info("LineGamePractice: start")
	//totalTimes := 1000000
	totalTimes := 1000000
	singleUser := &user.User{
		Id:      1,
		Balance: 0.0,
		Win:     0.0,
	}
	singleBet := 10.0
	shape := []int{3, 3, 3, 3, 3}
	log.Info("初始金額 ", singleUser.Balance, " 下注 ", singleBet)
	var settingDoc _interface.IReadLine
	if isMock == "true" {
		settingDoc = &settings.MockDoc{} //自訂 高機率觸發FreeGame
	} else {
		basePath, _ := os.Getwd()
		settingDoc = &settings.Doc{BasePath: basePath}
	}
	lineGame := &lineGamePractice.LineGamePractice{}
	if !lineGame.Init(settingDoc) {
		log.Info("讀取資料失敗")
		return
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < totalTimes; i++ {
		//log.Info("第", i+1, "次~~~")
		wg.Add(1)
		go lineGame.Spin(singleUser, singleBet, shape, &mu, &wg)
	}
	wg.Wait()
	Win := singleUser.Win
	Bet := singleBet * float64(totalTimes)
	RTP := Win / Bet
	percentageRTP := fmt.Sprintf("%.3f%%", RTP*100)
	log.Info("Win ", Win, " Bet ", Bet, " RTP ", percentageRTP)
}
