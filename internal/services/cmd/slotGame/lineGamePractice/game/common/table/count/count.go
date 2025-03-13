package count

import (
	"gamePractice/internal/pkg/const/symbol"
	"gamePractice/internal/pkg/entity/settings/reel"
)

func Symbol(reels []*reel.Reel, s string) int {
	reelCount := 0 // 計算連續符合 symbol 的 Reel 數量
	for _, reel := range reels {
		count := 0 // 該條輪帶基準
		for _, item := range reel.Items {
			if item == s || item == symbol.W {
				count++
			}
		}
		if count == 0 {
			//找不到，代表跨輪帶停止計算
			break
		}
		reelCount++
	}
	return reelCount
}
