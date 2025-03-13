package controller

import (
	"gamePractice/internal/pkg/interface"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/user"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/settings"
	log "github.com/sirupsen/logrus"
)

type Job struct{}

func (j *Job) LineGamePractice(isMock string) {
	/*
		Shape 3*5
		Sc*3 Up => +free 10
		Symbol*3 Up => payout
		max free =30
		W 百搭(L,H,Sc)
	*/
	log.Info("LineGamePractice: start")
	user := &user.User{
		Id:      1,
		Balance: 1100.0,
	}
	singleBet := 100.0
	shape := []int{3, 3, 3, 3, 3}
	log.Info("初始金額 ", user.Balance, " 下注 ", singleBet)

	var settingDoc _interface.IReadLine
	if isMock == "true" {
		settingDoc = &settings.MockDoc{} //自訂 高機率觸發FreeGame
	} else {
		settingDoc = &settings.Doc{}
	}
	lineGame := &lineGamePractice.LineGamePractice{}
	if !lineGame.Init(settingDoc) {
		log.Info("讀取資料失敗")
		return
	}
	lineGame.Spin(user, singleBet, shape)
}
