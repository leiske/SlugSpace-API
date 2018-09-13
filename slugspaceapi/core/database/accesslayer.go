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

func (d DBAccessLayer) GetLotInfo(lotID int) (models.Lot, error) {
	lotInfo := models.Lot{}

	if err := d.DB.QueryRow(utils.GetLotByID, lotID).Scan(&lotInfo.Id, &lotInfo.FullName, &lotInfo.Name, &lotInfo.Description, &lotInfo.ImageURI, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.LastUpdated); err == nil {
		return lotInfo, nil
	} else if err == sql.ErrNoRows {
		return lotInfo, errors.New("ID not found")
	} else {
		return lotInfo, err
	}
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
		if err = rows.Scan(&lotInfo.Id, &lotInfo.FullName, &lotInfo.Name, &lotInfo.Description, &lotInfo.ImageURI, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.LastUpdated); err == nil {
			lots = append(lots, lotInfo)
		} else {
			log.Println(err)
			continue
		}
	}

	if err = rows.Err(); err != nil {
		return lots, err
	}

	return lots, nil
}

func (d DBAccessLayer) GetLotDataOverTime(lotID int) ([]models.LotData, error) {
	lotData := make([]models.LotData, 0)
	rows, err := d.DB.Query(utils.GetLotDataOverTimeByID, lotID)
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

	if err = rows.Err(); err != nil {
		return []models.LotData{}, err
	}

	return lotData, nil
}

func (d DBAccessLayer) GetLotAverageFreespacesByDate(lotID int, checkDate time.Time , checkTime time.Time) (models.LotAverageFreespaces,error) {
	lotAverageFreespaces := models.LotAverageFreespaces{}
	//
	//tx , err := d.db.Begin()
	//if err != nil {
	//	return lotAverageFreespaces,err
	//}
	//
	//if _, err := tx.Exec("SET @date = STR_TO_DATE(?, '%Y:%c:%e')",checkDate.Format("2006-1-2")); err != nil {
	//	tx.Rollback() //returns error
	//	return lotAverageFreespaces, err
	//}
	//if _, err := tx.Exec("SET @time = STR_TO_DATE(?, '%k:%i:%s')",checkTime.Format("3:4:5")); err != nil {
	//	tx.Rollback()// returns error
	//	return lotAverageFreespaces, err
	//}
	//
	//if err := tx.QueryRow(utils.GetLotAverageFreespacesByDay, lotID).Scan(&lotAverageFreespaces.AverageFreeSpaces); err == sql.ErrNoRows {
	//	return lotAverageFreespaces, errors.New("ID not found")
	//} else if err != nil {
	//	return models.LotAverageFreespaces{},err
	//}
	//
	//if err := tx.Commit(); err != nil {
	//	tx.Rollback()
	//	return models.LotAverageFreespaces{},err
	//}
	return lotAverageFreespaces,nil
}

