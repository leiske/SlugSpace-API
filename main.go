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
	"github.com/colbyleiske/slugspace/go"
)

func main() {
	log.Printf("Server started")

	router := slugspace.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
