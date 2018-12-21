package main

import "fmt"

type contactinfo struct {
	emailid   string
	postalZip int
}

type person struct {
	firstName string
	lastName  string
	contactinfo
}

func main() {

	anil := person{
		firstName: "Anil",
		lastName:  "Anand",
		contactinfo: contactinfo{
			emailid:   "Anil@xinfin.org",
			postalZip: 400067,
		},
	}
	anil.updateData("Neel")
	anil.print()
}

func (p person) print() {

	fmt.Printf("%+v", p)
}

func (p person) updateData(firstName string) {
	p.firstName = firstName
}
