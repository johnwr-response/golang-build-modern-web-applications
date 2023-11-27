package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// Basic variables
	var myString string
	var myInt int
	myString = "Hi"
	myInt = 11
	mySecondString := "Another string"
	log.Println(myString, mySecondString, myInt)

	// String maps:
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	myMap["dog"] = "Fido"
	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	// Int maps:
	myMapInt := make(map[string]int)
	myMapInt["First"] = 1
	myMapInt["Second"] = 2
	log.Println(myMapInt["First"], myMapInt["Second"])

	// Struct maps:
	myMapUser := make(map[string]User)
	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}
	myMapUser["me"] = me
	log.Println(myMapUser["me"].FirstName)

	// Slices:
	var myStringSlice []string
	myStringSlice = append(myStringSlice, "Trevor")
	myStringSlice = append(myStringSlice, "John")
	myStringSlice = append(myStringSlice, "Mary")
	log.Println(myStringSlice)
	var myIntSlice []int
	myIntSlice = append(myIntSlice, 2)
	myIntSlice = append(myIntSlice, 1)
	myIntSlice = append(myIntSlice, 3)
	log.Println(myIntSlice)

	sort.Ints(myIntSlice)
	log.Println(myIntSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)
	log.Println(numbers[0:2])
	log.Println(numbers[6:9])

	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
