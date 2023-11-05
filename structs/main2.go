package main

import (
	"log"
 	"time"
)


type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	user := User {
		FirstName: "Wisdom",
		LastName: "Elue",
		PhoneNumber: "00000000000",
		Age: 25,
	}

	log.Println(user.FirstName, user.LastName)
}