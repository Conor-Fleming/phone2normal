package main

import (
	"database/sql"
	"fmt"
	"regexp"

	dbPack "github.com/Conor-Fleming/phone2normal/db"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "conor"
	password = "password1"
	dbname   = "sandbox"
)

func main() {
	//establishing connection
	pq := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := dbPack.Open("postgres", pq)
	errCheck(err)

	defer db.Close()
	err = db.Ping()
	errCheck(err)

	fmt.Println("Connected!")

	//removing rows at each run so I dont get the same rows over and over on each run
	err = db.Reset()
	errCheck(err)

	db.Insert()

	records, err := db.GetRecords()
	db.PrintRecords()
	errCheck(err)

	fmt.Println("Normalizing....:")

	for i, val := range records {
		newVal := normalize(val)
		err = db.FindPhone(newVal)
		if err == nil {
			err = db.DeleteRecord(i + 1)
			errCheck(err)
		} else {
			if err != sql.ErrNoRows {
				errCheck(err)
			}
			err = db.UpdateRecord(i+1, newVal)
			errCheck(err)
		}
	}

	db.PrintRecords()
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func normalize(number string) string {
	//matches any char that is not a digit and replaces (line 10) with empty string
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(number, "")
}
