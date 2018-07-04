package main

import (
	"database/sql"
	"github.com/colbyleiske/slugspace/utils"
	"errors"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
)

type DBAccessLayer struct {
	db *sql.DB
}

func (d DBAccessLayer)GetLotInfo(lotID int) (models.Lot,error) {
	lotInfo := models.Lot{}

	if err := d.db.QueryRow(utils.GetLotByID,lotID).Scan(&lotInfo.Id, &lotInfo.Name, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.LastUpdated); err == nil {
		return lotInfo,nil
	} else if err == sql.ErrNoRows {
		return lotInfo, errors.New("ID not found")
	} else {
		return lotInfo,err
	}
}
