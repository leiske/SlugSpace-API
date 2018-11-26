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

func TestGetUntrackedLotByID(t *testing.T) {
	req, _ := http.NewRequest("GET", constants.UntrackedLots+"/1", nil)
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	var lot models.UntrackedLot
	json.Unmarshal(res.Body.Bytes(), &lot)

	Assert(lot, untrackedLot, t)
}

func TestGetUntrackedLotByFakeID(t *testing.T) {
	req, _ := http.NewRequest("GET", constants.UntrackedLots+"/-1", nil)
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusNotFound, t)
}

func TestGetUntrackedLotByBadID(t *testing.T) {
	req, _ := http.NewRequest("GET", constants.UntrackedLots+"/bad_ID", nil)
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusBadRequest, t)
}
