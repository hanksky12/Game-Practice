package game

import (
	constGame "gamePractice/internal/pkg/const/game"
	"gamePractice/internal/pkg/entity/game"
	"gamePractice/internal/pkg/entity/settings"
	"gamePractice/internal/pkg/entity/settings/reel"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/session"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/table"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/user"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type Free struct {
	Shape   []int
	Session *session.Session
	User    *user.User
}

func (f *Free) Run(singleBet float64, settingsTables *settings.Tables) {
	log.Info("----------------------FG start----------------------")
	board := &game.Board{
		Name:  constGame.Free,
		Shape: f.Shape,
		Reels: make([]*reel.Reel, len(f.Shape)),
	}
	///////一樣
	base := table.Base{Board: board, Rand: rand.New(rand.NewSource(time.Now().UnixNano())), IsPrintMore: true}
	base.FillGameBoard(settingsTables.Wheel)
	score := base.CalculateWinScore(singleBet, settingsTables.Pay, settingsTables.Line)
	isWin, times := base.CalculateWinFreeGame()
	if isWin {
		log.Info("贏得免費場次 ", times)
		f.Session.AddFreeGameTimes(times)
	}
	///////一樣
	f.User.AddScore(score)
	log.Info("本局贏分 ", score, " 玩家累計總贏分 ", f.User.Balance)
	log.Info("----------------------FG end----------------------")
}
