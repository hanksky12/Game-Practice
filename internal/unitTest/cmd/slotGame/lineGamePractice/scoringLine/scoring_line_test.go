package scoringLine

import (
	"gamePractice/internal/pkg/entity/settings/line"
	"gamePractice/internal/pkg/entity/settings/payout"
	"gamePractice/internal/pkg/entity/settings/reel"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/table/scoring"
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/settings"
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
	PayoutTable *payout.Table
	BetBase     float64
	SingleBet   float64
}

func (suite *Suite) SetupTest() {
	basePath := suite.getBasePath()
	s := &settings.Doc{basePath}
	var _ bool
	_, suite.LineTable = s.ReadLineTable()
	_, suite.PayoutTable = s.ReadPayTable()
	suite.BetBase = 10.0
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

func (suite *Suite) TestCountAndSymbol_Basic_3() {
	log.Info("TestCountAndSymbol_Basic_3")
	reelItems := []string{"H1", "H1", "H1", "H4", "L1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(3, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_Basic_4() {
	log.Info("TestCountAndSymbol_Basic_4")
	reelItems := []string{"H1", "H1", "H1", "H1", "L1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(4, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_Basic_5() {
	log.Info("TestCountAndSymbol_Basic_5")
	reelItems := []string{"H1", "H1", "H1", "H1", "H1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(5, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_Fail() {
	log.Info("TestCountAndSymbol_Fail")
	reelItems := []string{"H1", "L1", "H1", "H1", "H1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(1, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_Fail_With_W() {
	log.Info("TestCountAndSymbol_Fail_With_W")
	reelItems := []string{"H1", "L1", "W", "H1", "H1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(1, count)
	suite.Equal("H1", symbol)
}

// W百搭
func (suite *Suite) TestCountAndSymbol_With_In_Last_1W() {
	log.Info("TestCountAndSymbol_With_In_Last_1W")
	reelItems := []string{"H1", "H1", "W", "H4", "L1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(3, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_In_Last_2W() {
	log.Info("TestCountAndSymbol_With_In_Last_2W")
	reelItems := []string{"H1", "W", "W", "H4", "L1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(3, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_In_Last_3W() {
	log.Info("TestCountAndSymbol_With_In_Last_3W")
	reelItems := []string{"H1", "W", "W", "W", "L1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(4, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_In_Last_4W() {
	log.Info("TestCountAndSymbol_With_In_Last_4W")
	reelItems := []string{"H1", "W", "W", "W", "W"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(5, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_In_Middle_1W() {
	log.Info("TestCountAndSymbol_With_In_Middle_1W")
	reelItems := []string{"H1", "H1", "W", "H1", "L1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(4, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_In_Middle_2W() {
	log.Info("TestCountAndSymbol_With_In_Middle_2W")
	reelItems := []string{"H1", "W", "W", "H1", "L1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(4, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_In_Middle_3W() {
	log.Info("TestCountAndSymbol_With_In_Middle_3W")
	reelItems := []string{"H1", "W", "W", "W", "H1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(5, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_In_First_1W() {
	log.Info("TestCountAndSymbol_With_In_First_1W")
	reelItems := []string{"W", "H1", "H2", "H1", "L1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(2, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_In_First_2W() {
	log.Info("TestCountAndSymbol_With_In_First_2W")
	reelItems := []string{"W", "W", "H1", "H2", "W"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(3, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_In_First_3W() {
	log.Info("TestCountAndSymbol_With_In_First_3W")
	reelItems := []string{"W", "W", "W", "H1", "L1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(4, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_In_First_4W() {
	log.Info("TestCountAndSymbol_With_In_First_4W")
	reelItems := []string{"W", "W", "W", "W", "H1"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(5, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_W_Jump() {
	log.Info("TestCountAndSymbol_With_W_Jump")
	reelItems := []string{"W", "W", "W", "H1", "W"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(5, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_W_Jump1() {
	log.Info("TestCountAndSymbol_With_W_Jump1")
	reelItems := []string{"W", "H1", "W", "H2", "W"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(3, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_W_Jump2() {
	log.Info("TestCountAndSymbol_With_W_Jump2")
	reelItems := []string{"W", "H1", "H1", "W", "W"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(5, count)
	suite.Equal("H1", symbol)
}

func (suite *Suite) TestCountAndSymbol_With_5W() {
	log.Info("TestCountAndSymbol_With_5W")
	reelItems := []string{"W", "W", "W", "W", "W"}
	count, symbol := suite.getCountAndSymbol(reelItems)
	suite.Equal(5, count)
	suite.Equal("W", symbol)
}

func (suite *Suite) getCountAndSymbol(reelItems []string) (int, string) {
	scoringLine := scoring.Line{}
	lineRow := &line.Row{
		Case: []int{0, 0, 0, 0, 0},
	}
	reels := []*reel.Reel{
		{
			Items: []string{reelItems[0]},
		},
		{
			Items: []string{reelItems[1]},
		},
		{
			Items: []string{reelItems[2]},
		},
		{
			Items: []string{reelItems[3]},
		},
		{
			Items: []string{reelItems[4]},
		},
	}
	count, symbol := scoringLine.GetCountAndSymbol(lineRow, reels)
	log.Info("count: ", count, " symbol: ", symbol)
	return count, symbol
}

// ////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////
/*
計分的正確性
*/

func (suite *Suite) TestQueryScore_H1() {
	log.Info("TestQueryScore_H1")
	multiples := []float64{0, 10, 100, 1000, 5000}
	symbol := "H1"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_H2() {
	log.Info("TestQueryScore_H2")
	multiples := []float64{0, 5, 40, 400, 2000}
	symbol := "H2"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_H3() {
	log.Info("TestQueryScore_H3")
	multiples := []float64{0, 5, 30, 100, 750}
	symbol := "H3"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_H4() {
	log.Info("TestQueryScore_H4")
	multiples := []float64{0, 5, 30, 100, 750}
	symbol := "H4"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_L1() {
	log.Info("TestQueryScore_L1")
	multiples := []float64{0, 0, 5, 40, 150}
	symbol := "L1"
	suite.QueryScore(symbol, multiples)
}
func (suite *Suite) TestQueryScore_L2() {
	log.Info("TestQueryScore_L2")
	multiples := []float64{0, 0, 5, 40, 150}
	symbol := "L2"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_L3() {
	log.Info("TestQueryScore_L3")
	multiples := []float64{0, 0, 5, 25, 100}
	symbol := "L3"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_L4() {
	log.Info("TestQueryScore_L4")
	multiples := []float64{0, 0, 5, 25, 100}
	symbol := "L4"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_L5() {
	log.Info("TestQueryScore_L5")
	multiples := []float64{0, 0, 5, 25, 100}
	symbol := "L5"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_SC() {
	log.Info("TestQueryScore_SC")
	multiples := []float64{0, 0, 2, 20, 200}
	symbol := "SC"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) TestQueryScore_W() {
	log.Info("TestQueryScore_W")
	multiples := []float64{0, 0, 2, 20, 200}
	symbol := "W"
	suite.QueryScore(symbol, multiples)
}

func (suite *Suite) QueryScore(symbol string, multiples []float64) {
	scoringLine := &scoring.Line{}
	for i := 0; i < 1; i++ {
		score := scoringLine.QueryScore(5-i, symbol, suite.SingleBet, suite.LineTable, suite.PayoutTable)
		log.Info("score: ", score)
		suite.Equal(suite.scoreRule(multiples[5-1-i]), score)
	}
}

func (suite *Suite) scoreRule(multiples float64) float64 {
	return suite.SingleBet * multiples / suite.BetBase
}

// ////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////

/*
line 表 從盤面能正確的取出位置
*/

func (suite *Suite) TestGetReelItems_0() {
	log.Info("TestGetReelItems_0")
	lineRow := suite.LineTable.Rows[0]
	except := []string{"H2", "H2", "H2", "H2", "H2"}
	suite.getItems(lineRow, except)
}

func (suite *Suite) TestGetReelItems_1() {
	log.Info("TestGetReelItems_1")
	lineRow := suite.LineTable.Rows[1]
	except := []string{"H1", "H2", "H3", "H4", "L1"}
	suite.getItems(lineRow, except)
}

func (suite *Suite) TestGetReelItems_2() {
	log.Info("TestGetReelItems_2")
	lineRow := suite.LineTable.Rows[2]
	except := []string{"H3", "H3", "H3", "H3", "H3"}
	suite.getItems(lineRow, except)
}

func (suite *Suite) TestGetReelItems_3() {
	log.Info("TestGetReelItems_3")
	lineRow := suite.LineTable.Rows[3]
	except := []string{"H1", "H2", "H3", "H2", "L1"}
	suite.getItems(lineRow, except)
}

func (suite *Suite) TestGetReelItems_4() {
	log.Info("TestGetReelItems_4")
	lineRow := suite.LineTable.Rows[4]
	except := []string{"H3", "H2", "H3", "H2", "H3"}
	suite.getItems(lineRow, except)
}

func (suite *Suite) TestGetReelItems_5() {
	log.Info("TestGetReelItems_5")
	lineRow := suite.LineTable.Rows[5]
	except := []string{"H2", "H2", "H3", "H4", "H2"}
	suite.getItems(lineRow, except)
}

func (suite *Suite) TestGetReelItems_6() {
	log.Info("TestGetReelItems_6")
	lineRow := suite.LineTable.Rows[6]
	except := []string{"H2", "H3", "H3", "H3", "H2"}
	suite.getItems(lineRow, except)
}

func (suite *Suite) TestGetReelItems_7() {
	log.Info("TestGetReelItems_7")
	lineRow := suite.LineTable.Rows[7]
	except := []string{"H1", "H2", "H2", "H3", "H3"}
	suite.getItems(lineRow, except)
}

func (suite *Suite) TestGetReelItems_8() {
	log.Info("TestGetReelItems_8")
	lineRow := suite.LineTable.Rows[8]
	except := []string{"H3", "H3", "H2", "H4", "L1"}
	suite.getItems(lineRow, except)
}

func (suite *Suite) TestGetReelItems_9() {
	log.Info("TestGetReelItems_9")
	lineRow := suite.LineTable.Rows[9]
	except := []string{"H2", "H3", "H2", "H4", "H2"}
	suite.getItems(lineRow, except)
}

func (suite *Suite) getItems(lineRow *line.Row, except []string) {
	reels := []*reel.Reel{
		{
			Items: []string{"H1", "H2", "H3", "H4", "L1"},
		},
		{
			Items: []string{"H2", "H2", "H3", "H4", "L1"},
		},
		{
			Items: []string{"H3", "H2", "H3", "H4", "L1"},
		},
		{
			Items: []string{"H4", "H2", "H3", "H4", "L1"},
		},
		{
			Items: []string{"L1", "H2", "H3", "H4", "L1"},
		},
	}
	scoringLine := &scoring.Line{}
	items := scoringLine.GetReelItems(lineRow, reels)
	log.Info("items: ", items)
	suite.Equal(except, items)
}
