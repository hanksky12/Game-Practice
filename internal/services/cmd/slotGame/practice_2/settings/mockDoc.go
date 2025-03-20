package settings

import (
	"gamePractice/internal/pkg/const/game"
	"gamePractice/internal/pkg/const/symbol"
	"gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
)

type MockDoc struct {
}

func (d *MockDoc) ReadWheelTable() (bool, *reel.Table, *reel.Table) {
	mgT := &reel.Table{
		Name: game.Main,
		Reels: []*reel.Reel{
			{Items: []string{"L1", "W", "H2", "SC"}},
			{Items: []string{"L1", "SC", "H2", "W"}},
			{Items: []string{"L5", "SC", "W"}},
			{Items: []string{"SC", "L5"}},
			{Items: []string{"H1", "H2", "H3", "H4", "L1", "L2", "L3", "L4", "L5", "SC", "W"}},

			//{Items: []string{"L5", "W", "L5"}},
			//{Items: []string{"L5", "H2", "W"}},
			//{Items: []string{"L5", "SC", "W"}},
			//{Items: []string{"SC", "W"}},
			//{Items: []string{"H1", "H2", "H3", "H4", "L1", "L2", "L3", "L4", "L5", "SC", "W"}},
		},
	}
	fgT := &reel.Table{
		Name: game.Free,
		Reels: []*reel.Reel{
			{Items: []string{"L4", "W", "L5"}},
			{Items: []string{"L5", "SC", "W"}},
			{Items: []string{"L4", "SC", "W"}},
			{Items: []string{"SC", "W"}},
			{Items: []string{"H1", "H2", "H3", "H4", "L1", "L2", "L3", "L4", "L5", "SC", "W"}},
		},
	}

	return true, mgT, fgT
}

func (d *MockDoc) ReadPayTable() (bool, *payout.Table) {
	table := &payout.Table{
		Name: "payTable",
		Rows: []*payout.Row{
			{Symbol: symbol.H1, Pays: []float64{0, 0, 50.0, 20.0, 100.0}},
			{Symbol: symbol.H2, Pays: []float64{0, 0, 60, 70, 80.0}},
			{Symbol: symbol.H3, Pays: []float64{0, 0, 20, 30, 40.0}},
			{Symbol: symbol.H4, Pays: []float64{0, 0, 10, 30, 40.0}},
			{Symbol: symbol.H5, Pays: []float64{0, 0, 10, 30, 40.0}},

			{Symbol: symbol.L1, Pays: []float64{0, 0, 10, 20, 30.0}},
			{Symbol: symbol.L2, Pays: []float64{0, 0, 10, 30, 40.0}},
			{Symbol: symbol.L3, Pays: []float64{0, 0, 10, 30, 40.0}},
			{Symbol: symbol.L4, Pays: []float64{0, 0, 10, 30, 40.0}},
			{Symbol: symbol.L5, Pays: []float64{0, 0, 10, 30, 40.0}},
			{Symbol: symbol.L6, Pays: []float64{0, 0, 10, 30, 40.0}},

			{Symbol: symbol.SC, Pays: []float64{0, 0, 0, 0, 0}},
			{Symbol: symbol.W, Pays: []float64{0, 0, 0, 0, 0}},
		},
	}
	return true, table
}
