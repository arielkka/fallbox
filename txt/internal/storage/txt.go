package storage

import (
	"github.com/jmoiron/sqlx"
)

type TXT struct {
	db *sqlx.DB
}

func NewTxtStorage(db *sqlx.DB) *TXT {
	return &TXT{db: db}
}

func (p *TXT) GetExcel(userID string, txtID int) ([]byte, error) {
	panic("implement me")

}

func (p *TXT) AddExcel(userID string, txt []byte) (int, error) {
	panic("implement me")

}

func (p *TXT) DeleteExcel(userID string, txtID int) error {

	panic("implement me")
}
