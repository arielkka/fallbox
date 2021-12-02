package storage

import (
	"encoding/hex"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Excel struct {
	db *sqlx.DB
}

func NewExcelStorage(db *sqlx.DB) *Excel {
	return &Excel{db: db}
}

func (e *Excel) GetExcel(userID string, excelID int) ([]byte, error) {
	excel := ""
	query := fmt.Sprintf("select body from excel where id=%v and user_id='%s'", excelID, userID)
	row := e.db.QueryRow(query)
	err := row.Scan(&excel)
	if err != nil {
		return nil, err
	}
	exc, err := hex.DecodeString(excel)
	if err != nil {
		return nil, err
	}

	return exc, nil
}

func (e *Excel) AddExcel(userID string, excel []byte) (int, error) {
	exc := hex.EncodeToString(excel)
	query := fmt.Sprintf("insert into excel(user_id,body) VALUES ('%s','%s')", userID, exc)
	_, err := e.db.Exec(query)
	if err != nil {
		return -1, err
	}

	var id int
	query = fmt.Sprintf("select id from excel where user_id='%s' and body='%s'", userID, exc)
	row := e.db.QueryRow(query)
	err = row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (e *Excel) DeleteExcel(userID string, excelID int) error {
	query := fmt.Sprintf("delete from excel where user_id='%s' and id=%v", userID, excelID)
	_, err := e.db.Exec(query)
	return err
}
