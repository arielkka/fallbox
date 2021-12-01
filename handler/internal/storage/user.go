package storage

import (
	"fmt"
	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u *User) GetUser(login, password string) (string, error) {
	user := new(models.User)
	query := fmt.Sprintf("select id, login, password from user where login='%s' and password='%s'", login, password)
	err := u.db.Get(user, query)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (u *User) CreateUser(login, password, id string) error {
	query := fmt.Sprintf("INSERT INTO positions(login, password, id) VALUES ('%s','%s','%s')", login, password, id)
	_, err := u.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
