package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"medicalApp/data"

	"github.com/gorilla/mux"
)

var patientDAO = data.NewPatientDAO()

func GetPatientByUUID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSinglePatient")
	queryParams := mux.Vars(r)
	reqPatientId := queryParams["Id"]

	patient := patientDAO.GetByUUID(reqPatientId)
	fmt.Println(patient)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}
