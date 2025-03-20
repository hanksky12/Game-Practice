package game

import (
	constGame "gamePractice/internal/pkg/const/game"
	gameEntity "gamePractice/internal/pkg/entity/game/param"
	"gamePractice/internal/pkg/entity/settings"
	"gamePractice/internal/services/cmd/slotGame/common/gameBoard"
	"gamePractice/internal/services/cmd/slotGame/common/session"
	"gamePractice/internal/services/cmd/slotGame/common/user"
	"gamePractice/internal/services/cmd/slotGame/practice_2/game/common/table"
	log "github.com/sirupsen/logrus"
)

type Main struct {
	Session   *session.Session
	User      *user.User
	GameParam *gameEntity.Parameter
}

func (m *Main) Run(settingsTables *settings.WayTables) {
	log.Info("----------------------MG start----------------------")
	game := gameBoard.Board{Name: constGame.Main, Shape: m.GameParam.BoardShape.Mg}
	board := game.GetByReel(settingsTables.Wheel, m.GameParam.Rand)
	base := table.Base{Board: board, MinimumCount: m.GameParam.PayOutMinimumCount.Mg}
	score := base.CalculateWinScore(m.GameParam.SingleBet, settingsTables.Pay)
	isWin, times := base.CalculateWinFreeGame(m.GameParam.FreeAddTimes.Mg)

	m.User.AddScore(score, m.GameParam.SingleBet)

	if isWin {
		log.Info("贏得免費場次 ", times)
		m.Session.AddFreeGameTimes(times)
	}
	log.Info("本局贏分 ", score, " 玩家累計總贏分 ", m.User.Balance)
	log.Info("----------------------MG end----------------------")
}
