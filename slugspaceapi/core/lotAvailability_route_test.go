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

var testLotAvailability = models.LotAvailability {
	Id:   1,
	Name: "Test Lot Available Time",
}

func (t TestStoreAccessLayer) GetLotAvailabilities() ([]models.LotAvailability, error) {
	lotAvailabilities := []models.LotAvailability{testLotAvailability}
	return lotAvailabilities, nil
}

func (t TestStoreAccessLayer) GetLotAvailabilityByID(id int) (models.LotAvailability, error) {
	if id == -1 {
		return models.LotAvailability{}, errors.New("ID not found")
	}
	return testLotAvailability, nil
}

//Tests

func TestGetLotAvailabilityByID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.LotAvailabilities+"/1")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	var lotavailability models.LotAvailability
	json.Unmarshal(res.Body.Bytes(), &lotavailability)

	Assert(lotavailability, testLotAvailability, t)
}

func TestStore_GetLotAvailabilityByID_FakeID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.LotAvailabilities+"/-1")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusNotFound, t)
}

func TestStore_GetLotAvailabilityByID_BadID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.LotAvailabilities+"/bad_ID")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusBadRequest, t)
}

func TestStore_GetLotAvailabilities(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.LotAvailabilities)
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	var lotavailabilities []models.LotAvailability
	json.Unmarshal(res.Body.Bytes(), &lotavailabilities)

	Assert(len(lotavailabilities), 1, t)
	Assert(lotavailabilities[0], testLotAvailability, t)
}
