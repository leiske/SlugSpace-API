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

func TestGetLotDataOverTimeByID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.LotDataOverTimeNoID+"/1")
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	var lotData []models.LotData
	json.Unmarshal(res.Body.Bytes(), &lotData)

	Assert(len(lotData), 1, t)
	Assert(lotData[0].FreeSpaces, 50, t)
	Assert(lotData[0], []models.LotData{trackedLotData}[0], t)
}

func TestGetLotDataOverTimeByFakeID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.LotDataOverTimeNoID+"/-1")
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	var lotData []models.LotData
	json.Unmarshal(res.Body.Bytes(), &lotData)

	Assert(res.Code, http.StatusOK, t) //Make sure we are good here still. Just return an empty dataset
	Assert(len(lotData), 0, t)
}

func TestGetLotDataOverTimeByBadID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.LotDataOverTimeNoID+"/bad_ID")
	res := httptest.NewRecorder()
	CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusBadRequest, t)
}
