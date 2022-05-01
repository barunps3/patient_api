package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var patients = []Patient{
	{
		Id:        "1",
		FirstName: "Barun",
		LastName:  "Mazumdar",
		Age:       29,
		Gender:    "Male",
		XRays: []XRay{
			{
				Date:     "19042022",
				BodyPart: "Right Hand",
				ImageUrl: "https://patientappdata.s3.eu-central-1.amazonaws.com/xrayData/barun-mazumdar.jpg"},
			},
		BodyTemperature: []BodyTemperature{
			{
				Date: "19042022",
				FullDayTempReadings: []TempReading{
					{
						BodyTemperature: 98,
						Unit: "Celsius",
						Time: "0930",
					},
					{
						BodyTemperature: 100,
						Unit: "Celsius",
						Time: "1230",
					},
				},
			},
		},
	},
	{
		Id:        "2",
		FirstName: "Ume",
		LastName:  "Hani",
		Age:       29,
		Gender:    "Female",
	},
}

var allPatients = PatientsList{Patients: patients}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Patient App")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllPatients(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPatients: List of all patients was retrieved by Barun Mazumdar")
	
	var firstName string = r.URL.Query().Get("firstName")
	var lastName string =  r.URL.Query().Get("lastName")

	filter := get_patient_filter(firstName, lastName)
	var patients []Patient = get_patient_by_filter(filter)

	printPatientData(w, patients)
}

func returnSinglePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSinglePatient")
	vars := mux.Vars(r)
	key := vars["Id"]

	if allPatients.patientExists(key) {
		printPatientData(w, allPatients.getPatient(key))
	} else {
		fmt.Println("No patients found on Endpoint returnSinglePatient")
	}
}

func createNewPatient(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var patient Patient
	json.Unmarshal(reqBody, &patient)

	if allPatients.patientExists(patient.Id) {
		fmt.Printf("Patient with ID: %v already exists on Endpoint createNewPatient", patient.Id)
	} else {
		allPatients.addPatient(patient)
		json.NewEncoder(w).Encode(patient)
	}
}

func patchPatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: patchPatient")
	vars := mux.Vars(r)
	Id := vars["Id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var key Patient
	json.Unmarshal(reqBody, &key)

	if allPatients.patientExists(Id) {
		patientIndex := allPatients.getPatientIndex(Id)
		//fmt.Printf("patient = %v\n", patient)
		//fmt.Printf("key = %v\n", key)
		//fmt.Printf("patient.FirstName = %v\n", patient.FirstName)
		allPatients.Patients[patientIndex].FirstName = key.FirstName
	} else {
		fmt.Println("No patients found to patch")
	}
}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deletePatient")
	vars := mux.Vars(r)
	id := vars["Id"]

	if allPatients.patientExists(id) {
		allPatients.deletePatient(id)
	} else {
		fmt.Println("Endpoint Hit: Patient does not exist.")
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/patients", returnAllPatients)
	myRouter.HandleFunc("/patient/{Id}", deletePatient).Methods("DELETE")
	myRouter.HandleFunc("/patient/{Id}", returnSinglePatient).Methods("GET")
	myRouter.HandleFunc("/patient/{Id}", patchPatient).Methods("PATCH")
	myRouter.HandleFunc("/patient", createNewPatient).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
