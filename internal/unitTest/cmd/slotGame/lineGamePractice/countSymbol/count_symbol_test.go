package scoring_line

import (
	"gamePractice/internal/pkg/const/symbol"
	"gamePractice/internal/pkg/entity/settings/reel"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/table/count"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
}

func (suite *Suite) SetupTest() {
}

func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))

}

/* 計算Sc的數量
W 不取代Sc*/
// ///////////////////////
// /0~5 SC
func (suite *Suite) TestCountSymbol_Basic_0() {
	log.Info("TestCountSymbol_Basic_0")
	reels := []*reel.Reel{
		{Items: []string{"L5", "L5", "L5"}},
		{Items: []string{"L5", "L5", "L5"}},
		{Items: []string{"L4", "L5", "L5"}},
		{Items: []string{"L4", "L4", "L4"}},
		{Items: []string{"H1", "H2", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(0, count)
}

func (suite *Suite) TestCountSymbol_Basic_1() {
	log.Info("TestCountSymbol_Basic_1")
	reels := []*reel.Reel{
		{Items: []string{"SC", "L5", "L5"}},
		{Items: []string{"L5", "L5", "L5"}},
		{Items: []string{"SC", "L5", "L5"}},
		{Items: []string{"L4", "L4", "L4"}},
		{Items: []string{"H1", "H2", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(2, count)
}

func (suite *Suite) TestCountSymbol_Basic_2() {
	log.Info("TestCountSymbol_Basic_2")
	reels := []*reel.Reel{
		{Items: []string{"SC", "L5", "L5"}},
		{Items: []string{"SC", "L5", "L5"}},
		{Items: []string{"L4", "L5", "L5"}},
		{Items: []string{"L4", "SC", "L4"}},
		{Items: []string{"H1", "H2", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(3, count)
}

func (suite *Suite) TestCountSymbol_Basic_3() {
	log.Info("TestCountSymbol_Basic_3")
	reels := []*reel.Reel{
		{Items: []string{"SC", "L4", "L4"}},
		{Items: []string{"L4", "L4", "SC"}},
		{Items: []string{"L4", "L4", "SC"}},
		{Items: []string{"L4", "L4", "L4"}},
		{Items: []string{"H1", "SC", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(4, count)
}

func (suite *Suite) TestCountSymbol_Basic_4() {
	log.Info("TestCountSymbol_Basic_4")
	reels := []*reel.Reel{
		{Items: []string{"L5", "L5", "SC"}},
		{Items: []string{"L5", "SC", "L5"}},
		{Items: []string{"L4", "SC", "L5"}},
		{Items: []string{"L4", "L4", "SC"}},
		{Items: []string{"SC", "H2", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(5, count)
}

func (suite *Suite) TestCountSymbol_Basic_5() {
	log.Info("TestCountSymbol_Basic_5")
	reels := []*reel.Reel{
		{Items: []string{"SC", "L5", "L5"}},
		{Items: []string{"L5", "L5", "L5"}},
		{Items: []string{"L4", "SC", "L5"}},
		{Items: []string{"L4", "L4", "L5"}},
		{Items: []string{"SC", "L5", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(3, count)
}

// /////////////////////////////////////////////////
// 頭Sc尾Ｗ
func (suite *Suite) TestCountSymbol_With_Sc_Last_W() {
	log.Info("TestCountSymbol_With_Sc_Last_W")
	reels := []*reel.Reel{
		{Items: []string{"SC", "W", "W"}},
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L4", "L4", "L5"}},
		{Items: []string{"L4", "L4", "L4"}},
		{Items: []string{"SC", "L4", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(2, count)
}

func (suite *Suite) TestCountSymbol_With_Sc_Last_2W() {
	log.Info("TestCountSymbol_With_Sc_Last_2W")
	reels := []*reel.Reel{
		{Items: []string{"SC", "L5", "L5"}},
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L4", "W", "L5"}},
		{Items: []string{"L4", "L4", "L4"}},
		{Items: []string{"SC", "L5", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(2, count)
}

// /////////////////////////////////////////////////
// 頭Ｗ尾Sc
func (suite *Suite) TestCountSymbol_With_First_W_Last_Sc() {
	log.Info("TestCountSymbol_With_First_W_Last_Sc")
	reels := []*reel.Reel{
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L5", "SC", "L5"}},
		{Items: []string{"L4", "L5", "L5"}},
		{Items: []string{"L4", "L4", "L4"}},
		{Items: []string{"SC", "L5", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(2, count)
}

func (suite *Suite) TestCountSymbol_With_First_2W_Last_Sc() {
	log.Info("TestCountSymbol_With_First_2W_Last_Sc")
	reels := []*reel.Reel{
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L4", "SC", "L5"}},
		{Items: []string{"L4", "L4", "L4"}},
		{Items: []string{"L4", "SC", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(2, count)
}

func (suite *Suite) TestCountSymbol_With_First_3W_Last_Sc() {
	log.Info("TestCountSymbol_With_First_3W_Last_Sc")
	reels := []*reel.Reel{
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L4", "W", "L5"}},
		{Items: []string{"L4", "SC", "L4"}},
		{Items: []string{"L4", "L4", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(1, count)
}

// /////////////////////////////////////////////////
// 3Ｗ 0Sc
func (suite *Suite) TestCountSymbol_With_3W_No_Sc() {
	log.Info("TestCountSymbol_With_First_3W_Last_Sc")
	reels := []*reel.Reel{
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L4", "W", "L5"}},
		{Items: []string{"L4", "L4", "L4"}},
		{Items: []string{"L4", "L4", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(0, count)
}

// /////////////////////////////////////////////////
// ScＷSc
func (suite *Suite) TestCountSymbol_With_W_Between_Sc() {
	log.Info("TestCountSymbol_With_W_Between_Sc")
	reels := []*reel.Reel{
		{Items: []string{"L5", "SC", "L5"}},
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L4", "SC", "L5"}},
		{Items: []string{"L4", "L4", "L4"}},
		{Items: []string{"L4", "L4", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(2, count)
}

// /////////////////////////////////////////////////
// WScＷ
func (suite *Suite) TestCountSymbol_With_Sc_Between_W() {
	log.Info("TestCountSymbol_With_Sc_Between_W")
	reels := []*reel.Reel{
		{Items: []string{"L5", "W", "L5"}},
		{Items: []string{"L5", "SC", "L5"}},
		{Items: []string{"L4", "W", "L5"}},
		{Items: []string{"L4", "L4", "L4"}},
		{Items: []string{"L4", "L4", "L4"}},
	}
	count := count.Symbol(reels, symbol.SC)
	suite.Equal(1, count)
}
