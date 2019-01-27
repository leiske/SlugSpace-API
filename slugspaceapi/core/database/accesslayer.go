package database

import (
	"database/sql"
	"errors"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	"github.com/colbyleiske/slugspace/utils"
	"log"
	"time"
)

type DBAccessLayer struct {
	DB *sql.DB
}

var IDNotFoundError = errors.New("ID not found")

func (d DBAccessLayer) GetLotByID(id int) (models.Lot, error) {
	lotInfo := models.Lot{}
	err := d.DB.QueryRow(utils.GetLotByID, id).Scan(&lotInfo.Id, &lotInfo.FullName, &lotInfo.Name, &lotInfo.Description, &lotInfo.ImageURI, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.Longitude, &lotInfo.Latitude, &lotInfo.LastUpdated)
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
		if err = rows.Scan(&lotInfo.Id, &lotInfo.FullName, &lotInfo.Name, &lotInfo.Description, &lotInfo.ImageURI, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.Longitude, &lotInfo.Latitude, &lotInfo.LastUpdated); err == nil {
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

