package scoring

import (
	symbolConst "gamePractice/internal/pkg/const/symbol"
	"gamePractice/internal/pkg/entity/settings/reel"
)

type Way struct {
}

func (w *Way) GetCountAndWay(symbol string, reels []*reel.Reel) (int, float64) {
	ways := 1
	reelCount := 0 // 連續出現的 Reel 數量
	for i := 0; i < len(reels); i++ {
		//log.Info("~~~~第", i+1, "輪~~~~")
		count := 0 // 該 Reel 出現的次數(累積後for 相乘)
		for _, item := range reels[i].Items {
			if item == symbol || item == symbolConst.W {
				count++
			}
		}
		if count == 0 {
			break
		}
		ways *= count // 組合數
		reelCount++
	}
	return reelCount, float64(ways)
}
