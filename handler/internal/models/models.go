package models

import "encoding/json"

type (
	User struct {
		ID       string `db:"id" json:"id,omitempty"`
		Login    string `db:"login" json:"login"`
		Password string `db:"password" json:"password"`
	}

	UserID struct {
		ID string `json:"id"`
	}

	Response struct {
		Body []byte `json:"body,omitempty"`
		ID   int    `json:"id,omitempty"`
	}

	IsDeleted struct {
		Flag bool `json:"is_deleted"`
	}

	Request struct {
		ID     int    `json:"id,omitempty"`
		UserID string `json:"user_id,omitempty"`
		Body   []byte `json:"body,omitempty"`
	}
)

func NewUser(login, password string) *User {
	return &User{
		Login:    login,
		Password: password,
	}
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (r *Request) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Response) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (i *IsDeleted) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
