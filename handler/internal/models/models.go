package models

import "encoding/json"

type User struct {
	id       string
	Login    string
	Password string
}

func NewUser(login, password string) *User {
	return &User{
		Login:    login,
		Password: password,
	}
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
