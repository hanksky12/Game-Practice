package scoring

import (
	"gamePractice/internal/pkg/const/symbol"
	"gamePractice/internal/pkg/entity/settings/line"
	"gamePractice/internal/pkg/entity/settings/reel"
)

type Line struct {
}

func (line *Line) GetCountAndSymbol(lineRow *line.Row, reels []*reel.Reel) (int, string) {
	count := 1
	continuousSymbol := ""
	for reelIndex, itemIndex := range lineRow.Case {
		//拿出跨輪帶對應的位置
		item := reels[reelIndex].Items[itemIndex]
		if reelIndex == 0 {
			continuousSymbol = item
			continue
		}

		// 如果前一個是 W，則繼續累積並取代符號
		if continuousSymbol == symbol.W {
			continuousSymbol = item
			count++
			continue
		}

		// 相同符號||當前符號是 W 繼續累積
		if item == continuousSymbol || item == symbol.W {
			count++
			continue
		}
		// 連線斷開，直接返回當前 count 和 symbol
		break
	}
	return count, continuousSymbol
}

func (l *Line) GetReelItems(lineRow *line.Row, reels []*reel.Reel) []string {
	var items []string
	for reelIndex, itemIndex := range lineRow.Case {
		items = append(items, reels[reelIndex].Items[itemIndex])
	}
	return items
}
