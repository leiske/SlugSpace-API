package database

import (
	"database/sql"
	"errors"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	"github.com/colbyleiske/slugspace/utils"
	"log"
	"time"
	"strings"
	"strconv"
	)

type DBAccessLayer struct {
	DB *sql.DB
}

var IDNotFoundError = errors.New("ID not found")

func (d DBAccessLayer) GetLotByID(id int) (models.Lot, error) {
	lotInfo := models.Lot{}
	err := d.DB.QueryRow(utils.GetLotByID, id).Scan(&lotInfo.Id, &lotInfo.FullName, &lotInfo.Name, &lotInfo.Description, &lotInfo.ImageURI, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.Longitude, &lotInfo.Latitude, &lotInfo.LastUpdated, &lotInfo.UntrackedID)
	if err == sql.ErrNoRows {
		err = IDNotFoundError
	}

	return lotInfo, err //the err will either be nil or contain something
}

func (d DBAccessLayer) GetLots() ([]models.Lot, error) {
	var lots []models.Lot

	rows, err := d.DB.Query(utils.GetLots)
	if err != nil {
		return lots, err
	}
	defer rows.Close()

	for rows.Next() {
		lotInfo := models.Lot{}
		if err = rows.Scan(&lotInfo.Id, &lotInfo.FullName, &lotInfo.Name, &lotInfo.Description, &lotInfo.ImageURI, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.Longitude, &lotInfo.Latitude, &lotInfo.LastUpdated, &lotInfo.UntrackedID); err == nil {
			lots = append(lots, lotInfo)
			continue
		}
		log.Println(err)
	}

	return lots, rows.Err()
}

func (d DBAccessLayer) GetUntrackedLotByID(id int) (models.UntrackedLot, error) {
	lotInfo := models.UntrackedLot{}
	err := d.DB.QueryRow(utils.GetUntrackedLotByID, id).Scan(&lotInfo.Id, &lotInfo.Name, &lotInfo.LotNumber, &lotInfo.Longitude, &lotInfo.Latitude, &lotInfo.Permits, &lotInfo.FreeAfter, &lotInfo.PayStations)
	if err == sql.ErrNoRows {
		err = IDNotFoundError
	}

	return lotInfo, err
}


func (d DBAccessLayer) GetUntrackedLots() ([]models.UntrackedLot, error) {
	var lots []models.UntrackedLot

	rows, err := d.DB.Query(utils.GetUntrackedLots)
	if err != nil {
		return lots, err
	}
	defer rows.Close()

	for rows.Next() {
		lotInfo := models.UntrackedLot{}
		if err = rows.Scan(&lotInfo.Id, &lotInfo.Name, &lotInfo.LotNumber, &lotInfo.Longitude, &lotInfo.Latitude, &lotInfo.Permits, &lotInfo.FreeAfter, &lotInfo.PayStations); err == nil {
			lots = append(lots, lotInfo)
			continue
		}
		log.Println(err)
	}

	return lots, rows.Err()
}

func (d DBAccessLayer) GetPermitByID(id int) (models.Permit, error) {
	permit := models.Permit{}

	err := d.DB.QueryRow(utils.GetPermitByID, id).Scan(&permit.Id, &permit.Name, &permit.Info)
	if err == sql.ErrNoRows {
		err = IDNotFoundError
	}

	return permit, err
}

func (d DBAccessLayer) GetPermits() ([]models.Permit, error) {
	var permits []models.Permit

	rows, err := d.DB.Query(utils.GetPermits)
	if err != nil {
		return permits, err
	}
	defer rows.Close()

	for rows.Next() {
		permit := models.Permit{}
		if err = rows.Scan(&permit.Id, &permit.Name, &permit.Info); err == nil {
			permits = append(permits, permit)
			continue
		}
		log.Println(err)
	}

	return permits, rows.Err()
}

