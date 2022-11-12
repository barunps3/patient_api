package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const mongo_uri = "mongodb+srv://master:Master_1@cluster0.sqvis.mongodb.net/?w=majority"

var dbName string = "test"
var patientDetails string = "patient_details"
var patientXrays string = "patient_xrays"


func getDbClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connection established")
	return client
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
