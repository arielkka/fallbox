package mysql

import (
	"fmt"
	"sync"

	"github.com/arielkka/fallbox/handler/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB   *sqlx.DB
	once sync.Once
	err  error
)

func NewMySQL(cfg *config.Config) (*sqlx.DB, error) {
	once.Do(func() {
		DB, err = sqlx.Open("mysql", fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Name,
		))
	})
	if err != nil {
		return nil, err
	}
	err = DB.Ping()
	if err != nil {
		return nil, err
	}
	return DB, nil
}
