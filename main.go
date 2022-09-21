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
	password = "password1"
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

	err = resetTable(db)
	errCheck(err)

	//inserting some records
	_, err = insertNum(db, "(530) 514 4505")
	errCheck(err)
	_, err = insertNum(db, "(530) 517-4585")
	errCheck(err)
	_, err = insertNum(db, "530-514-4505")
	errCheck(err)
	_, err = insertNum(db, "(530)5142505")
	errCheck(err)
	_, err = insertNum(db, "1234567890")
	errCheck(err)
	_, err = insertNum(db, "123 456 7890")
	errCheck(err)
	_, err = insertNum(db, "123-456-7894")
	errCheck(err)
	_, err = insertNum(db, "(123)456-7890")
	errCheck(err)

	checkRecords(db)
	fmt.Println("Normalizing....")
	normalizeRecords(db)
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
	statement := `INSERT INTO phone_numbers (number) VALUES ($1) RETURNING id;`
	var id int
	err := db.QueryRow(statement, number).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func checkRecords(db *sql.DB) error {
	queryString := `SELECT number FROM phone_numbers;`
	rows, err := db.Query(queryString)
	if err != nil {
		return err
	}
	defer rows.Close()

	//numberSlice := make([]string, 0)
	i := 0
	for rows.Next() {
		var number string
		err = rows.Scan(&number)
		if err != nil {
			return err
		}
		fmt.Println(i, number)
		i++
	}
	return nil
}

func resetTable(db *sql.DB) error {
	command := `TRUNCATE TABLE phone_numbers RESTART IDENTITY;`
	_, err := db.Exec(command)
	if err != nil {
		return err
	}
	return nil
}

func normalizeRecords(db *sql.DB) error {
	queryString := `SELECT number FROM phone_numbers;`
	rows, err := db.Query(queryString)
	if err != nil {
		return err
	}
	defer rows.Close()

	//numberSlice := make([]string, 0)
	for rows.Next() {
		var number string
		err = rows.Scan(&number)
		if err != nil {
			return err
		}
		//numberSlice = append(numberSlice, number)
		fmt.Println(normalize(number))
	}
	return nil
}
