package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "conor"
	password = ""
	dbname   = "sandbox"
)

func main() {
	pq := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", pq)
	errCheck(err)

	defer db.Close()

	err = db.Ping()
	errCheck(err)

	fmt.Println("Connected!")
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
