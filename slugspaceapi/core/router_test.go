package slugspace

import (
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
	. "github.com/colbyleiske/slugspace/utils"
	"testing"
)

func TestCreateRouter(t *testing.T) {
	router := CreateRouter(tStore)

	//Ensure our routes are here
	AssertNonNil(router.Get(constants.LotDataOverTimeFull), t)
	AssertNonNil(router.Get(constants.Lots), t)
	AssertNonNil(router.Get(constants.LotByIDFull), t)
}
