package data

import (
	"database/sql"
	"fmt"
	"time"
)

type PatientDAO struct {
	db *sql.DB
}

type IPatientDAO interface {
	GetAll() ([]Patient, error)
	GetByUUID(id string) (Patient, error)
	Add(patient Patient) error
}

var patientColumns = `SELECT 
    patientUUID,
    firstName,
    lastName,
    gender,
    dateOfBirth,
    COALESCE(insuranceId, '') AS insuranceId,
    phoneNum,
    COALESCE(emergencyPhoneNum, '') AS emergencyPhoneNum,
    address
    FROM patients`

var receptionistEditsColumns = `SELECT
	editdatetime,
	edittype,
	editorid,
	patientuuid,
	assigneddept,
	assignedcare,
	comments
	FROM receptionistedits`

func (dao *PatientDAO) GetAll() ([]Patient, error) {
	rows, err := dao.db.Query(patientColumns)
	defer dao.db.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("row:", rows)

	var patients []Patient
	for rows.Next() {
		var patient Patient
		var timeOfBirth time.Time
		if err := rows.Scan(
			&patient.UUID,
			&patient.FirstName,
			&patient.LastName,
			&patient.Gender,
			&timeOfBirth,
			&patient.InsuranceId,
			&patient.PhoneNum,
			&patient.EmergencyPhoneNum,
			&patient.Address,
		); err != nil {
			return nil, fmt.Errorf("err: %v", err)
		}
		patient.DateOfBirth = timeOfBirth.Format("2006-01-02")
		patients = append(patients, patient)
	}
	return patients, nil
}

func (dao *PatientDAO) GetByUUID(id string) Patient {
	fmt.Println(id)
	row := dao.db.QueryRow(fmt.Sprintf("%s WHERE patientUUID=$1", patientColumns), id)
	defer dao.db.Close()

	var patient Patient
	var timeOfBirth time.Time
	if err := row.Scan(
		&patient.UUID,
		&patient.FirstName,
		&patient.LastName,
		&patient.Gender,
		&timeOfBirth,
		&patient.InsuranceId,
		&patient.PhoneNum,
		&patient.EmergencyPhoneNum,
		&patient.Address,
	); err != nil {
		fmt.Printf("Could not fetch from patients table: %v", err)
	}

	var receptionistEdit ReceptionistEdit
	row = dao.db.QueryRow(fmt.Sprintf("%s WHERE patientUUID=$1", receptionistEditsColumns), id)
	if err := row.Scan(
		&receptionistEdit.Timestamp,
		&receptionistEdit.Type,
		&receptionistEdit.EditorId,
		&receptionistEdit.PatientUUID,
		&receptionistEdit.AssignedDept,
		&receptionistEdit.AssignedCare,
		&receptionistEdit.Comments,
	); err != nil {
		fmt.Printf("Could not fetch from receptionistedits table: %v", err)
	}

	patient.CurrentCare = receptionistEdit.AssignedCare
	patient.CurrentDept = receptionistEdit.AssignedDept
	patient.DateOfBirth = timeOfBirth.Format("2006-01-02")
	return patient
}

func (dao *PatientDAO) GetByIdType(idType string, idVal string) Patient {
	row := dao.db.QueryRow(`SELECT patientuuid FROM patientids
		WHERE idtype=$1 AND idvalue=$2`, idType, idVal)
	defer dao.db.Close()
	var patient Patient
	if err := row.Scan(&patient.UUID); err != nil {
		fmt.Printf("Could not fetch ID of patient from DB: %v", err)
	}
	fmt.Println(patient.UUID)
	patient = dao.GetByUUID(patient.UUID)
	return patient
}
