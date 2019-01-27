package slugspace_test

import (
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
	. "github.com/colbyleiske/slugspace/utils"
	"testing"
	"github.com/colbyleiske/slugspace/slugspaceapi/core"
)

func TestCreateRouter(t *testing.T) {
	router := slugspace.CreateRouter(tStore)

	//Ensure our routes are here
	AssertNonNil(router.Get(constants.LotDataOverTimeFull), t)
	AssertNonNil(router.Get(constants.Lots), t)
	AssertNonNil(router.Get(constants.LotByID), t)
	AssertNonNil(router.Get(constants.RegisterAppInstance), t)

}