func (d DBAccessLayer) GetPayStationByID(id int) (models.PayStation, error) {
	payStation := models.PayStation{}
	err := d.DB.QueryRow(utils.GetPayStationByID, id).Scan(&payStation.Id, &payStation.Name)
	if err == sql.ErrNoRows {
		err = IDNotFoundError
	}

	return payStation, err //the err will either be nil or contain something
}

func (d DBAccessLayer) GetPayStations() ([]models.PayStation, error) {
	var payStations []models.PayStation

	rows, err := d.DB.Query(utils.GetPayStations)
	if err != nil {
		return payStations, err
	}
	defer rows.Close()

	for rows.Next() {
		payStation := models.PayStation{}
		if err = rows.Scan(&payStation.Id, &payStation.Name); err == nil {
			payStations = append(payStations, payStation)
			continue
		}
		log.Println(err)
	}

	return payStations, rows.Err()
}

func (d DBAccessLayer) GetLotAvailabilityByID(id int) (models.LotAvailability, error) {
	availability := models.LotAvailability{}
	err := d.DB.QueryRow(utils.GetLotAvailabilityByID, id).Scan(&availability.Id, &availability.Name)
	if err == sql.ErrNoRows {
		err = IDNotFoundError
	}

	return availability, err //the err will either be nil or contain something
}

func (d DBAccessLayer) GetLotAvailabilities() ([]models.LotAvailability, error) {
	var availabilities []models.LotAvailability

	rows, err := d.DB.Query(utils.GetLotAvailabilities)
	if err != nil {
		return availabilities, err
	}
	defer rows.Close()

	for rows.Next() {
		availability := models.LotAvailability{}
		if err = rows.Scan(&availability.Id, &availability.Name); err == nil {
			availabilities = append(availabilities, availability)
			continue
		}
		log.Println(err)
	}

	return availabilities, rows.Err()
}

func (d DBAccessLayer) GetLotDataOverTime(id int) ([]models.LotData, error) {
	lotData := make([]models.LotData, 0)
	rows, err := d.DB.Query(utils.GetLotDataOverTimeByID, id)
	if err != nil {
		return []models.LotData{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var spaces int
		var time string
		if err = rows.Scan(&spaces, &time); err != nil {
			return []models.LotData{}, err
		}
		lotData = append(lotData, models.LotData{FreeSpaces: spaces, TimeTaken: time})
	}

	return lotData, rows.Err()
}

func (d DBAccessLayer) GetLotAverageFreespacesByDate(id int, checkDate time.Time, checkTime time.Time) (models.LotAverageFreespaces, error) {
	lotAverageFreespaces := models.LotAverageFreespaces{}

	tx , err := d.DB.Begin()
	if err != nil {
		return lotAverageFreespaces,err
	}

	if _, err := tx.Exec("SET @date = STR_TO_DATE(?, '%Y:%c:%e')",checkDate.Format("2006-1-2")); err != nil {
		tx.Rollback() //returns error
		return lotAverageFreespaces, err
	}
	if _, err := tx.Exec("SET @time = STR_TO_DATE(?, '%k:%i:%s')",checkTime.Format("3:4:5")); err != nil {
		tx.Rollback()// returns error
		return lotAverageFreespaces, err
	}

	err = tx.QueryRow(utils.GetLotAverageFreespacesByDay, id).Scan(&lotAverageFreespaces.AverageFreeSpaces)
	if err == sql.ErrNoRows {
		err = IDNotFoundError
	}
	if err != nil {
		return models.LotAverageFreespaces{},err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return models.LotAverageFreespaces{},err
	}
	return lotAverageFreespaces, nil
}

//Refactor and make cleaner......
func (d DBAccessLayer) GetTrackedLotFullInfoByID(id int) (models.TrackedLotFullInfo, error) {
	lotInfo,err := d.GetLotByID(id)
	if err != nil {
		return models.TrackedLotFullInfo{}, err
	}

	untrackedLotInfo, err := d.GetUntrackedLotByID(int(lotInfo.UntrackedID))
	if err != nil {
		return models.TrackedLotFullInfo{}, err
	}

	permits := []models.Permit{}
	for _,val := range strings.Split(strings.TrimSpace(untrackedLotInfo.Permits),",") {
		permitID , err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			return models.TrackedLotFullInfo{}, err
		}
		permit,err := d.GetPermitByID(permitID)
		if err != nil {
			return models.TrackedLotFullInfo{}, err
		}

		permits = append(permits, permit)
	}

	payStations := []models.PayStation{}
	for _,val := range strings.Split(strings.TrimSpace(untrackedLotInfo.PayStations),",") {
		payStationID , err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			return models.TrackedLotFullInfo{}, err
		}
		payStation,err := d.GetPayStationByID(payStationID)
		if err != nil {
			return models.TrackedLotFullInfo{}, err
		}

		payStations = append(payStations, payStation)
	}

	freeAfterID, err := strconv.Atoi(untrackedLotInfo.FreeAfter)
	if err != nil {
		return models.TrackedLotFullInfo{}, err
	}

	freeAfter, err := d.GetLotAvailabilityByID(freeAfterID)
	if err != nil {
		return models.TrackedLotFullInfo{}, err
	}

	fullInfo := models.TrackedLotFullInfo{
		Id:lotInfo.Id,
		Name:lotInfo.Name,
		FullName:lotInfo.FullName,
		Latitude:lotInfo.Latitude,
		Longitude:lotInfo.Longitude,
		LastUpdated:lotInfo.LastUpdated,
		FreeSpaces: lotInfo.FreeSpaces,
		TotalSpaces: lotInfo.TotalSpaces,
		Description: lotInfo.Description,
		ImageURI: lotInfo.ImageURI,
		Permits: permits,
		PayStations: payStations,
		LotAvailability: freeAfter,
	}

	return fullInfo, nil
}

