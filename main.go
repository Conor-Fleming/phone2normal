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

	records, err := getRecords(db)
	errCheck(err)

	fmt.Println("Normalizing....")
	err = normalizeRecords(db, records)
	errCheck(err)

	records, err = getRecords(db)
	errCheck(err)
	//fmt.Println(results)
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

func getRecords(db *sql.DB) ([]string, error) {
	queryString := `SELECT number FROM phone_numbers;`
	rows, err := db.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	numberSlice := make([]string, 0)
	for rows.Next() {
		var number string
		err = rows.Scan(&number)
		if err != nil {
			return nil, err
		}
		numberSlice = append(numberSlice, number)
		fmt.Println(number)
	}
	return numberSlice, nil
}

func resetTable(db *sql.DB) error {
	command := `TRUNCATE TABLE phone_numbers RESTART IDENTITY;`
	_, err := db.Exec(command)
	if err != nil {
		return err
	}
	return nil
}

func normalizeRecords(db *sql.DB, records []string) error {
	updateString := `UPDATE phone_numbers SET number = $2 WHERE id = $1;`
	for i, v := range records {
		newV := normalize(v)
		//fmt.Println("this should be inserted:", newV)
		_, err := db.Exec(updateString, i+1, newV)
		if err != nil {
			return err
		}
	}
	return nil
}
