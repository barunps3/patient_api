package main

type XRay struct {
	Date     string
	BodyPart string
	ImageUrl string
}

type CTScan struct {
	Date     string
	BodyPart string
	ImageUrl string
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

type MedicalAppointment struct {
	Id          string
	IsBooked    bool
	BookingDate string
	Date        string
	Time        string
}
