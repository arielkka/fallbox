package storage

import (
	"fmt"
	"github.com/arielkka/fallbox/excel/internal/entities"
	"github.com/jmoiron/sqlx"
)

type Excel struct {
	db *sqlx.DB
}

func NewExcelStorage(db *sqlx.DB) *Excel {
	return &Excel{db: db}
}

func (e *Excel) GetExcel(userID string, excelID int) ([]byte, error) {
	excel := new(entities.Excel)
	query := fmt.Sprintf("select body from excel where id=%v and user_id='%s'", excelID, userID)
	err := e.db.Get(excel, query)
	return excel.Body, err
}

func (e *Excel) AddExcel(userID string, excel []byte) (int, error) {
	query := fmt.Sprintf("insert into excel(user_id,body) VALUES ('%s','%s')", userID, string(excel))
	_, err := e.db.Exec(query)
	if err != nil {
		return -1, err
	}

	var id int
	query = fmt.Sprintf("select id from excel where user_id='%s' and body='%s'", userID, string(excel))
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
