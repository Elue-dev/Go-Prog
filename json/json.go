package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func main() {
	myJson := `
		[
			{
				"first_name": "Dubem",
				"last_name": "Elue"
			},
			{
				"first_name": "Wisdom",
				"last_name": "Elue"
			}
		]
	`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error Unmarshalling JSON")
	}

	fmt.Printf("unmarshalled : %v", unmarshalled)

	// write json from a struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Luchi"
	m1.LastName = "Elue"

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Awele"
	m2.LastName = "Elue"

	mySlice = append(mySlice, m2)
	newJson, err := json.MarshalIndent(mySlice, "", "      ")
	
	if err != nil {
		fmt.Println("Error Marshalling", err)
	}

		fmt.Println(string(newJson))

}