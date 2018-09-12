package slugspace

import (
	"net/http"
	"time"
	"github.com/pkg/errors"
	"encoding/json"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	"strconv"
)

//use on app https://stackoverflow.com/questions/10599148/how-do-i-get-the-current-time-only-in-javascript/25164024
//use on app https://stackoverflow.com/questions/1531093/how-do-i-get-the-current-date-in-javascript
func (s *Store) GetLotAverageFreespaces() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		lotString, err := getURLParameter("id",r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) //TODO make these actually reply something useful
			return
		}
		lotID, err := strconv.Atoi(lotString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dateString, err := getURLParameter("date",r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		timeString, err := getURLParameter("time",r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) //TODO make these actually reply something useful
			return
		}

		checkDate , err := time.Parse("1-2-2006",dateString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		checkTime,err := time.Parse("3:4:5",timeString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}


		lotData, err := s.DAL().GetLotAverageFreespacesByDate(lotID, checkDate,checkTime)
		if err != nil {
			if err.Error() == "ID not found" {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode([]models.LotAverageFreespaces{})
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(lotData)
	})
}

func getURLParameter(parameter string,r *http.Request)(string, error) {
	parameterArray , ok := r.URL.Query()[parameter] //returns an array, to access, use the first index
	if !ok || len(parameterArray[0]) < 1 {
		return "", errors.New("Parameter not found")
	}
	return parameterArray[0], nil
}
