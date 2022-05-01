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

func get_db_client() *mongo.Client {
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

func get_patient_by_filter(filter bson.M) []Patient {
	client := get_db_client()
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

func get_patient_filter(firstName string, lastName string) bson.M {
	if firstName != "" {
		if lastName != "" {
			return bson.M{"FirstName": firstName, "LastName": lastName}
		} else {
			return bson.M{"FirstName": firstName}
		}
	} else if lastName != "" {
		return bson.M{"LastName": lastName}
	} else {
		return nil
	}
}
