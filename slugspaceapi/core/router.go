package slugspace

import (
	"github.com/gorilla/mux"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/middleware"
)

func CreateRouter(s *Store) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)

	//route registration
	router.Handle("/v1/lot/{lotID}", s.GetLotByID()).Methods("GET")
	return router
}
