package slugspace

import (
	"net/http"

	"encoding/json"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
		)

func (s *Store) GetLotDataOverTime() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		id, err := s.GetID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		lotData, err := s.DAL().GetLotDataOverTime(id)
		if err != nil {
			if err.Error() == "ID not found" {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode([]models.LotData{})
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(lotData)
	})
}
