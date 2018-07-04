package main

import (
	"log"
	"net/http"
	"time"
	"os"
	"os/signal"
	"context"
	"flag"
	"github.com/colbyleiske/slugspace/slugspaceapi/core"
	"fmt"
	"database/sql"
	"github.com/colbyleiske/slugspace/utils"
	"github.com/gorilla/mux"
)

var s *slugspace.Store

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", 15*time.Second, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	db, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	dal := DBAccessLayer{db:db}
	s = slugspace.NewStore(db,dal)
	defer s.CloseDB()

	router := mux.NewRouter()
	router.Use(LoggerMiddleware)

	router.Handle("/v1/lot/{lotID}",GetLotByID(s)).Methods("GET")

	srv := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		log.Printf("Server started")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	//Graceful Shutdowns. Courtesy of Mux github
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c //block until CTRL+C signal is given

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting down")
}


func connectToDB() (*sql.DB,error) {
	fmt.Println("Connecting to DB")

	db, err := sql.Open("mysql", utils.SQLCredentials)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Checking connection...")

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	fmt.Println("Connected")
	return db,nil
}
