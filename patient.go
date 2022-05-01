package main


type Patient struct {
	Id                   string `bson:"Id"`
	FirstName            string `bson:"FirstName"`
	LastName             string `bson:"LastName"`
	Age                  uint   `bson:"Age"`
	Gender               string `bson:"Gender"`
	DateOfBirth          string `bson:"DateOfBirth"`
	HouseNumber          string `bson:"HouseNumber"`
	StreetName           string `bson:"StreetName"`
	CountryName          string `bson:"CountryName"`
	PhoneNumber          uint   `bson:"PhoneNumber"`
	ValidHealthInsurance bool   `bson:"ValidHealthInsurance"`
	EmergencyPhoneNumber uint   `bson:"EmergencyPhoneNumber"`
	EmergencyContactName string `bson:"EmergencyContactName"`
	IsAdmitted           bool   `bson:"IsAdmitted"`
}

type PatientXrayData struct {
	Id                   string
	FirstName            string
	LastName             string
	Age                  uint
	Gender               string
	XRays                []XRay
}

type PatientBodyTemperature struct {
	Id                   string
	FirstName            string
	LastName             string
	Age                  uint
	Gender               string
	BodyTemperature      []BodyTemperature
}

type PatientsList struct {
	Patients []Patient
}


func (patients PatientsList) getPatient(id string) *Patient {
	for _, patient := range patients.Patients {
		if patient.Id == id {
			return &patient
		}
	}
	return nil
}

func (patients PatientsList) getPatientIndex(id string) int {
	for index, patient := range patients.Patients {
		if patient.Id == id {
			return index
		}
	}
	return -1
}

func (patients PatientsList) patientExists(id string) bool {
	if patients.getPatient(id) == nil {
		return false
	} else {
		return true
	}
}


func (patients *PatientsList) addPatient(patient Patient) {
	patients.Patients = append(patients.Patients, patient)
}


func (patients *PatientsList) deletePatient(id string) {
	index := patients.getPatientIndex(id)
	patients.Patients = append(patients.Patients[:index], patients.Patients[index+1:]...)	
}


