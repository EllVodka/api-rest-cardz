package store

import (
	"log"

	"Angular/api-rest/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var noRowsSql string = "sql: no rows in result set"

type Store interface {
	GetCardz() ([]models.Cardz, error)
	GetCardzById(id int64) (models.Cardz, error)
	DeleteCardz(id int64) error
	CreateCardz(c models.Cardz) error

	Open() error
	Close() error
}

type DbStore struct {
	db *sqlx.DB
}

func (store *DbStore) Open() error {
	db, err := sqlx.Connect("mysql", "root:@(localhost:3306)/cardz")
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	store.db = db
	return nil
}

func (store *DbStore) Close() error {
	return store.db.Close()
}
