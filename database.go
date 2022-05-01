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
var collections string = "patient_details"


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
	patient_collection := client.Database(dbName).Collection(collections)

	var result []Patient

	if filter != nil {
		cursor, err := patient_collection.Find(context.TODO(), filter)
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
	patient_collection := client.Database(dbName).Collection(collections)

	patientBson, _ := bson.Marshal(&patient)
	_, err := patient_collection.InsertOne(context.TODO(), patientBson)

	if err != nil {
		panic(err)
	}
}
