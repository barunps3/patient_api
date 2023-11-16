package api

import (
	"fmt"
	"medicalApp/data"
)

func example() {
	test := data.PatientExample{Example: "abbc"}
	fmt.Println(test)
}
