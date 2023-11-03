package model

type User struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Info     UserInfo `json:"-"`
}

type UserInfo struct {
	UUID string
	Hash string
	Salt string
}
