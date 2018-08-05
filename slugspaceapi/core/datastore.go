package slugspace

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	"log"
)

type Store struct {
	db  *sql.DB
	dal DataAccessLayer
}

type DataAccessLayer interface {
	GetLotInfo(lotID int) (models.Lot, error)
	GetLots() ([]models.Lot, error)
}

func NewStore(db *sql.DB, dal DataAccessLayer) *Store {
	return &Store{db: db, dal: dal}
}

func (s *Store) DB() *sql.DB {
	return s.db
}

func (s *Store) DAL() DataAccessLayer {
	return s.dal
}

func (s *Store) CloseDB() {
	s.db.Close()
	log.Println("DB Closed")
}
