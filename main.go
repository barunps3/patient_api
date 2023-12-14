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
	myRouter.HandleFunc("/xrays/{uuid}", api.GetXrays).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
