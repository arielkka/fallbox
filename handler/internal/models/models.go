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

	ImagePath struct {
		Path string `json:"path"`
	}

	PNG struct {
		ID   string `json:"id,omitempty"`
		Body []byte `json:"body"`
	}

	PngID struct {
		ID string `json:"id"`
	}

	JPG struct {
		ID   string `json:"id,omitempty"`
		Body []byte `json:"body"`
	}

	JpgID struct {
		ID string `json:"id"`
	}

	IsDeleted struct {
		Flag bool `json:"is_deleted"`
	}

	Request struct {
		UserID string `json:"user_id,omitempty"`
		PngID  string `json:"png_id,omitempty"`
		JpgID  string `json:"jpg_id,omitempty"`
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

func (p *PNG) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (j *JPG) MarshalBinary() ([]byte, error) {
	return json.Marshal(j)
}

func (u *UserID) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (j *JpgID) MarshalBinary() ([]byte, error) {
	return json.Marshal(j)
}

func (p *PngID) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (i *ImagePath) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *IsDeleted) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
