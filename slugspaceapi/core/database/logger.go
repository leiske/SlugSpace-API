package database

import (
	"github.com/colbyleiske/slugspace/utils"
	"log"
	)

//Uses variadic strings. Limit to three ALWAYS... message is first, ip is second, error is third. IP and Error are optional
func (dal DBAccessLayer) Log(category string, severity int, messages ...string) {
	stmt, err := dal.DB.Prepare(utils.InsertLog)
	if err != nil {
		log.Printf("ERROR: Couldn't submit the following log under category %v and %v severity:\n", category, severity)
		log.Println(messages)
	}

	if _, err = stmt.Exec(category, severity, messages); err != nil {
		log.Printf("ERROR: Couldn't submit the following log under category %v and %v severity:\n", category, severity)
		log.Println(messages)
	}
}
