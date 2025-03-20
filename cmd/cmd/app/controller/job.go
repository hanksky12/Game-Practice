package controller

import (
	"gamePractice/internal/pkg/entity/game/param"
	randUtil "gamePractice/internal/pkg/util/rand"
	"gamePractice/internal/services/cmd/slotGame/common/game"
	"gamePractice/internal/services/cmd/slotGame/common/user"
	"gamePractice/internal/services/cmd/slotGame/practice_1"
	"gamePractice/internal/services/cmd/slotGame/practice_2"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type Job struct{}

// Practice1
/*
	lineGameBasic
	Shape 3*5
	Sc*3 Up(不管連續) => MG,FG+free 10
	Symbol*2 Up => payout
	max free =250
	W 百搭(L,H,not Sc)
	1百萬次=>RTP 97.287%
*/
func (j *Job) Practice1() {
	log.Info("Practice1: start")
	singleUser := &user.User{
		Id:      1,
		Balance: 0.0,
		Win:     0.0,
	}
	gameParam := &param.Parameter{
		SingleBet:        10.0,
		MaxFreeGameTimes: 250,
		Rand:             &randUtil.SafeRand{Rand: rand.New(rand.NewSource(time.Now().UnixNano()))},
		BoardShape: &param.BoardShape{
			Mg: []int{3, 3, 3, 3, 3},
			Fg: []int{3, 3, 3, 3, 3},
		},
		FreeAddTimes: &param.FreeAddTimes{
			Mg: [5]int{0, 0, 10, 10, 10},
			Fg: [5]int{0, 0, 10, 10, 10},
		},
		PayOutMinimumCount: &param.PayOutMinimumCount{
			Mg: 2,
			Fg: 2,
		},
	}
	practiceGame := &practice_1.LineGamePractice{GameParam: gameParam, User: singleUser}
	if !practiceGame.Init(false) {
		log.Error("讀取資料失敗")
		return
	}
	game.SpinByTimes(practiceGame, gameParam, singleUser, 1000000) //1000000
}

// Practice2
/*
	WayGameBasic
	Shape 3*5
	Sc*3 Up(不管連續) => MG,FG +free 10, 15, 20
	Symbol*3 Up => payout
	max free =60
	W 百搭(L,H,not Sc)
	1百萬次=>RTP 50.88%
*/
func (j *Job) Practice2() {
	log.Info("Practice2: start")
	singleUser := &user.User{
		Id:      1,
		Balance: 0.0,
		Win:     0.0,
	}
	gameParam := &param.Parameter{
		SingleBet:        50.0,
		MaxFreeGameTimes: 60,
		Rand:             &randUtil.SafeRand{Rand: rand.New(rand.NewSource(time.Now().UnixNano()))},
		BoardShape: &param.BoardShape{
			Mg: []int{3, 3, 3, 3, 3},
			Fg: []int{3, 3, 3, 3, 3},
		},
		FreeAddTimes: &param.FreeAddTimes{
			Mg: [5]int{0, 0, 10, 15, 20},
			Fg: [5]int{0, 0, 10, 15, 20},
		},
		PayOutMinimumCount: &param.PayOutMinimumCount{
			Mg: 3,
			Fg: 3,
		},
	}
	practiceGame := &practice_2.WayGamePractice{GameParam: gameParam, User: singleUser}
	if !practiceGame.Init(false) {
		log.Error("讀取資料失敗")
		return
	}
	game.SpinByTimes(practiceGame, gameParam, singleUser, 1000000) //1000000
}

// Practice3
/*

 */
func (j *Job) Practice3() {
	log.Info("Practice3: start")

}
