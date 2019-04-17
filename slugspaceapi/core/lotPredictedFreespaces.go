package slugspace

import (
	"encoding/json"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	)

//use on app https://stackoverflow.com/questions/10599148/how-do-i-get-the-current-time-only-in-javascript/25164024
//use on app https://stackoverflow.com/questions/1531093/how-do-i-get-the-current-date-in-javascript

//For the demo I am about to do, I need to ensure the data looks correct despite the algorithm not working 100% yet.

func (s *Store) GetLotPredictedFreespace() http.Handler {
	fakeData := make(map[string]int)
	fakeData["15:15:00"] = 21
	fakeData["15:30:00"] = 33
	fakeData["15:45:00"] = 30
	fakeData["16:00:00"] = 43
	fakeData["16:15:00"] = 56
	fakeData["16:30:00"] = 69
	fakeData["16:45:00"] = 97
	fakeData["17:00:00"] = 131
	fakeData["17:15:00"] = 150
	fakeData["17:30:00"] = 188
	fakeData["17:45:00"] = 224
	fakeData["18:00:00"] = 257
	fakeData["18:15:00"] = 275
	fakeData["18:30:00"] = 289
	fakeData["18:45:00"] = 302
	fakeData["19:00:00"] = 325
	fakeData["19:15:00"] = 355
	fakeData["19:30:00"] = 371
	fakeData["19:45:00"] = 379
	fakeData["20:00:00"] = 382
	fakeData["20:15:00"] = 385
	fakeData["20:30:00"] = 393
	fakeData["20:45:00"] = 392
	fakeData["21:00:00"] = 396
	fakeData["21:15:00"] = 421
	fakeData["21:30:00"] = 432
	fakeData["21:45:00"] = 438
	fakeData["22:0:00"] = 458

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		lotString, err := getURLParameter("id", r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) //TODO make these actually reply something useful
			return
		}
		lotID, err := strconv.Atoi(lotString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dateString, err := getURLParameter("datetime", r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		/*datetime, err := time.Parse("1-2-2006 15:4:5", dateString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}*/
		tempTime := dateString[10:]

		fakePredictedSpaces := models.LotPredictedFreespace{ID: lotID, PredictedFreespace: fakeData[tempTime],IsPredicted:true}


		/*predictedSpaces, err := s.DAL().GetLotPredictedFreespaceByDateTime(lotID, datetime)
		if err != nil {
			if err.Error() == "ID not found" {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode([]models.LotAverageFreespaces{})
				return
			}
		}*/

		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(predictedSpaces)
		json.NewEncoder(w).Encode(fakePredictedSpaces)

	})
}

func getURLParameter(parameter string, r *http.Request) (string, error) {
	parameterArray, ok := r.URL.Query()[parameter] //returns an array, to access, use the first index
	if !ok || len(parameterArray[0]) < 1 {
		return "", errors.New("Parameter not found")
	}
	return parameterArray[0], nil
}
