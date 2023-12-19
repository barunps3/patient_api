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
	blobObjectKey
	FROM patientXrays`

func (dao *XrayDAO) GetByPatientUUID(uuid string) []Xray {
	rows, err := dao.db.Query(fmt.Sprintf("%s WHERE patientUUID=$1", xrayColumns), uuid)
	if err != nil {
		fmt.Println(err)
	}
	defer dao.db.Close()

	var xrays []Xray
	var uploadDate time.Time
	for rows.Next() {
		var xray Xray
		var blobObjectKey string
		if err := rows.Scan(
			&xray.Id,
			&uploadDate,
			&xray.UploadedBy,
			&xray.PatientUUID,
			&blobObjectKey,
		); err != nil {
			fmt.Sprintf("err: %v", err)
		}
		xray.BlobUrl = generateS3Url(blobObjectKey)
		xray.UploadDate = uploadDate.Format(YYYY_MM_DD)
		xrays = append(xrays, xray)
	}
	return xrays
}
