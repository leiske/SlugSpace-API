package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	logPath    *string
	logPrefix  *string
	dateFormat *string
)

func main() {

	setupFlags()

	log.Printf("Clearing old logs in %v\n", filepath.Dir(*logPath))

	files, err := ioutil.ReadDir(filepath.Dir(*logPath))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() && file.ModTime().Before(time.Now().AddDate(0, 0, -14)) {
			err = os.Remove(filepath.Dir(*logPath) + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Printf("Reading %v\n", *logPath)

	logContents, err := ioutil.ReadFile(*logPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Compressing %v\n", *logPath)

	var buffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&buffer)
	gzipWriter.Write(logContents)
	gzipWriter.Close()
	err = ioutil.WriteFile(filepath.Dir(*logPath)+"/"+*logPrefix+"_"+time.Now().Format(*dateFormat)+".log.gz", buffer.Bytes(), 777)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Removing %v\n", *logPath)

	err = os.Remove(*logPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Creating new %v\n", *logPath)

	newLog, err := os.Create(*logPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer newLog.Close()

}

func setupFlags() {
	logPath = flag.String("logpath", "", "Full path to the log file to rotate")
	logPrefix = flag.String("logprefix", "api_caller", "Prefix for all rotated logs. EG: api_caller_09-12-2018.log")
	dateFormat = flag.String("dateformat", "01-02-2006", "Format to use when outputting the date. Follows Go's standard formatting system")
	flag.Parse()
}
