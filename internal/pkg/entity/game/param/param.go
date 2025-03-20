package param

import (
	"gamePractice/internal/pkg/util/rand"
)

type Parameter struct {
	SingleBet          float64
	MaxFreeGameTimes   int
	Rand               *rand.SafeRand
	BoardShape         *BoardShape
	FreeAddTimes       *FreeAddTimes
	PayOutMinimumCount *PayOutMinimumCount
}
