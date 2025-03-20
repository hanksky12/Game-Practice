package gameBoard

import (
	"gamePractice/internal/pkg/entity/game"
	"gamePractice/internal/pkg/entity/settings/reel"
	rand "gamePractice/internal/pkg/util/rand"
	log "github.com/sirupsen/logrus"
)

type Board struct {
	Name  string
	Shape []int
	board *game.Board
}

func (b *Board) GetByReel(table *reel.Table, rand *rand.SafeRand) *game.Board {
	b.createEmpty()
	b.fill(table, rand)
	log.Info(b.board)
	return b.board
}

func (b *Board) createEmpty() {
	b.board = &game.Board{
		Name:  b.Name,
		Shape: b.Shape,
		Reels: make([]*reel.Reel, len(b.Shape)),
	}
}

func (b *Board) fill(Table *reel.Table, rand *rand.SafeRand) {
	//log.Info("~~生成盤面~~")
	if Table == nil || len(Table.Reels) == 0 {
		log.Error("fill: Table.Reels 為空，無法生成盤面")
		return
	}
	for index, reelLen := range b.board.Shape {
		sourceReel := Table.Reels[index] // 取得輪帶
		//log.Info("~~取輪帶 ", sourceReel)
		selectedSymbols := b.selectSymbols(sourceReel.Items, reelLen, rand)
		//log.Info("~~生成Reel ", selectedSymbols)
		b.board.Reels[index] = &reel.Reel{Items: selectedSymbols}
	}
	//log.Info(b.board)
	//log.Info("~~完成盤面~~")
}

func (b *Board) selectSymbols(items []string, length int, rand *rand.SafeRand) []string {
	if len(items) == 0 || length <= 0 {
		log.Warn("selectSymbols: items 切片為空或選取長度為0")
		return nil
	}
	startIndex := rand.Intn(len(items)) // 隨機loc //startIndex := 101
	//log.Info("Start=>", startIndex)
	symbolSlice := make([]string, length)
	for i := 0; i < length; i++ {
		remainder := (startIndex + i) % len(items) // 餘數
		//log.Info("Case ", remainder)
		symbolSlice[i] = items[remainder]
	}
	return symbolSlice
}
