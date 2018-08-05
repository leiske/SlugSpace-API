package main

import (
	"database/sql"
	"errors"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	"github.com/colbyleiske/slugspace/utils"
	"log"
)

type DBAccessLayer struct {
	db *sql.DB
}

func (d DBAccessLayer) GetLotInfo(lotID int) (models.Lot, error) {
	lotInfo := models.Lot{}

	if err := d.db.QueryRow(utils.GetLotByID, lotID).Scan(&lotInfo.Id, &lotInfo.FullName, &lotInfo.Name, &lotInfo.Description, &lotInfo.ImageURI, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.LastUpdated); err == nil {
		return lotInfo, nil
	} else if err == sql.ErrNoRows {
		return lotInfo, errors.New("ID not found")
	} else {
		return lotInfo, err
	}
}

func (d DBAccessLayer) GetLots() ([]models.Lot, error) {
	var lots []models.Lot

	rows, err := d.db.Query(utils.GetLots)
	if err != nil {
		return lots, err
	}
	defer rows.Close()
	for rows.Next() {
		lotInfo := models.Lot{}
		if err = rows.Scan(&lotInfo.Id, &lotInfo.FullName, &lotInfo.Name, &lotInfo.Description, &lotInfo.ImageURI, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.LastUpdated); err == nil {
			lots = append(lots,lotInfo)
		} else {
			log.Fatal(err)
			continue
		}
	}
	err = rows.Err()
	if err != nil {
		return lots, err
	}

	return lots, nil
}
