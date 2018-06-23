package slugspace

import (
	"fmt"
	"database/sql"
	"log"

	"github.com/colbyleiske/slugspace/utils"
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

func GetDB() (*sql.DB) {
	return db
}


func CloseDB() {
	db.Close()
	fmt.Println("DB Connection Closed")
}
