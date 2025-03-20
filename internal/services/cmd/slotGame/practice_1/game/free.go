package game

import (
	constGame "gamePractice/internal/pkg/const/game"
	gameEntity "gamePractice/internal/pkg/entity/game/param"
	"gamePractice/internal/pkg/entity/settings"
	"gamePractice/internal/services/cmd/slotGame/common/gameBoard"
	"gamePractice/internal/services/cmd/slotGame/common/session"
	"gamePractice/internal/services/cmd/slotGame/common/user"
	"gamePractice/internal/services/cmd/slotGame/practice_1/game/common/table"
	log "github.com/sirupsen/logrus"
)

type Free struct {
	Session   *session.Session
	User      *user.User
	GameParam *gameEntity.Parameter
}

func (f *Free) Run(settingsTables *settings.LineTables, isPrintMore bool) {
	log.Info("----------------------FG start----------------------")
	game := gameBoard.Board{Name: constGame.Free, Shape: f.GameParam.BoardShape.Fg}
	board := game.GetByReel(settingsTables.Wheel, f.GameParam.Rand)
	base := table.Base{Board: board, IsPrintMore: isPrintMore, MinimumCount: f.GameParam.PayOutMinimumCount.Fg}
	score := base.CalculateWinScore(f.GameParam.SingleBet, settingsTables.Pay, settingsTables.Line)
	isWin, times := base.CalculateWinFreeGame(f.GameParam.FreeAddTimes.Fg)

	f.User.AddScore(score, 0)
	if isWin {
		log.Info("贏得免費場次 ", times)
		f.Session.AddFreeGameTimes(times)
	}

	log.Info("本局贏分 ", score, " 玩家累計總贏分 ", f.User.Balance)
	log.Info("----------------------FG end----------------------")
}
