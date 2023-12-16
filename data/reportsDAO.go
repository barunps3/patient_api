package data

import (
	"database/sql"
	"fmt"
	"time"
)

type ReportsDAO struct {
	db *sql.DB
}

type IReportsDAO interface {
	GetByPatientUUID(uuid string) (Xray, error)
}

var reportColumns = `SELECT
	DISTINCT uploadDate
	FROM patientXrays`

func (dao *ReportsDAO) GetByPatientUUID(uuid string) Reports {
	rows, err := dao.db.Query(fmt.Sprintf("%s WHERE patientUUID=$1", reportColumns), uuid)
	if err != nil {
		fmt.Println(err)
	}

	var reports Reports
	reports.PatientUUID = uuid
	var uploadDate time.Time
	for rows.Next() {
		if err := rows.Scan(
			&uploadDate,
		); err != nil {
			fmt.Sprintf("err: %v", err)
		}
		reports.Xrays = append(reports.Xrays, uploadDate.Format(YYYY_MM_DD))
	}
	// fmt.Println("reports:", reports)
	return reports
}
