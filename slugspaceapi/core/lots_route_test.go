package slugspace

import (
	"encoding/json"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	. "github.com/colbyleiske/slugspace/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
)

func init() {
	tal := TestStoreAccessLayer{}
	tStore = NewStore(nil, tal)
}

func TestGetLots(t *testing.T) {
	req, _ := http.NewRequest("GET", constants.Lots, nil)
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	var lot []models.Lot
	json.Unmarshal(res.Body.Bytes(), &lot)

	Assert(len(lot), 1, t)
	Assert(lot[0].FullName, "Core West Parking",t)
	Assert(lot[0], []models.Lot{newLot}[0], t)
}
