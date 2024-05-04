package main

import "fmt"


func main () {

	var firstName, lastName string
	var a, b float32
	a = 1
	b = 2

	firstName = "kato"
	lastName = "byantalo"

	if(firstName == "kato" && lastName == "byantalo"){
		fmt.Printf("My name is %s %s \n", firstName, lastName)
		var c = (a + b)
		fmt.Printf("The total is %f", c)
	}else{
		fmt.Println("not way man")
	}

}
