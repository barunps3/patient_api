package data

import (
	"database/sql"
	"fmt"
	"time"
)

type IPatientDAO interface {
	GetAll() ([]Patient, error)
	GetByUUID(id string) (Patient, error)
	Add(patient Patient) error
}

type PatientDAO struct {
	db *sql.DB
}

var PatientColumns = `SELECT 
    patientId,
    firstName,
    lastName,
    gender,
    dateOfBirth,
    COALESCE(insuranceId, '') AS insuranceId,
    phoneNum,
    COALESCE(emergencyPhoneNum, '') AS emergencyPhoneNum,
    address
    FROM patients`

func (dao *PatientDAO) GetAll() ([]Patient, error) {
	rows, err := dao.db.Query(PatientColumns)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("row:", rows)

	var patients []Patient
	for rows.Next() {
		var patient Patient
		var timeOfBirth time.Time
		if err := rows.Scan(
			&patient.Id,
			&patient.FirstName,
			&patient.LastName,
			&patient.Gender,
			&timeOfBirth,
			&patient.InsuranceId,
			&patient.PhoneNum,
			&patient.EmergencyPhoneNum,
			&patient.Address,
		); err != nil {
			fmt.Errorf("err: %v", err)
		}
		patient.DateOfBirth = timeOfBirth.Format("2006-01-02")
		patients = append(patients, patient)
	}
	return patients, nil
}

func (dao *PatientDAO) GetByUUID(id string) Patient {
	row := dao.db.QueryRow(fmt.Sprintf("%s WHERE patientId=$1", PatientColumns), id)
	defer dao.db.Close()

	var patient Patient
	var timeOfBirth time.Time
	if err := row.Scan(
		&patient.Id,
		&patient.FirstName,
		&patient.LastName,
		&patient.Gender,
		&timeOfBirth,
		&patient.InsuranceId,
		&patient.PhoneNum,
		&patient.EmergencyPhoneNum,
		&patient.Address,
	); err != nil {
		fmt.Printf("Could not fetch ID of patient from DB: %v", err)
	}

	patient.DateOfBirth = timeOfBirth.Format("2006-01-02")
	return patient
}

// func GetByFilter(filter bson.M) []Identity {
// 	client := getDbClient()
// 	patientCollection := client.Database(dbName).Collection(patientDetails)

// 	var result []Identity

// 	if filter != nil {
// 		cursor, err := patientCollection.Find(context.TODO(), filter)
// 		if err != nil {
// 			panic(err)
// 		}
// 		err = cursor.All(context.TODO(), &result)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return result
// 	}
// 	return nil
// }

// func CreateOneIdentity(patient *Identity) {
// 	patient.CreationDateTime = "14.11.2022-14:19:00"
// 	patient.Id = generatePatientId(patient)
// 	patient.Age = getPatientAge(patient)

// 	generatePatientDocument(patient)
// 	generateRelatedRegisters(patient)
// }

// func generatePatientId(patient *Identity) string {
// 	var input_string string = strings.Join([]string{patient.FirstName,
// 		patient.LastName, patient.DateOfBirth, patient.CreationDateTime}, "")
// 	var id_bytes [16]byte = md5.Sum([]byte(input_string))
// 	return hex.EncodeToString(id_bytes[:])
// }

// func getPatientAge(patient *Identity) uint {
// 	return 30
// }

// func generatePatientDocument(patient *Identity) {
// 	client := getDbClient()
// 	patientCollection := client.Database(dbName).Collection(patientDetails)

// 	patientBson, _ := bson.Marshal(patient)
// 	_, err := patientCollection.InsertOne(context.TODO(), patientBson)

// 	if err != nil {
// 		panic(err)
// 	}
// }

// func generateRelatedRegisters(patient *Identity) {
// 	createMetaInfo(patient)
// 	createEmptyAppointmentRegister(patient)
// 	createEmptyCTScanRegister(patient)
// 	createEmptyXRayRegister(patient)
// }

// func createEmptyXRayRegister(patient *Identity) {
// 	client := getDbClient()
// 	xraysCollection := client.Database(dbName).Collection(patient_xrays)

// 	var patientXrays XRays
// 	patientXrays.Id = patient.Id
// 	patientXrays.FirstName = patient.FirstName
// 	patientXrays.LastName = patient.LastName
// 	patientXrays.Gender = patient.Gender
// 	patientXrays.Age = patient.Age
// 	patientXrays.XRays = []XRay{}

// 	patientXRaysBson, _ := bson.Marshal(patientXrays)
// 	_, err := xraysCollection.InsertOne(context.TODO(), patientXRaysBson)

// 	if err != nil {
// 		panic(err)
// 	}
// }

// func createEmptyCTScanRegister(identity *Identity) {

// }

// func createEmptyAppointmentRegister(identity *Identity) {

// }

// func createMetaInfo(identiy *Identity) {

// }

// func DeleteOneByFilter(filter bson.M) {
// 	client := getDbClient()
// 	patientCollection := client.Database(dbName).Collection(patientDetails)

// 	if filter != nil {
// 		result, err := patientCollection.DeleteOne(context.TODO(), filter)
// 		fmt.Printf("result delete: %v", result)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }

// func UpdateOneByFilter(filter bson.M, update bson.D) []Identity {
// 	client := getDbClient()
// 	patientCollection := client.Database(dbName).Collection(patientDetails)

// 	update = bson.D{{Key: "$set", Value: update}}
// 	updateResult, err := patientCollection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if updateResult.MatchedCount == 1 && updateResult.ModifiedCount == 1 {
// 		return GetByFilter(filter)
// 	}

// 	return nil
// }

// func GetXraysByFilter(filter bson.M) []XRays {
// 	client := getDbClient()
// 	xraysCollection := client.Database(dbName).Collection(patient_xrays)

// 	var result []XRays

// 	if filter != nil {
// 		cursor, err := xraysCollection.Find(context.TODO(), filter)
// 		if err != nil {
// 			panic(err)
// 		}
// 		err = cursor.All(context.TODO(), &result)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return result
// 	}
// 	return nil
// }

// func GetCTScanByFilter(filter bson.M) []PatientCTScans {
// 	client := getDbClient()
// 	ctScansCollection := client.Database(dbName).Collection(patientCTScans)

// 	var result []PatientCTScans

// 	if filter != nil {
// 		cursor, err := ctScansCollection.Find(context.TODO(), filter)
// 		if err != nil {
// 			panic(err)
// 		}
// 		err = cursor.All(context.TODO(), &result)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return result
// 	}
// 	return nil
// }

// func GetAppointmentsByFilter(filter bson.M) []MedicalAppointments {
// 	client := getDbClient()
// 	appointmentCollection := client.Database(dbName).Collection(patientAppointments)
// 	var result []MedicalAppointments

// 	if filter != nil {
// 		cursor, err := appointmentCollection.Find(context.TODO(), filter)
// 		if err != nil {
// 			panic(err)
// 		}
// 		err = cursor.All(context.TODO(), &result)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return result
// 	}
// 	return nil
// }
