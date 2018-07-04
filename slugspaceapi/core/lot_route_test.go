package slugspace

import (
	"encoding/json"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	. "github.com/colbyleiske/slugspace/utils"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

var newLot = models.Lot{
	Id:          0,
	Name:        "Core West",
	FreeSpaces:  50,
	TotalSpaces: 100,
	LastUpdated: "2018",
}

var tStore *Store

func init() {
	tal := TestRouteAccessLayer{}
	tStore = NewStore(nil, tal)
}

type TestRouteAccessLayer struct{}

func (t TestRouteAccessLayer) GetLotInfo(lotID int) (models.Lot, error) {
	if lotID == -1 {
		return models.Lot{}, errors.New("ID not found")
	}
	return newLot, nil
}

func TestGetLotByID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/lot/1", nil)
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	var lot models.Lot
	json.Unmarshal(res.Body.Bytes(), &lot)

	Assert(lot, newLot, t)
}

func TestGetLotByFakeID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/lot/-1", nil)
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusNotFound, t)
}

func TestGetLotByBadID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/lot/bad_ID", nil)
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusBadRequest, t)
}
