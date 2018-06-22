package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/colbyleiske/ParkingLogixAPICaller/utils"

	_ "github.com/go-sql-driver/mysql"
)

type parkingLotData []struct {
	LocationName string `json:"location_name"`
	FreeSpaces   string `json:"free_spaces"`
	DateTime     string `json:"date_time"`
}

func main() {
	db, err := sql.Open("mysql", utils.SQLCredentials)
	if err != nil {
		log.Fatal(err) //fix this to write out to a text file that I can check up on every so often
	}
	defer db.Close()

	httpClient := http.Client{
		Timeout: 2 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, utils.ParkingLogixAPIURL, nil)
	if err != nil {
		panic(err) //write to DB with error ??
	}

	request.Header.Add(utils.ParkingLogixAPIHeader, utils.ParkingLogixAPIKey)
	request.Header.Set("User-Agent", "slugspace-data-gatherer")

	response, err := httpClient.Do(request)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	parkingLotData := parkingLotData{}
	if err = json.Unmarshal(body, &parkingLotData); err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE slugspace.LotOverview SET FreeSpaces= ?, LastUpdated= ? WHERE Name= ?")
	if err != nil {
		panic(err)
	}

	if _, err = stmt.Exec(parkingLotData[0].FreeSpaces, parkingLotData[0].DateTime, parkingLotData[0].LocationName); err != nil {
		panic(err)
	}
}
