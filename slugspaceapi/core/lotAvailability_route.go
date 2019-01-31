package slugspace

import (
	"net/http"
			"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
	"encoding/json"
)

func (s *Store) GetLotAvailabilityByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		id, err := s.GetID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		availability, err := s.DAL().GetLotAvailabilityByID(id)
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
		json.NewEncoder(w).Encode(availability)
	})
}

func (s *Store) GetLotAvailabilities() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		availabilities, err := s.DAL().GetLotAvailabilities()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.dal.Log(constants.DATA,constants.HIGH,err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(availabilities)
	})
}
