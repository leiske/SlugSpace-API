package slugspace

import (
	"fmt"
	"database/sql"
	"log"

	"github.com/colbyleiske/SlugSpace/utils"
	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

func ConnectToDB() {
	fmt.Println("Connecting to DB")

	var err error
	db, err = sql.Open("mysql", utils.SQLCredentials)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Checking connection...")

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	fmt.Println("Connected")

}

func GetLotInfo(lotID int) Lot {
	lotInfo := Lot{}
	rows, err := db.Query(utils.GetLotByID, lotID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&lotInfo.Id, &lotInfo.Name, &lotInfo.FreeSpaces, &lotInfo.TotalSpaces, &lotInfo.LastUpdated)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lotInfo
}

func CloseDB() {
	db.Close()
	fmt.Println("DB Connection Closed")
}
