package database_test

import (
	"testing"
	"github.com/colbyleiske/slugspace/slugspaceapi/core"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/database"
	"database/sql"
	"github.com/colbyleiske/slugspace/utils"
	"log"
	"os"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
)

var dbStore *slugspace.Store

func TestMain(m *testing.M) {
	db, err := sql.Open("mysql", utils.SQLCredentials)
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	dal := database.DBAccessLayer{DB: db}
	dbStore = slugspace.NewStore(db, dal)

	status := m.Run()

	dbStore.CloseDB() //close the resource

	os.Exit(status)
}

func TestGetLotByID(t *testing.T) {
	lot, err := dbStore.DAL().GetLotByID(1)

	utils.Assert(err,nil,t)
	utils.AssertNonNil(lot,t)
	utils.Assert(lot.Name,"Core West",t)
	utils.Assert(lot.FullName,"Core West Parking",t)
}

func TestGetLotByID_NotFound(t *testing.T) {
	lot, err := dbStore.DAL().GetLotByID(-1)

	utils.Assert(err,database.IDNotFoundError,t)
	utils.Assert(lot,models.Lot{},t)
}

func TestGetLots(t *testing.T) {
	lots, err := dbStore.DAL().GetLots()

	utils.Assert(err,nil,t)
	utils.AssertNonNil(lots,t)
	utils.Assert(len(lots) > 0, true, t)
	utils.Assert(lots[0].Name,"Core West",t)// Core West had ID = 1, always comes up first.
}

func TestGetUntrackedLotByID(t *testing.T) {
	lot, err := dbStore.DAL().GetUntrackedLotByID(1)

	utils.Assert(err,nil,t)
	utils.AssertNonNil(lot,t)
	utils.Assert(lot.Name,"Hahn Student Services",t)
	utils.Assert(lot.LotNumber,int64(101),t)
}

func TestGetUntrackedLotByID_NotFound(t *testing.T) {
	lot, err := dbStore.DAL().GetUntrackedLotByID(-1)

	utils.Assert(err,database.IDNotFoundError,t)
	utils.Assert(lot,models.UntrackedLot{},t)
}

func TestGetUntrackedLots(t *testing.T) {
	lots, err := dbStore.DAL().GetUntrackedLots()

	utils.Assert(err,nil,t)
	utils.AssertNonNil(lots,t)
	utils.Assert(len(lots) > 0, true, t)
	utils.Assert(lots[0].Name,"Hahn Student Services",t) // Hahn has ID = 1, always comes up first.
}

func TestGetPermitByID(t *testing.T) {
	permit, err := dbStore.DAL().GetPermitByID(1)

	utils.Assert(err,nil,t)
	utils.AssertNonNil(permit,t)
	utils.Assert(permit.Name,"A",t)
	utils.Assert(permit.Info,"",t)
}

func TestGetPermitByID_NotFound(t *testing.T) {
	permit, err := dbStore.DAL().GetPermitByID(-1)

	utils.Assert(err,database.IDNotFoundError,t)
	utils.Assert(permit,models.Permit{},t)
}

func TestGetPermits(t *testing.T) {
	permits, err := dbStore.DAL().GetPermits()

	utils.Assert(err,nil,t)
	utils.AssertNonNil(permits,t)
	utils.Assert(len(permits) > 0, true, t)
	utils.Assert(permits[0].Name,"A",t) // A has ID = 1, always comes up first.
	utils.Assert(permits[0].Info,"",t)
}

func TestGetPayStationByID(t *testing.T) {
	paystation, err := dbStore.DAL().GetPayStationByID(1)

	utils.Assert(err,nil,t)
	utils.AssertNonNil(paystation,t)
	utils.Assert(paystation.Name,"30-Min Meter",t) // 30 Min Meter has ID = 1, always comes up first.
}

func TestGetPayStationByID_NotFound(t *testing.T) {
	paystation, err := dbStore.DAL().GetPayStationByID(-1)

	utils.Assert(err,database.IDNotFoundError,t)
	utils.Assert(paystation,models.PayStation{},t)
}

func TestGetPayStations(t *testing.T) {
	paystations, err := dbStore.DAL().GetPayStations()

	utils.Assert(err,nil,t)
	utils.AssertNonNil(paystations,t)
	utils.Assert(len(paystations) > 0, true, t)
	utils.Assert(paystations[0].Name,"30-Min Meter",t) // 30 Min Meter has ID = 1, always comes up first.
}

