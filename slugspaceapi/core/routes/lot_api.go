/*
 * SlugSpace
 *
 * SlugSpace API - Realtime Parking Data and Metrics.
 *
 * API version: 0.0.1
 * Contact: colby.leiske@gmail.com
 */

package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
	"github.com/colbyleiske/slugspace/utils"
	"database/sql"
	"errors"
)

func GetLotByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	lotID,err := strconv.Atoi(vars["lotID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lotInfo,err := GetLotInfo(lotID)
	if err != nil {
		if err.Error() == "ID not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lotInfo)
}

func GetLotInfo(lotID int) (Lot,error) {
	lotInfo := Lot{}

	if err := db.QueryRow(utils.GetLotByID,lotID).Scan(&lotInfo.Id, &lotInfo.Name, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.LastUpdated); err == nil {
		return lotInfo,nil
	} else if err == sql.ErrNoRows {
		return lotInfo, errors.New("ID not found")
	} else {
		return lotInfo,err
	}

}
