package memory

import "time"

type Users map[string]User

func NewUsers() Users {
	return make(map[string]User)
}

type User struct {
	Hash        string
	Salt        string
	Orders      map[string]Order
	Withdrawals map[string]Withdrawals
	Balance     Balance
}

func NewUser(hash, salt string) User {
	return User{
		Hash:        hash,
		Salt:        salt,
		Orders:      make(map[string]Order),
		Withdrawals: make(map[string]Withdrawals),
	}
}

type Order struct {
	Status     string
	Accrual    *int
	UploadedAt time.Time
}

type Withdrawals struct {
	Sum         int
	ProcessedAt time.Time
}

type Balance struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}
