package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/colbyleiske/slugspace/utils"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type parkingLotData []struct {
	LocationName string `json:"location_name"`
	FreeSpaces   string `json:"free_spaces"`
	DateTime     string `json:"date_time"`
}

var (
	timeFormat   = "15:04:05"
	dateFormat   = "2006-01-02"
	offsetFormat = "-07:00"
)

func main() {
	db, err := sql.Open("mysql", utils.SQLCredentials)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	httpClient := http.Client{
		Timeout: 2 * time.Second,
	}

	for id, URL := range utils.ParkingLogixAPIURL {
		request, err := http.NewRequest(http.MethodGet, URL, nil)
		if err != nil {
			panic(err)
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

		stmt, err := db.Prepare(utils.UpdateLotInfoByName)
		if err != nil {
			panic(err)
		}

		if _, err = stmt.Exec(parkingLotData[0].FreeSpaces, parkingLotData[0].DateTime, parkingLotData[0].LocationName); err != nil {
			panic(err)
		}

		fmt.Printf("Updated %s at %v\n", parkingLotData[0].LocationName, time.Now())

		//Sucks we have to replace the time zone offset. Feels bad , but in this case the API isn't returning a standard time formatting. Easiest fix to make this work with Go's formatting system.
		date, err := time.Parse(dateFormat+" "+timeFormat+" "+offsetFormat, strings.Replace(parkingLotData[0].DateTime, "-8:00", "-08:00", -1))
		if err != nil {
			panic(err)
		}

		stmt, err = db.Prepare(utils.InsertLotDatapoint)
		if err != nil {
			panic(err)
		}

		timeLastUpdated := date.Format(timeFormat)
		dateLastUpdated := date.Format(dateFormat)

		if _, err = stmt.Exec(id+1, parkingLotData[0].FreeSpaces, dateLastUpdated, timeLastUpdated); err != nil {
			panic(err)
		}

		fmt.Printf("Inserted datapoint for %s at %v\n", parkingLotData[0].LocationName, time.Now())
	}
}
