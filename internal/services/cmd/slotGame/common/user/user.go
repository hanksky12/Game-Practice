package user

import "sync"

type User struct {
	Id      int
	Balance float64
	Win     float64
	mu      sync.Mutex
}

func (u *User) AddScore(score float64, singleBet float64) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Balance += score - singleBet
	u.Win += score
}
