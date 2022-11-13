package patient

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const mongo_uri = "mongodb+srv://master:Master_1@cluster0.sqvis.mongodb.net/?w=majority"

const dbName string = "test"
const patientDetails string = "patient_details"
const patient_xrays string = "patient_xrays"
const patientCTScans string = "patient_ctscans"
const patientAppointments string = "patient_appointments"

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
