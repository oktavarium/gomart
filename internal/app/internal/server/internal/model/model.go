package model

import "time"

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Order struct {
	Order      string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    *int      `json:"accrual,omitempty"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type Withdrawals struct {
	Order       string    `json:"number"`
	Sum         int       `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

type Withdrawal struct {
	Order string `json:"number"`
	Sum   int    `json:"sum"`
}

type Balance struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}
