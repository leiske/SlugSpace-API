package slugspace_test

import (
	"database/sql"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/database"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
	. "github.com/colbyleiske/slugspace/utils"
	"testing"
	"net/http"
	"log"
	"github.com/colbyleiske/slugspace/slugspaceapi/core"
)

type TestStoreAccessLayer struct{}

var tStore *slugspace.Store

func init() {
	tal := TestStoreAccessLayer{}
	tStore = slugspace.NewStore(nil, tal)
}


func (t TestStoreAccessLayer) CreateJWT(payload *database.JWTPayload) (string, error) {
	return "", nil //temp
}

func (t TestStoreAccessLayer) GetTokenSecret(guid interface{}) (interface{}, bool, error) {
	return []byte(constants.TestSecret), true, nil
}

func (t TestStoreAccessLayer) Log(category string, severity int, messages ...string) {
	log.Printf("CATEGORY: %v || SEVERITY: %v || MESSAGE: %v",category,severity,messages)
}

func CreateAuthenticatedRequest(endpoint string) (*http.Request, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization",constants.TestToken)
	return req, nil
}

func TestNewStore(t *testing.T) {
	db, _ := sql.Open("", "")
	tal := TestStoreAccessLayer{}
	s := slugspace.NewStore(db, tal)

	AssertNonNil(s.DB(), t)
	AssertNonNil(s.DAL(), t)

	Assert(s.DB(), db, t)
	Assert(s.DAL(), tal, t)
}
