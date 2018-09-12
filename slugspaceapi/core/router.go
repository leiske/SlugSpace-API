package slugspace

import (
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/middleware"
	"github.com/gorilla/mux"
)

func CreateRouter(s *Store) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)

	//route registration
	router.Handle(constants.LotByIDFull, middleware.AuthenticationMiddleware(s.GetLotByID(),s.db)).Methods("GET")
	router.Handle(constants.Lots, middleware.AuthenticationMiddleware(s.GetLots(),s.db)).Methods("GET")
	router.Handle(constants.LotDataOverTimeFull, middleware.AuthenticationMiddleware(s.GetLotDataOverTime(),s.db)).Methods("GET")
	router.Handle(constants.RegisterAppInstance, s.PostRegisterAppInstance()).Methods("POST") //todo: secure this route
	//router.Handle(constants.LotAverageFreespaceByDay,middleware.AuthenticationMiddleware(s.GetLotAverageFreespaces())).Methods("GET")


	return router
}
