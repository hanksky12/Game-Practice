package scoringLine

import (
	"gamePractice/internal/pkg/entity/settings/line"
	payoutEntity "gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
	"gamePractice/internal/services/cmd/slotGame/common/payout"
	settingsCommon "gamePractice/internal/services/cmd/slotGame/common/settings"
	"gamePractice/internal/services/cmd/slotGame/practice_2/game/common/table/scoring"
	"gamePractice/internal/services/cmd/slotGame/practice_2/settings"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type Suite struct {
	suite.Suite
	LineTable   *line.Table
	PayoutTable *payoutEntity.Table
	BetBase     float64
	SingleBet   float64
}

func (suite *Suite) SetupTest() {
	basePath := suite.getBasePath()
	s := &settings.Doc{BaseDoc: &settingsCommon.Doc{}, BasePath: basePath}
	var _ bool
	_, suite.PayoutTable = s.ReadPayTable()
	suite.BetBase = 50.0
	suite.SingleBet = 100.0

}

func (suite *Suite) getBasePath() string {
	basePath, _ := os.Getwd()
	// 檢查是否包含 "internal/unitTest"
	if strings.Contains(basePath, "internal/unitTest") {
		// 往上找到專案根目錄
		for !strings.HasSuffix(basePath, "gamePractice") {
			basePath = filepath.Dir(basePath)
		}
	}
	return basePath
}

func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

/*
取出對應輪帶之後的symbol計數
*/

func (suite *Suite) TestCountAndWay_1_1() {
	log.Info("TestCountAndWay_1_1")
	reels := []*reel.Reel{
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
	}
	count, way := suite.getCountAndWay(reels)
	suite.Equal(1, count)
	suite.Equal(1, way)
}

func (suite *Suite) TestCountAndWay_2_1() {
	log.Info("TestCountAndWay_2_1")
	reels := []*reel.Reel{
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
	}
	count, way := suite.getCountAndWay(reels)
	suite.Equal(2, count)
	suite.Equal(1, way)
}

func (suite *Suite) TestCountAndWay_3_1() {
	log.Info("TestCountAndWay_3_1")
	reels := []*reel.Reel{
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
	}
	count, way := suite.getCountAndWay(reels)
	suite.Equal(3, count)
	suite.Equal(1, way)
}

func (suite *Suite) TestCountAndWay_3_2() {
	log.Info("TestCountAndWay_3_2")
	reels := []*reel.Reel{
		{
			Items: []string{"H1", "H1", "L1"},
		},
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
	}
	count, way := suite.getCountAndWay(reels)
	suite.Equal(3, count)
	suite.Equal(2, way)
}

func (suite *Suite) TestCountAndWay_3_4() {
	log.Info("TestCountAndWay_3_4")
	reels := []*reel.Reel{
		{
			Items: []string{"H1", "H1", "L1"},
		},
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"H1", "H1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
	}
	count, way := suite.getCountAndWay(reels)
	suite.Equal(3, count)
	suite.Equal(4, way)
}

func (suite *Suite) TestCountAndWay_With_W_3_2() {
	log.Info("TestCountAndWay_With_W_3_2")
	reels := []*reel.Reel{
		{
			Items: []string{"H1", "H1", "L1"},
		},
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"H1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
		{
			Items: []string{"L1", "L1", "L1"},
		},
	}
	count, way := suite.getCountAndWay(reels)
	suite.Equal(3, count)
	suite.Equal(2, way)
}

func (suite *Suite) TestCountAndWay_With_W_5_8() {
	log.Info("TestCountAndWay_With_W_5_8")
	reels := []*reel.Reel{
		{
			Items: []string{"H1", "H1", "L1"},
		},
		{
			Items: []string{"W", "W", "L1"},
		},
		{
			Items: []string{"W", "W", "L1"},
		},
		{
			Items: []string{"W", "L1", "L1"},
		},
		{
			Items: []string{"W", "L1", "L1"},
		},
	}
	count, way := suite.getCountAndWay(reels)
	suite.Equal(5, count)
	suite.Equal(8, way)
}

func (suite *Suite) TestCountAndWay_With_W_5_24() {
	log.Info("TestCountAndWay_With_W_5_24")
	reels := []*reel.Reel{
		{
			Items: []string{"H1", "H1", "L1"},
		},
		{
			Items: []string{"W", "H1", "W"},
		},
		{
			Items: []string{"W", "L1", "W"},
		},
		{
			Items: []string{"W", "L1", "L1"},
		},
		{
			Items: []string{"W", "L1", "H1"},
		},
	}
	count, way := suite.getCountAndWay(reels)
	suite.Equal(5, count)
	suite.Equal(24, way) //????????
}

func (suite *Suite) getCountAndWay(reels []*reel.Reel) (int, int) {
	symbol := "H1"
	scoringWay := scoring.Way{}
	count, way := scoringWay.GetCountAndWay(symbol, reels)
	log.Info("count: ", count, " way: ", way)
	return count, int(way)
}

////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////
/*
計分的正確性
*/

func (suite *Suite) TestQueryScore_H1() {
	log.Info("TestQueryScore_H1")
	multiples := []float64{0, 0, 60, 120, 500}
	symbol := "H1"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_H2() {
	log.Info("TestQueryScore_H2")
	multiples := []float64{0, 0, 30, 60, 300}
	symbol := "H2"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_H3() {
	log.Info("TestQueryScore_H3")
	multiples := []float64{0, 0, 20, 50, 250}
	symbol := "H3"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_H4() {
	log.Info("TestQueryScore_H4")
	multiples := []float64{0, 0, 15, 30, 150}
	symbol := "H4"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_H5() {
	log.Info("TestQueryScore_H5")
	multiples := []float64{0, 0, 5, 10, 75}
	symbol := "H5"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_L1() {
	log.Info("TestQueryScore_L1")
	multiples := []float64{0, 0, 5, 10, 50}
	symbol := "L1"
	suite.QueryScore(symbol, multiples)
}
func (suite *Suite) TestQueryScore_L2() {
	log.Info("TestQueryScore_L2")
	multiples := []float64{0, 0, 5, 10, 50}
	symbol := "L2"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_L3() {
	log.Info("TestQueryScore_L3")
	multiples := []float64{0, 0, 5, 10, 50}
	symbol := "L3"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_L4() {
	log.Info("TestQueryScore_L4")
	multiples := []float64{0, 0, 5, 10, 50}
	symbol := "L4"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_L5() {
	log.Info("TestQueryScore_L5")
	multiples := []float64{0, 0, 5, 10, 50}
	symbol := "L5"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_L6() {
	log.Info("TestQueryScore_L6")
	multiples := []float64{0, 0, 5, 10, 50}
	symbol := "L6"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_SC() {
	log.Info("TestQueryScore_SC")
	multiples := []float64{0, 0, 0, 0, 0}
	symbol := "SC"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_W() {
	log.Info("TestQueryScore_W")
	multiples := []float64{0, 0, 0, 0, 0}
	symbol := "W"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) QueryScore(symbol string, multiples []float64) {
	pay := payout.Payout{BetBase: 50, SingleBet: suite.SingleBet, PayoutTable: suite.PayoutTable}
	way := 1.0
	for i := 0; i < 1; i++ {
		score := pay.QueryWayScore(5-i, symbol, way)
		log.Info("score: ", score)
		suite.Equal(suite.scoreRule(multiples[5-1-i], way), score)
	}
}

func (suite *Suite) scoreRule(multiples, way float64) float64 {
	return suite.SingleBet * multiples / suite.BetBase * way
}
