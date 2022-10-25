package domain

type Auth struct {
	UserId       UserId
	Login        Login    `json:"login"`
	Password     Password `json:"password"`
	PasswordHash PasswordHash
}
