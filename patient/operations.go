package patient

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetByFilter(filter bson.M) []Identity {
	client := getDbClient()
	patientCollection := client.Database(dbName).Collection(patientDetails)

	var result []Identity

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

func AddOne(patient Identity) {
	client := getDbClient()
	patientCollection := client.Database(dbName).Collection(patientDetails)

	patientBson, _ := bson.Marshal(&patient)
	_, err := patientCollection.InsertOne(context.TODO(), patientBson)

	if err != nil {
		panic(err)
	}
}

func DeleteOneByFilter(filter bson.M) {
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

func UpdateOneByFilter(filter bson.M, update bson.D) []Identity {
	client := getDbClient()
	patientCollection := client.Database(dbName).Collection(patientDetails)

	update = bson.D{{Key: "$set", Value: update}}
	updateResult, err := patientCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	if updateResult.MatchedCount == 1 && updateResult.ModifiedCount == 1 {
		return GetByFilter(filter)
	}

	return nil
}

func GetXraysByFilter(filter bson.M) []XRays {
	client := getDbClient()
	xraysCollection := client.Database(dbName).Collection(patient_xrays)

	var result []XRays

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

func GetCTScanByFilter(filter bson.M) []PatientCTScans {
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

func GetAppointmentsByFilter(filter bson.M) []PatientMedicalAppointments {
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
