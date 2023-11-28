package main

import (
	"log"
)

func main() {

	// Loop a fixed number of timer
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	// Loop a slice
	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for _, animal := range animals {
		log.Println(animal)
	}

	// Loop a map
	animalsMap := make(map[string]string)
	animalsMap["dog"] = "Fido"
	animalsMap["cat"] = "Fluffy"
	for animalType, animal := range animalsMap {
		log.Println(animalType, animal)
	}

	// Loop a string
	var firstLine = "Once upon a midnight dreary"
	firstLine = "x"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	// Loop a custom object
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{FirstName: "Mary", LastName: "Jones", Email: "mary@jones.com", Age: 20})
	users = append(users, User{FirstName: "Sally", LastName: "Brown", Email: "sally@smith.com", Age: 45})
	users = append(users, User{FirstName: "Alex", LastName: "Anderson", Email: "alex@smith.com", Age: 17})
	for _, u := range users {
		log.Println(u.FirstName, u.LastName, u.Email, u.Age)
	}

}
