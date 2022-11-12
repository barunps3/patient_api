package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Patient App")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllPatients(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPatients: List of all patients was retrieved by Barun Mazumdar")

	var filterMap = bson.M{}
	for _, param := range patientFilterParams {
		var paramVal string = r.URL.Query().Get(param)
		if paramVal != "" {
			filterMap[param] = paramVal
		}
	}
	var patients []Patient = getPatientByFilter(filterMap)
	printPatientData(w, patients)
}

func returnSinglePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSinglePatient")
	vars := mux.Vars(r)
	patientId := vars["Id"]
	filterMap := bson.M{"Id": patientId}

	var patient []Patient = getPatientByFilter(filterMap)
	if patient != nil {
		printPatientData(w, patient)
	} else {
		fmt.Println("No patients found on Endpoint returnSinglePatient")
	}
}

func createNewPatient(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var patient Patient
	json.Unmarshal(reqBody, &patient)

	if getPatientByFilter(bson.M{"Id": patient.Id}) != nil {
		fmt.Printf("Patient with ID: %v already exists on Endpoint createNewPatient", patient.Id)
	} else {
		addSinglePatient(patient)
		json.NewEncoder(w).Encode(patient)
	}
}

func patchPatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: patchPatient")
	vars := mux.Vars(r)
	Id := vars["Id"]
	reqBody, _ := ioutil.ReadAll(r.Body)

	var patchMap map[string]interface{}
	json.Unmarshal(reqBody, &patchMap)

	fmt.Printf("map[string]interface{} format patch: %v", patchMap)

	update := bson.D{}
	for k, v := range patchMap {
		fmt.Printf("type of value: %T", v)
    update = append(update, bson.E{Key: k, Value: v})
	}

	var filter = bson.M{"Id": Id}

	updatedPatient := updateOnePatientByFilter(filter, update)
	if updatedPatient != nil {
		printPatientData(w, updatedPatient)
	} else {
		fmt.Println("No patients found to patch")
	}
}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deletePatient")
	vars := mux.Vars(r)
	patientId := vars["Id"]
	fmt.Printf(patientId)
	filterMap := bson.M{"Id": patientId}

	var patients []Patient = getPatientByFilter(filterMap)
	if patients != nil {
		deleteOnePatientByFilter(filterMap)
	} else {
		fmt.Println("Endpoint Hit: Patient does not exist.")
	}
}

func getPatientXrays(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getPatientXrays")
	vars := mux.Vars(r)
	patientId := vars["Id"]
	filterMap := bson.M{"Id": patientId}

	var xrays []PatientXrayData = getPatientXraysByFilter(filterMap)
	if xrays != nil {
		printPatientData(w, xrays)
	} else {
		fmt.Println("No Xray found on Endpoint getPatientXrays")
	}
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/patients", returnAllPatients)
	myRouter.HandleFunc("/patient/{Id}", deletePatient).Methods("DELETE")
	myRouter.HandleFunc("/patient/{Id}", returnSinglePatient).Methods("GET")
	myRouter.HandleFunc("/patient/{Id}", patchPatient).Methods("PATCH")
	myRouter.HandleFunc("/patient", createNewPatient).Methods("POST")

	myRouter.HandleFunc("/xrays/{Id}", getPatientXrays).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
