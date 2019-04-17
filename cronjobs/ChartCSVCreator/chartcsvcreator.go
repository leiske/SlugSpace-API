package main

import (
			"log"
		"encoding/csv"
	"os"
		time2 "time"
	"strconv"
		"github.com/iancoleman/orderedmap"
	"fmt"
	"io/ioutil"
)

func main() {
	rawdata, err := os.Open("/home/valid/go/src/github.com/colbyleiske/slugspace/cronjobs/ChartCSVCreator/rawdata.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer rawdata.Close()

	csvreader := csv.NewReader(rawdata)

	csvreader.Comma = '	'

	csvData, err := csvreader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	o := orderedmap.New()
	for _,v := range csvData {
		timeVal , err := time2.Parse("15:04:05", v[2])
		if err != nil {
			log.Fatal(err)
		}
		/*dateVal, err := time2.Parse("2006-01-02",v[0]) // dateval
		if err != nil {
			log.Fatal(err)
		}*/
		freespace , err := strconv.Atoi(v[1])
		if err != nil {
			log.Println(err)
			continue
		}

		roundedTime := timeVal.Round(5*time2.Minute).Format("15:04:05.999999999")
		if val, ok := o.Get(roundedTime); !ok {
			o.Set(roundedTime,[]int{freespace})
		} else {
			o.Set(roundedTime, append(val.([]int), freespace))
		}
	}


	output := ""
	for _, k := range o.Keys() {
		output += fmt.Sprintf("%v,",k)
		val, _ := o.Get(k)
		for k, v := range val.([]int) {
			comma := ","
			if k == len(val.([]int)) - 1 {
				comma = "\n"
			}
			output += fmt.Sprintf("%v%v",v,comma)
		}
	}
	err = ioutil.WriteFile("/home/valid/go/src/github.com/colbyleiske/slugspace/cronjobs/ChartCSVCreator/chartdata.txt",[]byte(output),0644)
	if err != nil {
		log.Fatal(err)
	}


}

//problem
//need to sort by large groups and the insertion order of them inbound is random.......
//make map of strings to int array
//keep appending the location of that onto it. We then sort