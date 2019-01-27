package slugspace

import (
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/middleware"
	"github.com/gorilla/mux"
	"net/http"
		)

func CreateRouter(s *Store) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)

	//route registration
	router.Handle(constants.LotByID, s.AuthMiddleware(s.GetLotByID())).Methods("GET")
	router.Handle(constants.Lots, s.AuthMiddleware(s.GetLots())).Methods("GET")

	router.Handle(constants.UntrackedLots, s.AuthMiddleware(s.GetUntrackedLots())).Methods("GET")
	router.Handle(constants.UntrackedLotsByID, s.AuthMiddleware(s.GetUntrackedLotByID())).Methods("GET")

	router.Handle(constants.Permits, s.AuthMiddleware(s.GetPermits())).Methods("GET")
	router.Handle(constants.PermitByID, s.AuthMiddleware(s.GetPermitByID())).Methods("GET")

	router.Handle(constants.PayStations, s.AuthMiddleware(s.GetPayStations())).Methods("GET")
	router.Handle(constants.PayStationByID, s.AuthMiddleware(s.GetPayStationByID())).Methods("GET")

	router.Handle(constants.LotAvailabilities, s.AuthMiddleware(s.GetLotAvailabilities())).Methods("GET")
	router.Handle(constants.LotAvailabilityByID, s.AuthMiddleware(s.GetLotAvailabilityByID())).Methods("GET")

	router.Handle(constants.LotDataOverTimeFull, s.AuthMiddleware(s.GetLotDataOverTime())).Methods("GET")

	router.Handle(constants.RegisterAppInstance, s.PostRegisterAppInstance()).Methods("POST") //todo: secure this route

	router.Handle(constants.LotAverageFreespaceByDay,s.AuthMiddleware(s.GetLotAverageFreespaces())).Methods("GET")

	return router
}

func (s *Store)AuthMiddleware(endpoint http.Handler) (http.Handler) {
	return middleware.AuthenticationMiddleware(endpoint, s.dal.GetTokenSecret)
}

