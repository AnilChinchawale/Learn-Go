package main

import "fmt"

type contactinfo struct {
	emailid   string
	postalZip int
}

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

func main() {

	anil := person{
		firstName: "Anil",
		lastName:  "Anand",
		contact: contactinfo{
			emailid:   "Anil@xinfin.org",
			postalZip: 400067,
		},
	}
	fmt.Printf("%+v", anil)
}
