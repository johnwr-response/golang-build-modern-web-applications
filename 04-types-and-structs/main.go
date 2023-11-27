package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	var s2 = "six"
	// s := "eight" // Don't do this, it's confusing which s you set
	log.Println("s is", s)
	log.Println("s2 is", s2)
	saySomething("xxx")

	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "1-555-555-1212",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s3, "World"
}
