package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const postgresConnStr = "user=barun dbname=patients sslmode=disable"

// patient DAO factory
func newPatientDAO() (*patientDAO, error) {
	db, err := sql.Open("postgres", postgresConnStr)
	if err != nil {
		fmt.Println("Could not connect to Database: ", err)
		return nil, err
	}
	return &patientDAO{db: db}, nil
}
