package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const postgresConnStr = "user=barun dbname=patients sslmode=disable"

func createSqlClient() *sql.DB {
	client, err := sql.Open("postgres", postgresConnStr)
	if err != nil {
		fmt.Println("Could not connect to Database: ", err)
		return nil
	}
	return client
}

// patient DAO factory
func NewPatientDAO() *PatientDAO {
	db := createSqlClient()
	return &PatientDAO{db: db}
}

// Xray DAO factory
func NewXraysDAO() *XrayDAO {
	db := createSqlClient()
	return &XrayDAO{db: db}
}
