/*
 * SlugSpace
 *
 * SlugSpace API - Realtime Parking Data and Metrics.
 *
 * API version: 0.0.1
 * Contact: colby.leiske@gmail.com
 */

package slugspace

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routeList {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routeList = Routes{
	Route{
		"GetLotByID",
		http.MethodGet,
		"/v1/lot/{lotID}",
		GetLotByID,
	},
}
