/*
 * SlugSpace
 *
 * SlugSpace API - Realtime Parking Data and Metrics.
 *
 * API version: 0.0.1
 * Contact: colby.leiske@gmail.com
 */

package main

import (
	"log"
	"net/http"
	"github.com/colbyleiske/slugspace/core"
	"time"
	"os"
	"os/signal"
	"context"
	"flag"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", 15*time.Second, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	slugspace.ConnectToDB()
	defer slugspace.CloseDB() //Uncomment this and please commit later :)
	log.Printf("Server started")

	router := slugspace.NewRouter()

	srv := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
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
	os.Exit(0)
}
