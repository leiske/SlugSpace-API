package slugspace

import (
	"encoding/json"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	. "github.com/colbyleiske/slugspace/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	tal := TestStoreAccessLayer{}
	tStore = NewStore(nil, tal)
}

func TestGetLotByID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Lots+"/1")
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	var lot models.Lot
	json.Unmarshal(res.Body.Bytes(), &lot)

	Assert(lot, trackedLot, t)
}

func TestGetLotByFakeID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Lots+"/-1")
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusNotFound, t)
}

func TestGetLotByBadID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Lots+"/bad_ID")
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusBadRequest, t)
}
