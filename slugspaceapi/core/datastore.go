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
	GetLotByID(id int) (models.Lot, error)
	GetLots() ([]models.Lot, error)

	GetUntrackedLots() ([]models.UntrackedLot, error)
	GetUntrackedLotByID(id int) (models.UntrackedLot, error)

	GetPermits() ([]models.Permit, error)
	GetPermitByID(id int) (models.Permit, error)

	GetPayStations() ([]models.PayStation, error)
	GetPayStationByID(id int) (models.PayStation, error)

	GetLotAvailabilities() ([]models.LotAvailability, error)
	GetLotAvailabilityByID(id int) (models.LotAvailability, error)

	GetLotDataOverTime(id int) ([]models.LotData, error)

	GetLotAverageFreespacesByDate(lotID int, checkDate time.Time, checkTime time.Time) (models.LotAverageFreespaces, error)

	GetTrackedLotFullInfoByID(lotID int) (models.TrackedLotFullInfo, error)

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
