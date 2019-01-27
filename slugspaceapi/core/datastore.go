package slugspace

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/colbyleiske/slugspace/slugspaceapi/core/database"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	"log"
	"time"
)

type Store struct {
	db  *sql.DB
	dal DataAccessLayer
}

type DataAccessLayer interface {
	GetLotByID(lotID int) (models.Lot, error)
	GetLots() ([]models.Lot, error)

	GetUntrackedLots() ([]models.UntrackedLot, error)
	GetUntrackedLotByID(lotID int) (models.UntrackedLot, error)

	GetPermits() ([]models.Permit, error)
	GetPermitByID(permitID int) (models.Permit, error)


	GetLotDataOverTime(lotID int) ([]models.LotData, error)

	GetLotAverageFreespacesByDate(lotID int, checkDate time.Time, checkTime time.Time) (models.LotAverageFreespaces, error)

	CreateJWT(payload *database.JWTPayload) (string, error)
	GetTokenSecret(guid interface{}) (interface{}, bool, error)

	Log(category string, severity int, messages ...string)
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
