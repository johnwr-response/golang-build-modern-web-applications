package main

import "log"

func main() {

	// Simple if bool
	var isTrue bool
	isTrue = true
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	// Simple if string
	cat := "cat2"
	if cat == "cat" {
		log.Println("cat is cat")
	} else {
		log.Println("cat is not cat")
	}

	// Simple if int
	myNum := 100
	isTrue = false
	if myNum > 99 && isTrue {
		log.Println("myNum is greater than 99 and isTrue is true")
	} else if myNum < 100 && isTrue {
		log.Println("myNum is greater than 100 and isTrue is true")
	} else if myNum == 101 || isTrue {
		log.Println("myNum is 101 and isTrue is true")
	} else if myNum > 1000 && isTrue == false {
		log.Println("myNum is greater than 999 and isTrue is not true")
	}

	// Simple switch
	myVar := "cat"
	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is set to something else")
	}

}
