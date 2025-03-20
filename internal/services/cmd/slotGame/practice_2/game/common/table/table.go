package table

import (
	symbolConst "gamePractice/internal/pkg/const/symbol"
	"gamePractice/internal/pkg/entity/game"
	payoutEntity "gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/services/cmd/slotGame/common/countSymbol/count"
	"gamePractice/internal/services/cmd/slotGame/common/freeGameTimes"
	"gamePractice/internal/services/cmd/slotGame/common/payout"
	"gamePractice/internal/services/cmd/slotGame/practice_2/game/common/table/scoring"
	log "github.com/sirupsen/logrus"
)

type Base struct {
	Board        *game.Board
	MinimumCount int
}

func (b *Base) CalculateWinScore(singleBet float64, payoutTable *payoutEntity.Table) float64 {
	log.Info("~~盤面計分~~")
	singleGameScore := 0.0
	payout := payout.Payout{BetBase: 50.0, SingleBet: singleBet, PayoutTable: payoutTable}
	scoringWay := scoring.Way{}
	alreadyCalculatedSymbol := make(map[string]bool) //算過的避免重複計算
	//只需檢查第一個Reel
	for i, symbol := range b.Board.Reels[0].Items {
		log.Info("~~~~第", i+1, "種~~~~", symbol)
		if symbol == symbolConst.SC || symbol == symbolConst.W {
			log.Info(symbol, " 沒有賠付")
			continue
		}
		if alreadyCalculatedSymbol[symbol] {
			log.Info(symbol, " 已經算過")
			continue
		}
		count, way := scoringWay.GetCountAndWay(symbol, b.Board.Reels)
		if count < b.MinimumCount {
			log.Info("不足連線數")
			continue
		}
		// 計算得分
		wayScore := payout.QueryWayScore(count, symbol, way)
		log.Info(" 符號:", symbol, " 連線數:", count, " way:", way, " 得分:", wayScore)
		singleGameScore += wayScore
		alreadyCalculatedSymbol[symbol] = true
	}
	return singleGameScore
}

func (b *Base) CalculateWinFreeGame(freeAddTimes [5]int) (bool, int) {
	log.Info("~~計算贏得免費場次~~")
	count := count.Any(b.Board.Reels, symbolConst.SC)
	times := &freeGameTimes.Win{FreeAddTimes: freeAddTimes}
	ok, countWin := times.Get(count)
	return ok, countWin
}
