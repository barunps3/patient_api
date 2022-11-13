package main

import (
	"encoding/json"
	"fmt"
	"io"
	"medicalApp/patient"
)

func printPatientData(writeOutput io.Writer, patients any) {
	enc := json.NewEncoder(writeOutput)
	enc.SetIndent("", " ")
	enc.Encode(patients)
}

func writeJson(writeOutput io.Writer, patient patient.Identity) {
	fmt.Println(patient)
	fmt.Println("Using write json")
	b, _ := json.MarshalIndent(patient, "", "\t")
	fmt.Println(string(b))
	_, _ = writeOutput.Write(b)
}
