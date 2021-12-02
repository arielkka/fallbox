package storage

import (
	"encoding/hex"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Txt struct {
	db *sqlx.DB
}

func NewTxtStorage(db *sqlx.DB) *Txt {
	return &Txt{db: db}
}

func (t *Txt) GetTxt(userID string, txtID int) ([]byte, error) {
	text := ""
	query := fmt.Sprintf("select body from txt where id=%v and user_id='%s'", txtID, userID)
	row := t.db.QueryRow(query)
	err := row.Scan(&text)
	if err != nil {
		return nil, err
	}

	txt, err := hex.DecodeString(text)
	if err != nil {
		return nil, err
	}
	return txt, nil
}

func (t *Txt) AddTxt(userID string, txt []byte) (int, error) {
	text := hex.EncodeToString(txt)
	query := fmt.Sprintf("insert into txt(user_id,body) VALUES ('%s','%s')", userID, text)
	_, err := t.db.Exec(query)
	if err != nil {
		return -1, err
	}

	var id int
	query = fmt.Sprintf("select id from txt where user_id='%s' and body='%s'", userID, text)
	row := t.db.QueryRow(query)
	err = row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (t *Txt) DeleteTxt(userID string, txtID int) error {
	query := fmt.Sprintf("delete from txt where user_id='%s' and id=%v", userID, txtID)
	_, err := t.db.Exec(query)
	return err
}
