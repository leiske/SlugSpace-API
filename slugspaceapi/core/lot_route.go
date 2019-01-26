package slugspace

import (
	"net/http"

	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"log"
)

func (s *Store) GetLotByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		vars := mux.Vars(r)
		lotID, err := strconv.Atoi(vars["lotID"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		lotInfo, err := s.DAL().GetLotInfo(lotID)
		if err != nil {
			if err.Error() == "ID not found" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(lotInfo)
	})
}

func (s *Store) GetLots() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		lots, err := s.DAL().GetLots()
		if err != nil {
			log.Fatal(err)
			if err.Error() == "ID not found" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(lots)
	})
}
