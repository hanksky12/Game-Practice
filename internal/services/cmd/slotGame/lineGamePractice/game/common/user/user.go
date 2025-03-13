package user

type User struct {
	Id      int
	Balance float64
}

func (u *User) AddScore(score float64) {
	u.Balance += score
}
