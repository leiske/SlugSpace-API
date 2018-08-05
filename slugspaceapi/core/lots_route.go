package slugspace

import (
	"net/http"

	"encoding/json"
	"log"
)

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
