package settings

import (
	"gamePractice/internal/pkg/const/game"
	"gamePractice/internal/pkg/const/symbol"
	"gamePractice/internal/pkg/entity/settings/line"
	"gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
)

type MockDoc struct {
}

func (d *MockDoc) ReadWheelTable() (bool, *reel.Table, *reel.Table) {
	mgT := &reel.Table{
		Name: game.Main,
		Reels: []*reel.Reel{
			{Items: []string{"Sc", "W"}},
			{Items: []string{"L5", "Sc", "W"}},
			{Items: []string{"L4", "Sc", "W"}},
			{Items: []string{"Sc", "W"}},
			{Items: []string{"H1", "H2", "H3", "H4", "L1", "L2", "L3", "L4", "L5", "Sc", "W"}},
		},
	}
	fgT := &reel.Table{
		Name: game.Free,
		Reels: []*reel.Reel{
			{Items: []string{"Sc", "W"}},
			{Items: []string{"L5", "Sc", "W"}},
			{Items: []string{"L4", "Sc", "W"}},
			{Items: []string{"Sc", "W"}},
			{Items: []string{"H1", "H2", "H3", "H4", "L1", "L2", "L3", "L4", "L5", "Sc", "W"}},
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

			{Symbol: symbol.L1, Pays: []float64{0, 0, 10, 20, 30.0}},
			{Symbol: symbol.L2, Pays: []float64{0, 0, 10, 30, 40.0}},
			{Symbol: symbol.L3, Pays: []float64{0, 0, 10, 30, 40.0}},
			{Symbol: symbol.L4, Pays: []float64{0, 0, 10, 30, 40.0}},
			{Symbol: symbol.L5, Pays: []float64{0, 0, 10, 30, 40.0}},
			{Symbol: symbol.Sc, Pays: []float64{0, 0, 10, 30, 40.0}},
			{Symbol: symbol.W, Pays: []float64{0, 10, 20, 30, 70.0}},
		},
	}
	return true, table
}
func (d *MockDoc) ReadLineTable() (bool, *line.Table) {
	table := &line.Table{
		Name: "lineTable",
		Rows: []*line.Row{
			{Case: []int{0, 0, 0, 0, 0}},
			{Case: []int{1, 1, 1, 1, 1}},
			{Case: []int{2, 2, 2, 2, 2}},
			{Case: []int{0, 1, 2, 1, 0}},
			{Case: []int{2, 1, 0, 1, 2}},
			{Case: []int{0, 0, 1, 0, 0}},
			{Case: []int{2, 2, 1, 2, 2}},
			{Case: []int{1, 0, 0, 0, 1}},
			{Case: []int{1, 2, 2, 2, 1}},
			{Case: []int{0, 1, 1, 1, 0}},
		},
	}
	return true, table
}
