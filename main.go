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
	//establishing connection
	pq := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", pq)
	errCheck(err)

	defer db.Close()

	err = db.Ping()
	errCheck(err)

	fmt.Println("Connected!")

	//inserting some records
	_, err = insertNum(db, "(530) 514 4505")
	errCheck(err)
	_, err = insertNum(db, "(530) 514-4505")
	errCheck(err)
	_, err = insertNum(db, "530-514-4505")
	errCheck(err)
	_, err = insertNum(db, "(530)5144505")
	errCheck(err)
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

func insertNum(db *sql.DB, number string) (int, error) {
	statement := `INSERT INTO phone_numbers(number) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, number).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
