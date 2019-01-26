package slugspace

import (
	"encoding/json"
		"github.com/colbyleiske/slugspace/slugspaceapi/models"
	. "github.com/colbyleiske/slugspace/utils"
		"net/http/httptest"
	"testing"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
)

func init() {
	tal := TestStoreAccessLayer{}
	tStore = NewStore(nil, tal)
}

func TestGetUntrackedLots(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.UntrackedLots)
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	var lot []models.UntrackedLot
	json.Unmarshal(res.Body.Bytes(), &lot)

	Assert(len(lot), 1, t)
	Assert(lot[0], []models.UntrackedLot{untrackedLot}[0], t)
}
