package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

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

func getPatientXraysByFilter(filter bson.M) []PatientXrays {
	client := getDbClient()
	xraysCollection := client.Database(dbName).Collection(patientXrays)

	var result []PatientXrays

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

func getPatientCTScanByFilter(filter bson.M) []PatientCTScans {
	client := getDbClient()
	ctScansCollection := client.Database(dbName).Collection(patientCTScans)

	var result []PatientCTScans

	if filter != nil {
		cursor, err := ctScansCollection.Find(context.TODO(), filter)
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

func getPatientAppointmentsByFilter(filter bson.M) []PatientMedicalAppointments {
	client := getDbClient()
	appointmentCollection := client.Database(dbName).Collection(patientAppointments)
	var result []PatientMedicalAppointments

	if filter != nil {
		cursor, err := appointmentCollection.Find(context.TODO(), filter)
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
