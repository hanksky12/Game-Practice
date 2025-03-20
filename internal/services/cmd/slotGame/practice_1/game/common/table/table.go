package table

import (
	"gamePractice/internal/pkg/const/symbol"
	"gamePractice/internal/pkg/entity/game"
	"gamePractice/internal/pkg/entity/settings/line"
	payoutEntity "gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/services/cmd/slotGame/common/countSymbol/count"
	"gamePractice/internal/services/cmd/slotGame/common/freeGameTimes"
	"gamePractice/internal/services/cmd/slotGame/common/payout"
	"gamePractice/internal/services/cmd/slotGame/practice_1/game/common/table/scoring"
	log "github.com/sirupsen/logrus"
)

type Base struct {
	Board        *game.Board
	IsPrintMore  bool
	MinimumCount int
}

func (b *Base) CalculateWinScore(singleBet float64, payoutTable *payoutEntity.Table, lineTable *line.Table) float64 {
	//log.Info("~~盤面計分~~")
	singleGameScore := 0.0
	scoringLine := scoring.Line{}
	payout := payout.Payout{BetBase: float64(len(lineTable.Rows)), SingleBet: singleBet, PayoutTable: payoutTable}
	for _, lineRow := range lineTable.Rows {
		//log.Info("~~~~第", i+1, "種~~~~")
		count, matchingSymbol := scoringLine.GetCountAndSymbol(lineRow, b.Board.Reels)
		if b.IsPrintMore {
			items := scoringLine.GetReelItems(lineRow, b.Board.Reels)
			log.Info(" Case:", lineRow.Case, " 對應值:", items)
		}
		if count < b.MinimumCount {
			//log.Info("不足連線數")
			continue
		}
		// 計算得分
		lineScore := payout.QueryLineScore(count, matchingSymbol)
		//log.Info(" 符號:", matchingSymbol, " 連線數:", count, " 線得分:", lineScore)
		singleGameScore += lineScore
	}
	//log.Info("~~單局計分~~ SingleGameScore ", singleGameScore)
	return singleGameScore
}

func (b *Base) CalculateWinFreeGame(freeAddTimes [5]int) (bool, int) {
	count := count.Any(b.Board.Reels, symbol.SC)
	times := &freeGameTimes.Win{FreeAddTimes: freeAddTimes}
	return times.Get(count)
}
