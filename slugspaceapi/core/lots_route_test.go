package slugspace

import (
	"encoding/json"
		"github.com/colbyleiske/slugspace/slugspaceapi/models"
	. "github.com/colbyleiske/slugspace/utils"
		"net/http/httptest"
	"testing"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
)


func TestGetLots(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Lots)
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	var lot []models.Lot
	json.Unmarshal(res.Body.Bytes(), &lot)

	Assert(len(lot), 1, t)
	Assert(lot[0].FullName, "Core West Parking", t)
	Assert(lot[0], []models.Lot{trackedLot}[0], t)
}
