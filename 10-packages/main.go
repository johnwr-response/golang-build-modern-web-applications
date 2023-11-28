package main

import (
	"github.com/tsawler/myniceprogram/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	myVar.TypeNumber = 1
	log.Println(myVar.TypeName)
}
