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

var testPayStation = models.PayStation {
	Id:   1,
	Name: "Test PayStation Name",
}

func (t TestStoreAccessLayer) GetPayStations() ([]models.PayStation, error) {
	PayStation := []models.PayStation{testPayStation}
	return PayStation, nil
}

func (t TestStoreAccessLayer) GetPayStationByID(payStationID int) (models.PayStation, error) {
	if payStationID == -1 {
		return models.PayStation{}, errors.New("ID not found")
	}
	return testPayStation, nil
}

//Tests

func TestGetPayStationByID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.PayStations+"/1")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	var paystation models.PayStation
	json.Unmarshal(res.Body.Bytes(), &paystation)

	Assert(paystation, testPayStation, t)
}

func TestGetPayStationByFakeID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.PayStations+"/-1")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusNotFound, t)
}

func TestGetPayStationByBadID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.PayStations+"/bad_ID")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusBadRequest, t)
}

func TestGetPayStations(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.PayStations)
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	var paystations []models.PayStation
	json.Unmarshal(res.Body.Bytes(), &paystations)

	Assert(len(paystations), 1, t)
	Assert(paystations[0], []models.PayStation{testPayStation}[0], t)
}
