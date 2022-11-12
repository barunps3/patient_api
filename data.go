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

type PatientsList struct {
	Patients []Patient
}

type XRay struct {
	Date     string
	BodyPart string
	ImageUrl string
}

type PatientXrays struct {
	Id        string
	FirstName string
	LastName  string
	Age       uint
	Gender    string
	XRays     []XRay
}

type CTScan struct {
	Date     string
	BodyPart string
	ImageUrl string
}

type PatientCTScans struct {
	Id        string
	FirstName string
	LastName  string
	Age       string
	Gender    string
	CTScan    []CTScan
}

type TempReading struct {
	BodyTemperature float32
	Unit            string
	Time            string
}

type BodyTemperature struct {
	Date                string
	FullDayTempReadings []TempReading
}

type PatientBodyTemperatures struct {
	Id              string
	FirstName       string
	LastName        string
	Age             uint
	Gender          string
	BodyTemperature []BodyTemperature
}

type MedicalAppointment struct {
	BookingDate string
	Date        string
	Time        string
}

type PatientMedicalAppointments struct {
	Id                 string
	FirstName          string
	LastName           string
	Age                uint
	Gender             string
	MedicalAppointment []MedicalAppointment
}