//This should be a stored procedure. DBeaver sucks and Linux needs a better mariadb client lol
//Right now I just want to be demo ready... TODO refactor this
func (d DBAccessLayer) GetLotPredictedFreespaceByDateTime(lotID int, datetime time.Time) (models.LotPredictedFreespace, error) {
	tx, err := d.DB.Begin()
	if err != nil {
		return models.LotPredictedFreespace{}, err
	}

	_, err = tx.Exec("set @parkingPeriod = ?;",datetime)
	if err != nil {
		return models.LotPredictedFreespace{}, err
	}

	_, err = tx.Exec("set @lotID = ?;",lotID)
	if err != nil {
		return models.LotPredictedFreespace{}, err
	}


	_, err = tx.Exec(`set @date = convert(@parkingPeriod,DATE);`)
	if err != nil {
		tx.Rollback()
		return models.LotPredictedFreespace{}, err
	}
	_, err = tx.Exec(`set @time = convert(@parkingPeriod,TIME);`)
	if err != nil {
		return models.LotPredictedFreespace{}, err
	}

	//This is not the end-all of statistacal algorithms hahahaha
	//In talks with a data science consulting firm, gonna make this way better later
	res := tx.QueryRow("select median(freeSpaces) over (partition by lotID) as predictedFreeSpace from tbl_LotDataOverTime where ((`date` = @date) OR (`date` = DATE_SUB( @date, INTERVAL 7 DAY )) " +
		"OR (`date` = DATE_SUB( @date, INTERVAL 14 DAY )) OR (`date` = DATE_SUB( @date, INTERVAL 21 DAY ))) AND (`time` > SUBTIME(@time,\"0:15:0\") and `time` < ADDTIME(@time,\"0:15:0\")) AND lotID = @lotID LIMIT 1;")

	var freespaces int
	err = res.Scan(&freespaces)
	if err != nil {
		return models.LotPredictedFreespace{}, err
	}

	isPredicted := datetime.After(time.Now())

	err = tx.Commit()
	if err != nil {
		return models.LotPredictedFreespace{}, err
	}

	return models.LotPredictedFreespace{PredictedFreespace: freespaces, IsPredicted: isPredicted}, nil

}