package slugspace

import (
	"github.com/colbyleiske/slugspace/slugspaceapi/core/middleware"
	"github.com/gorilla/mux"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
)

func CreateRouter(s *Store) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)

	//route registration
	router.Handle(constants.LotByIDFull, s.GetLotByID()).Methods("GET")
	router.Handle(constants.Lots, s.GetLots()).Methods("GET")
	router.Handle(constants.LotDataOverTimeFull, s.GetLotDataOverTime()).Methods("GET")

	return router
}
