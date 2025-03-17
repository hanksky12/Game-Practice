package user

type User struct {
	Id      int
	Balance float64
	Win     float64
}

func (u *User) AddScore(score float64) {
	u.Balance += score
}

func (u *User) AddWin(score float64) {
	u.Win += score
}
