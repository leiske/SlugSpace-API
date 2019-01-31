package slugspace

import (
	"net/http"

	"encoding/json"
			"github.com/colbyleiske/slugspace/slugspaceapi/core/constants"
)

func (s *Store) GetPermitByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		id, err := s.GetID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		permitInfo, err := s.DAL().GetPermitByID(id)
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
		json.NewEncoder(w).Encode(permitInfo)
	})
}

func (s *Store) GetPermits() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		permits, err := s.DAL().GetPermits()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.dal.Log(constants.DATA,constants.HIGH,err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(permits)
	})
}
