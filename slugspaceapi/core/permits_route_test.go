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

var permit = models.Permit{
	Id:   1,
	Name: "Test Permit Name",
	Info: "Test Permit Info",
}

func (t TestStoreAccessLayer) GetPermits() ([]models.Permit, error) {
	permits := []models.Permit{permit}
	return permits, nil
}

func (t TestStoreAccessLayer) GetPermitByID(permitID int) (models.Permit, error) {
	if permitID == -1 {
		return models.Permit{}, errors.New("ID not found")
	}
	return permit, nil
}

//Tests

func TestGetPermitByID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Permits+"/1")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	var permitT models.Permit
	json.Unmarshal(res.Body.Bytes(), &permitT)

	Assert(permitT, permit, t)
}

func TestGetPermitByFakeID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Permits+"/-1")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusNotFound, t)
}

func TestGetPermitByBadID(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Permits+"/bad_ID")
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	Assert(res.Code, http.StatusBadRequest, t)
}

func TestGetPermits(t *testing.T) {
	req, _ := CreateAuthenticatedRequest(constants.Permits)
	res := httptest.NewRecorder()
	slugspace.CreateRouter(tStore).ServeHTTP(res, req)

	var permits []models.Permit
	json.Unmarshal(res.Body.Bytes(), &permits)

	Assert(len(permits), 1, t)
	Assert(permits[0], []models.Permit{permit}[0], t)
}
