package table

import (
	"gamePractice/internal/pkg/const/symbol"
	"gamePractice/internal/pkg/entity/game"
	"gamePractice/internal/pkg/entity/settings/line"
	"gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/table/count"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/table/scoring"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

type Base struct {
	Board       *game.Board
	Rand        *rand.Rand
	IsPrintMore bool
}

func (b *Base) FillGameBoard(table *reel.Table) {
	log.Info("~~生成盤面~~")
	for index, reelLen := range b.Board.Shape {
		sourceReel := table.Reels[index] // 取得輪帶
		//log.Info("~~生成輪帶 ", sourceReel.Name)
		selectedSymbols := b.selectSymbols(sourceReel.Items, reelLen)
		b.Board.Reels[index] = &reel.Reel{Items: selectedSymbols}
	}
	log.Info(b.Board)
	log.Info("~~完成盤面~~")
}

func (b *Base) CalculateWinScore(singleBet float64, payoutTable *payout.Table, lineTable *line.Table) float64 {
	log.Info("~~盤面計分~~")
	var singleGameScore = 0.0
	var minimumCount = 3
	scoringLine := scoring.Line{LineTable: lineTable, PayoutTable: payoutTable}
	scoringLine.Init()
	for i, lineRow := range lineTable.Rows {
		log.Info("~~~~第", i+1, "種~~~~")
		count, matchingSymbol := scoringLine.GetCountAndSymbol(lineRow, b.Board.Reels)
		if b.IsPrintMore {
			items := scoringLine.GetReelItems(lineRow, b.Board.Reels)
			log.Info(" Case:", lineRow.Case, " 對應值:", items)
		}
		if count < minimumCount {
			log.Info("不足連線數")
			continue
		}
		// 計算得分
		lineScore := scoringLine.QueryScore(count, matchingSymbol, singleBet)
		log.Info(" 符號:", matchingSymbol, " 連線數:", count, " 線得分:", lineScore)
		singleGameScore += lineScore
	}
	log.Info("~~單局計分~~ SingleGameScore ", singleGameScore)
	return singleGameScore
}

func (b *Base) CalculateWinFreeGame() (bool, int) {
	count := count.Symbol(b.Board.Reels, symbol.Sc)
	log.Info("Scatter count ", count)
	if count >= 3 {
		return true, 10
	}
	return false, 0
}

func (b *Base) selectSymbols(items []string, length int) []string {
	if len(items) == 0 || length <= 0 {
		return nil
	}
	startIndex := b.Rand.Intn(len(items)) // 隨機loc //startIndex := 101
	//log.Info("Start=>", startIndex)
	symbolSlice := make([]string, length)
	for i := 0; i < length; i++ {
		remainder := (startIndex + i) % len(items) // 餘數
		//log.Info("Case ", remainder)
		symbolSlice[i] = items[remainder]
	}
	return symbolSlice
}
