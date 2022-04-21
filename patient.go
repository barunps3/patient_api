package main


type Patient struct {
	Id                   string
	FirstName            string
	LastName             string
	Age                  uint
	Gender               string
	DataOfBirth          string
	HouseNumber          string
	StreetName           string
	CountryName          string
	PhoneNumber          uint
	ValidHealthInsurance bool
	EmergencyPhoneNumber uint
	EmergencyContactName string
	IsAdmmitted          bool
	XRays                []XRay
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


