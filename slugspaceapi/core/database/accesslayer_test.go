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
	utils.Assert(lot.LotName,"Hahn Student Services",t)
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
	utils.Assert(lots[0].LotName,"Hahn Student Services",t)// Hahn has ID = 1, always comes up first.
}

func TestGetPermitByID(t *testing.T) {
	permit, err := dbStore.DAL().GetPermitByID(1)

	utils.Assert(err,nil,t)
	utils.AssertNonNil(permit,t)
	utils.Assert(permit.PermitName,"A",t)
	utils.Assert(permit.PermitInfo,"",t)
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
	utils.Assert(permits[0].PermitName,"A",t)// A has ID = 1, always comes up first.
	utils.Assert(permits[0].PermitInfo,"",t)// A has ID = 1, always comes up first.

}