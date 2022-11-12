package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

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
	Id        string
	FirstName string
	LastName  string
	Age       uint
	Gender    string
	XRays     []XRay
}

type PatientBodyTemperature struct {
	Id              string
	FirstName       string
	LastName        string
	Age             uint
	Gender          string
	BodyTemperature []BodyTemperature
}

type PatientsList struct {
	Patients []Patient
}

func getPatientByFilter(filter bson.M) []Patient {
	client := getDbClient()
	patientCollection := client.Database(dbName).Collection(patientDetails)

	var result []Patient

	if filter != nil {
		cursor, err := patientCollection.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}
		err = cursor.All(context.TODO(), &result)
		if err != nil {
			panic(err)
		}
		return result
	}
	return nil
}

func addSinglePatient(patient Patient) {
	client := getDbClient()
	patientCollection := client.Database(dbName).Collection(patientDetails)

	patientBson, _ := bson.Marshal(&patient)
	_, err := patientCollection.InsertOne(context.TODO(), patientBson)

	if err != nil {
		panic(err)
	}
}

func deleteOnePatientByFilter(filter bson.M) {
	client := getDbClient()
	patientCollection := client.Database(dbName).Collection(patientDetails)

	if filter != nil {
		result, err := patientCollection.DeleteOne(context.TODO(), filter)
		fmt.Printf("result delete: %v", result)
		if err != nil {
			panic(err)
		}
	}
}

func updateOnePatientByFilter(filter bson.M, update bson.D) []Patient {
	client := getDbClient()
	patientCollection := client.Database(dbName).Collection(patientDetails)

	update = bson.D{{Key: "$set", Value: update}}
	updateResult, err := patientCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	if updateResult.MatchedCount == 1 && updateResult.ModifiedCount == 1 {
		return getPatientByFilter(filter)
	}

	return nil
}

func getPatientXraysByFilter(filter bson.M) []PatientXrayData {
	client := getDbClient()
	xraysCollection := client.Database(dbName).Collection(patientXrays)

	var result []PatientXrayData

	if filter != nil {
		cursor, err := xraysCollection.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}
		err = cursor.All(context.TODO(), &result)
		if err != nil {
			panic(err)
		}
		return result
	}
	return nil
}
