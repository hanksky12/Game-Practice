package session

import (
	"gamePractice/internal/services/cmd/slotGame/lineGamePractice/game/common/session"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
	Session          *session.Session
	MaxFreeGameTimes int
}

func (suite *Suite) SetupTest() {
	suite.MaxFreeGameTimes = 20
	suite.Session = &session.Session{MaxFreeGameTimes: suite.MaxFreeGameTimes}
	suite.Session.Init()
}

func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

/*
計算FreeGame 次數正確性
*/

func (suite *Suite) TestFreeGame_0() {
	log.Info("TestFreeGame_0")
	// no add
	ok, count := suite.Session.IsAnyFreeGameTimes()
	suite.Equal(false, ok)
	suite.Equal(0, count)
}

func (suite *Suite) TestFreeGame_1() {
	log.Info("TestFreeGame_1")
	suite.Session.AddFreeGameTimes(1)
	ok, count := suite.Session.IsAnyFreeGameTimes()
	suite.Equal(true, ok)
	suite.Equal(1, count)
}

func (suite *Suite) TestFreeGame_RunFinish() {
	log.Info("TestFreeGame_RunFinish")
	times := 5
	suite.Session.AddFreeGameTimes(times)
	for i := 0; i < times; i++ {
		ok, count := suite.Session.IsAnyFreeGameTimes()
		log.Info(" ok:", ok, " count:", count)
	}
	ok, count := suite.Session.IsAnyFreeGameTimes()
	suite.Equal(false, ok)
	suite.Equal(0, count)
}

func (suite *Suite) TestFreeGame_RunMaxFreeLimit() {
	log.Info("TestFreeGame_RunMaxFreeLimit")
	suite.Session.AddFreeGameTimes(suite.MaxFreeGameTimes + 10)
	for i := 0; i < suite.MaxFreeGameTimes; i++ {
		ok, count := suite.Session.IsAnyFreeGameTimes()
		log.Info(" ok:", ok, " count:", count)
	}
	ok, count := suite.Session.IsAnyFreeGameTimes()
	suite.Equal(false, ok)
	suite.Equal(0, count)
}
