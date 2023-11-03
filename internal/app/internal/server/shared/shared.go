package shared

type User struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Info     UserInfo `json:"-"`
}

type UserInfo struct {
	Hash string
	Salt string
}
