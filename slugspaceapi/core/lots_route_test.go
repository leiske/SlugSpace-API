package slugspace_test

import (
	"encoding/json"
	"errors"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	. "github.com/colbyleiske/slugspace/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/colbyleiske/slugspace/slugspaceapi/core"
)

var trackedLot = models.Lot{
	Id:          0,
	Name:        "Core West",
	FullName:    "Core West Parking",
	FreeSpaces:  50,
	TotalSpaces: 100,
	LastUpdated: "2018",
}

func (t TestStoreAccessLayer) GetLotByID(lotID int) (models.Lot, error) {
	if lotID == -1 {
		return models.Lot{}, errors.New("ID not found")
	}
	return trackedLot, nil
}

func (t TestStoreAccessLayer) GetLots() ([]models.Lot, error) {
	return []models.Lot{trackedLot}, nil
}

//Tests

func TestGetLotByID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Lots+"/1")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	var lot models.Lot
	json.Unmarshal(res.Body.Bytes(), &lot)

	Assert(lot, trackedLot, t)
}

func TestGetLotByFakeID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Lots+"/-1")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusNotFound, t)
}

func TestGetLotByBadID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Lots+"/bad_ID")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusBadRequest, t)
}

func TestGetLots(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Lots)
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	var lot []models.Lot
	json.Unmarshal(res.Body.Bytes(), &lot)

	Assert(len(lot), 1, t)
	Assert(lot[0].FullName, "Core West Parking", t)
	Assert(lot[0], []models.Lot{trackedLot}[0], t)
}

