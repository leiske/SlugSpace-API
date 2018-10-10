package slugspace

import (
	"github.com/colbyleiske/slugspace/utils"
	"log"
)

//My crappy "enums"

type Category string

const (
	AUTH     = Category("AUTH")
	DATA     = Category("DATA")
	INTERNET = Category("INTERNET")
)

type Severity int

const (
	HIGH = Severity(3)
	MID  = Severity(2)
	LOW  = Severity(1)
)

//Uses variadic strings. Limit to three ALWAYS... message is first, ip is second, error is third. IP and Error are optional
func (s *Store) Log(category Category, severity Severity, messages ...string) {
	stmt, err := s.db.Prepare(utils.InsertLog)
	if err != nil {
		log.Printf("ERROR: Couldn't submit the following log under category %v and %v severity:\n", category, severity)
		log.Println(messages)
	}

	if _, err = stmt.Exec(category, severity, messages); err != nil {
		log.Printf("ERROR: Couldn't submit the following log under category %v and %v severity:\n", category, severity)
		log.Println(messages)
	}
}
