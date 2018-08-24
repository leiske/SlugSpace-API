package slugspace

import (
	"database/sql"
	"errors"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	. "github.com/colbyleiske/slugspace/utils"
	"testing"
)

type TestStoreAccessLayer struct{}

var tStore *Store

var newLot = models.Lot{
	Id:          0,
	Name:        "Core West",
	FullName:    "Core West Parking",
	FreeSpaces:  50,
	TotalSpaces: 100,
	LastUpdated: "2018",
}

var newData = models.LotData{
	FreeSpaces: 50,
	TimeTaken: "20:12:42",
}

func (t TestStoreAccessLayer) GetLotInfo(lotID int) (models.Lot, error) {
	if lotID == -1 {
		return models.Lot{}, errors.New("ID not found")
	}
	return newLot, nil
}

func (t TestStoreAccessLayer) GetLots() ([]models.Lot, error) {
	return []models.Lot{newLot}, nil
}

func (t TestStoreAccessLayer) GetLotDataOverTime(lotID int) ([]models.LotData, error) {
	if lotID == -1 {
		return []models.LotData{},errors.New("ID not found")
	}
	return []models.LotData{newData}, nil
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
