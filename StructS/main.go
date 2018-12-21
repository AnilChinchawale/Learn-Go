package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	// alex := person{}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
