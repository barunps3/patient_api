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
	Id          string
	UploadDate  string
	UploadedBy  string
	PatientUUID string
	BlobUrl     string
}

type Reports struct {
	PatientUUID string
	Xrays       []string
}

const (
	YYYY_MM_DD string = "2006-01-02"
)
