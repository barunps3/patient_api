package main

import (
	"encoding/json"
	"io"
	"fmt"
)

func printPatientData(writeOutput io.Writer, patients any) {
	enc := json.NewEncoder(writeOutput)
	enc.SetIndent("", " ")
	enc.Encode(patients)
}

func writeJson(writeOutput io.Writer, patient Patient) {
	fmt.Println(patient)
	fmt.Println("Using write json")
	b, _ := json.MarshalIndent(patient, "", "\t")
	fmt.Println(string(b))
	_, _ = writeOutput.Write(b)
}
