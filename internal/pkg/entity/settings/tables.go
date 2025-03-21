package settings

import (
	"gamePractice/internal/pkg/entity/settings/line"
	"gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
)

type LineTables struct {
	Wheel *reel.Table
	Pay   *payout.Table
	Line  *line.Table
}

type WayTables struct {
	Wheel *reel.Table
	Pay   *payout.Table
}
