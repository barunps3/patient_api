package data

import (
	"time"
)

type Patient struct {
	UUID              string
	FirstName         string
	LastName          string
	Gender            string
	DateOfBirth       string
	InsuranceId       string
	PhoneNum          string
	EmergencyPhoneNum string
	Address           string
	CurrentDept       string
	CurrentCare       string
}

type ReceptionistEdit struct {
	Timestamp    time.Time
	Type         string
	EditorId     string
	PatientUUID  string
	AssignedDept string
	AssignedCare string
	Comments     string
}

type Xray struct {
	Id           string
	UploadDate   string
	UploadedBy   string
	PatientUUID  string
	BlobLocation string
}

// 	Id                   string `bson:"Id"`
// 	FirstName            string `bson:"FirstName"`
// 	MiddleName           string `bson:"MiddleName"`
// 	LastName             string `bson:"LastName"`
// 	Age                  uint   `bson:"Age"`
// 	Gender               string `bson:"Gender"`
// 	DateOfBirth          string `bson:"DateOfBirth"`
// 	HouseNumber          string `bson:"HouseNumber"`
// 	StreetName           string `bson:"StreetName"`
// 	CountryName          string `bson:"CountryName"`
// 	PhoneNumber          uint   `bson:"PhoneNumber"`
// 	ValidHealthInsurance bool   `bson:"ValidHealthInsurance"`
// 	EmergencyPhoneNumber uint   `bson:"EmergencyPhoneNumber"`
// 	EmergencyContactName string `bson:"EmergencyContactName"`
// 	IsAdmitted           bool   `bson:"IsAdmitted"`
// 	CreationDateTime     string `bson:"CreationDateTime"`
// }

// type MetaInfo struct {
// 	Id               string `bson:"Id"`
// 	CreationDateTime string `bson:"DateOfEntry"`
// }

// type XRay struct {
// 	Date     string `bson:"Date"`
// 	BodyPart string `bson:"BodyPart"`
// 	ImageUrl string `bson:"ImageUrl"`
// }

// type XRays struct {
// 	Id        string `bson:"Id"`
// 	FirstName string `bson:"FirstName"`
// 	LastName  string `bson:"LastName"`
// 	Age       uint   `bson:"Age"`
// 	Gender    string `bson:"Gender"`
// 	XRays     []XRay `bson:"XRays"`
// }

// type CTScan struct {
// 	Date     string
// 	BodyPart string
// 	ImageUrl string
// }

// type PatientCTScans struct {
// 	Id        string
// 	FirstName string
// 	LastName  string
// 	Age       string
// 	Gender    string
// 	CTScan    []CTScan
// }

// type TempReading struct {
// 	BodyTemperature float32
// 	Unit            string
// 	Time            string
// }

// type BodyTemperature struct {
// 	Date                string
// 	FullDayTempReadings []TempReading
// }

// type PatientBodyTemperatures struct {
// 	Id              string
// 	FirstName       string
// 	LastName        string
// 	Age             uint
// 	Gender          string
// 	BodyTemperature []BodyTemperature
// }

// type MedicalAppointment struct {
// 	BookingDate string
// 	Date        string
// 	Time        string
// }

// type MedicalAppointments struct {
// 	Id                 string
// 	FirstName          string
// 	LastName           string
// 	Age                uint
// 	Gender             string
// 	MedicalAppointment []MedicalAppointment
// }
