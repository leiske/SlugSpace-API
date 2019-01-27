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

var untrackedLot = models.UntrackedLot{
	Id:        0,
	Name:      "Test Untracked Lot Name",
	LotNumber: 1,
}

func (t TestStoreAccessLayer) GetUntrackedLotByID(lotID int) (models.UntrackedLot, error) {
	if lotID == -1 {
		return models.UntrackedLot{}, errors.New("ID not found")
	}
	return untrackedLot, nil
}

func (t TestStoreAccessLayer) GetUntrackedLots() ([]models.UntrackedLot, error) {
	return []models.UntrackedLot{untrackedLot}, nil
}

//Tests

func TestGetUntrackedLotByID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.UntrackedLots+"/1")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	var lot models.UntrackedLot
	json.Unmarshal(res.Body.Bytes(), &lot)

	Assert(lot, untrackedLot, t)
}

func TestGetUntrackedLotByFakeID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.UntrackedLots+"/-1")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusNotFound, t)
}

func TestGetUntrackedLotByBadID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.UntrackedLots+"/bad_ID")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusBadRequest, t)
}

func TestGetUntrackedLots(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.UntrackedLots)
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	var lot []models.UntrackedLot
	json.Unmarshal(res.Body.Bytes(), &lot)

	Assert(len(lot), 1, t)
	Assert(lot[0], []models.UntrackedLot{untrackedLot}[0], t)
}
