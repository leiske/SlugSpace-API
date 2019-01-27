package slugspace_test

import (
	"time"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
)

func (t TestStoreAccessLayer) GetLotAverageFreespacesByDate(lotID int, checkDate time.Time, checkTime time.Time) (models.LotAverageFreespaces, error) {
	return models.LotAverageFreespaces{}, nil //temp
}

//tests

//none yet