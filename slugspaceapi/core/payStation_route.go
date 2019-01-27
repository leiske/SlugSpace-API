package slugspace

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
	"encoding/json"
)

func (s *Store) GetPayStationByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		vars := mux.Vars(r)
		stationID, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		stationInfo, err := s.DAL().GetPayStationByID(stationID)
		if err != nil {
			if err.Error() == "ID not found" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			s.dal.Log(constants.DATA,constants.HIGH,err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stationInfo)
	})
}

func (s *Store) GetPayStations() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		payStations, err := s.DAL().GetPayStations()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.dal.Log(constants.DATA,constants.HIGH,err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(payStations)
	})
}
