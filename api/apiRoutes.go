package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"medicalApp/data"

	"github.com/gorilla/mux"
)

func GetPatientByUUID(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	reqPatientId := queryParams["uuid"]

	var patientDAO = data.NewPatientDAO()
	patient := patientDAO.GetByUUID(reqPatientId)
	fmt.Println(patient)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

func GetPatientByIdType(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Search by Nation ID type")

	idType := r.URL.Query().Get("idType")
	idValue := r.URL.Query().Get("idVal")

	if len(idType) != 0 && len(idValue) != 0 {
		var patientDAO = data.NewPatientDAO()
		patient := patientDAO.GetByIdType(idType, idValue)
		fmt.Println(patient)
		json.NewEncoder(w).Encode(patient)
	} else {
		http.NotFound(w, r)
	}
}

func GetXrays(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	reqPatientId := queryParams["uuid"]

	var xraysDao = data.NewXraysDAO()
	xrays := xraysDao.GetByPatientUUID(reqPatientId)
	fmt.Println(xrays)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(xrays)
}
