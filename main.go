package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"medicalApp/api"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Patient App")
	fmt.Println("Endpoint Hit: homePage")
}

// func returnSinglePatient(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnSinglePatient")
// 	vars := mux.Vars(r)
// 	patientId := vars["Id"]
// 	filterMap := bson.M{"Id": patientId}

// 	var patient []patient.Identity = patient.GetByFilter(filterMap)
// 	if patient != nil {
// 		printPatientData(w, patient)
// 	} else {
// 		fmt.Println("No patients found on Endpoint returnSinglePatient")
// 	}
// }

// func createNewPatient(w http.ResponseWriter, r *http.Request) {
// 	reqBody, _ := io.ReadAll(r.Body)
// 	var patientInfo patient.Identity
// 	json.Unmarshal(reqBody, &patientInfo)

// 	if patient.GetByFilter(bson.M{"Id": patientInfo.Id}) != nil {
// 		fmt.Printf("Patient with ID: %v already exists on Endpoint createNewPatient", patientInfo.Id)
// 	} else {
// 		// TODO: Validate before insertion
// 		patient.CreateOneIdentity(&patientInfo)
// 		json.NewEncoder(w).Encode(patientInfo)
// 	}
// }

// func patchPatient(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: patchPatient")
// 	vars := mux.Vars(r)
// 	Id := vars["Id"]
// 	reqBody, _ := io.ReadAll(r.Body)

// 	var patchMap map[string]interface{}
// 	json.Unmarshal(reqBody, &patchMap)

// 	fmt.Printf("map[string]interface{} format patch: %v", patchMap)

// 	update := bson.D{}
// 	for k, v := range patchMap {
// 		fmt.Printf("type of value: %T", v)
// 		update = append(update, bson.E{Key: k, Value: v})
// 	}

// 	var filter = bson.M{"Id": Id}

// 	updatedPatient := patient.UpdateOneByFilter(filter, update)
// 	if updatedPatient != nil {
// 		printPatientData(w, updatedPatient)
// 	} else {
// 		fmt.Println("No patients found to patch")
// 	}
// }

// func deletePatient(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: deletePatient")
// 	vars := mux.Vars(r)
// 	patientId := vars["Id"]
// 	fmt.Print(patientId)
// 	filterMap := bson.M{"Id": patientId}

//		var patients []patient.Identity = patient.GetByFilter(filterMap)
//		if patients != nil {
//			patient.DeleteOneByFilter(filterMap)
//		} else {
//			fmt.Println("Endpoint Hit: Patient does not exist.")
//		}
//	}
//
// func getpatientXrays(w http.ResponseWriter, r *http.Request) { // 	fmt.Println("Endpoint Hit: getpatient.Xrays") // 	vars := mux.Vars(r) // 	patientId := vars["Id"] // 	filterMap := bson.M{"Id": patientId} // 	var xrays []patient.XRays = patient.GetXraysByFilter(filterMap) // 	if xrays != nil { // 		printPatientData(w, xrays) // 	} else { // 		fmt.Println("No Xray found on Endpoint getpatient.Xrays") // 	} // } // func getVaccinations() {}
// func getCTScans() {}

// Custom middleware to set standard headers
func standardHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set standard headers here
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// Use the custom middleware for all routes
	myRouter.Use(standardHeadersMiddleware)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/patient/{uuid}", api.GetPatientByUUID).Methods("GET")
	myRouter.HandleFunc("/patient", api.GetPatientByIdType).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
