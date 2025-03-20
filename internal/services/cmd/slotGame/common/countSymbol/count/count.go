package count

import (
	"gamePractice/internal/pkg/entity/settings/reel"
)

func Any(reels []*reel.Reel, s string) int {
	reelCount := 0 // 計算有出現 symbol 的 Reel 數量

	for _, reel := range reels {
		for _, item := range reel.Items {
			if item == s {
				reelCount++ // 只要該 Reel 內出現過一次，就計數
				break       // 跳出內部迴圈，避免重複計算
			}
		}
	}

	return reelCount
}
