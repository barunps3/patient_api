package data

import (
	"database/sql"
	"fmt"
	"time"
)

type XrayDAO struct {
	db *sql.DB
}

type IXrayDAO interface {
	GetByPatientUUID(uuid string) (Xray, error)
}

var xrayColumns = `SELECT
	xrayHash,
	uploadDate,
	uploadedBy,
	patientUUID,
	cloudLocation
	FROM patientXrays`

func (dao *XrayDAO) GetByPatientUUID(uuid string) []Xray {
	rows, err := dao.db.Query(fmt.Sprintf("%s WHERE patientUUID=$1", xrayColumns), uuid)
	if err != nil {
		fmt.Println(err)
	}

	var xrays []Xray
	var uploadDate time.Time
	for rows.Next() {
		var xray Xray
		if err := rows.Scan(
			&xray.Id,
			&uploadDate,
			&xray.UploadedBy,
			&xray.PatientUUID,
			&xray.BlobLocation,
		); err != nil {
			fmt.Sprintf("err: %v", err)
		}
		xray.UploadDate = uploadDate.Format(YYYY_MM_DD)
		xrays = append(xrays, xray)
	}
	return xrays
}
