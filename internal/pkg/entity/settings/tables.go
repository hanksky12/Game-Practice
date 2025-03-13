package settings

import (
	"gamePractice/internal/pkg/entity/settings/line"
	"gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
)

type Tables struct {
	Wheel *reel.Table
	Pay   *payout.Table
	Line  *line.Table
}
