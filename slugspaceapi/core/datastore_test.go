package slugspace

import (
	"database/sql"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	. "github.com/colbyleiske/slugspace/utils"
	"testing"
)

type TestStoreAccessLayer struct{}

func (t TestStoreAccessLayer) GetLotInfo(lotID int) (models.Lot, error) {
	return models.Lot{}, nil
}

func TestNewStore(t *testing.T) {
	db, _ := sql.Open("", "")
	tal := TestStoreAccessLayer{}
	s := NewStore(db, tal)

	AssertNonNil(s.db, t)
	AssertNonNil(s.dal, t)

	Assert(s.db, db, t)
	Assert(s.dal, tal, t)
}
