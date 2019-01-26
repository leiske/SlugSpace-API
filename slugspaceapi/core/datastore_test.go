package slugspace

import (
	"database/sql"
	"errors"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/database"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
	. "github.com/colbyleiske/slugspace/utils"
	"testing"
	"time"
	"net/http"
	)

type TestStoreAccessLayer struct{}

var tStore *Store

var trackedLot = models.Lot{
	Id:          0,
	Name:        "Core West",
	FullName:    "Core West Parking",
	FreeSpaces:  50,
	TotalSpaces: 100,
	LastUpdated: "2018",
}

var trackedLotData = models.LotData{
	FreeSpaces: 50,
	TimeTaken:  "20:12:42",
}

var untrackedLot = models.UntrackedLot{
	Id: 0,
	LotName: "Test Untracked Lot Name",
	LotNumber: 1,
	/* omitting other fields */
}

func (t TestStoreAccessLayer) GetLotInfo(lotID int) (models.Lot, error) {
	if lotID == -1 {
		return models.Lot{}, errors.New("ID not found")
	}
	return trackedLot, nil
}

func (t TestStoreAccessLayer) GetLots() ([]models.Lot, error) {
	return []models.Lot{trackedLot}, nil
}

func (t TestStoreAccessLayer) GetUntrackedLotInfo(lotID int) (models.UntrackedLot, error) {
	if lotID == -1 {
		return models.UntrackedLot{}, errors.New("ID not found")
	}
	return untrackedLot, nil
}

func (t TestStoreAccessLayer) GetUntrackedLots() ([]models.UntrackedLot, error) {
	return []models.UntrackedLot{untrackedLot}, nil
}

func (t TestStoreAccessLayer) GetLotDataOverTime(lotID int) ([]models.LotData, error) {
	if lotID == -1 {
		return []models.LotData{}, errors.New("ID not found")
	}
	return []models.LotData{trackedLotData}, nil
}

func (t TestStoreAccessLayer) CreateJWT(payload *database.JWTPayload) (string, error) {
	return "", nil //temp
}

func (t TestStoreAccessLayer) GetTokenSecret(guid interface{}) (interface{}, bool, error) {
	return []byte(constants.TestSecret), true, nil
}

func (t TestStoreAccessLayer) GetLotAverageFreespacesByDate(lotID int, checkDate time.Time, checkTime time.Time) (models.LotAverageFreespaces, error) {
	return models.LotAverageFreespaces{}, nil //temp
}

func CreateAuthenticatedRequest(endpoint string) (*http.Request, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization",constants.TestToken)
	return req, nil
}

func TestNewStore(t *testing.T) {
	db, _ := sql.Open("", "")
	tal := TestStoreAccessLayer{}
	s := NewStore(db, tal)

	AssertNonNil(s.db, t)
	AssertNonNil(s.dal, t)

	Assert(s.db, db, t)
	Assert(s.dal, tal, t)
}
