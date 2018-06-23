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
	"strconv"
	"log"
	"encoding/json"
)

func GetLotByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	lotID,err := strconv.Atoi(vars["lotID"])
	if err != nil {
		log.Panic(err)
	}
	lotInfo := GetLotInfo(lotID)
	json.NewEncoder(w).Encode(lotInfo)
}
