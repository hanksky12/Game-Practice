package session

type Session struct {
	MaxFreeGameTimes       int
	freeGameTimes          int
	freeGameProcessedTimes int
}

func (s *Session) Init() {
	s.freeGameTimes = 0
	s.freeGameProcessedTimes = 0
}

func (s *Session) AddFreeGameTimes(times int) {
	s.freeGameTimes += times
}

func (s *Session) IsAnyFreeGameTimes() (bool, int) {
	if s.freeGameTimes == 0 {
		//log.Info("沒有任何免費場次~~~")
		return false, 0
	}
	if s.freeGameProcessedTimes >= s.freeGameTimes {
		//log.Info("已經處理完所有免費場次~~~")
		//已經處理完所有場次(因為先加1再判斷，所以有等於)
		return false, 0
	}
	if s.freeGameProcessedTimes >= s.MaxFreeGameTimes {
		//log.Info("已經到達免費場次上限~~~")
		//已經處理完max場次(因為先加1再判斷，所以有等於)
		return false, 0
	}
	s.freeGameProcessedTimes += 1
	return true, s.freeGameProcessedTimes // 第幾場
}
