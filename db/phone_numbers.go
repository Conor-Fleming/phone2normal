package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Phone struct {
	ID     int
	Number string
}

type DB struct {
	db *sql.DB
}

func Open(driver, conString string) (*DB, error) {
	db, err := sql.Open(driver, conString)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Ping() error {
	return db.db.Ping()
}

func (db *DB) Insert() error {
	phoneData := []string{
		"(530) 514 4505",
		"(530) 517-4585",
		"530-514-4505",
		"(530)5142505",
		"1234567890",
		"123 456 7890",
		"123-456-7894",
		"(123)456-7890",
	}
	for _, numbs := range phoneData {
		_, err := insertNum(db.db, numbs)
		if err != nil {
			return err
		}
	}
	return nil
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

func (db *DB) Reset() error {
	//removing rows from table
	command := `TRUNCATE TABLE phone_numbers RESTART IDENTITY;`
	_, err := db.db.Exec(command)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetRecords() ([]string, error) {
	queryString := `SELECT number FROM phone_numbers;`
	rows, err := db.db.Query(queryString)
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
	}
	return numberSlice, nil
}

func (db *DB) PrintRecords() {
	data, err := db.GetRecords()
	if err != nil {
		panic(err)
	}
	for _, val := range data {
		fmt.Println(val)
	}

}

func (db *DB) FindPhone(newV string) error {
	recordCheckString := `SELECT number FROM phone_numbers where number = $1;`
	row := db.db.QueryRow(recordCheckString, newV)
	var number string
	err := row.Scan(&number)
	return err
}

func (db *DB) UpdateRecord(id int, newV string) error {
	updateString := `UPDATE phone_numbers SET number = $2 WHERE id = $1;`
	_, err := db.db.Exec(updateString, id, newV)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteRecord(id int) error {
	deleteString := `DELETE FROM phone_numbers WHERE id = $1;`
	_, err := db.db.Exec(deleteString, id)
	if err != nil {
		return err
	}
	return nil
}
