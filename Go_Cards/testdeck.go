package main

import (
	"fmt"
	"io/ioutil"
)

type testdeck = []string

func (d deck) testSave(testfilename string) error {

	return ioutil.WriteFile(testfilename, []byte(d.toString()), 0666)

}

func (d deck) testRead(testfilename string) ([]byte, error) {

	// fmt.Println(ioutil.ReadFile(testfilename))
	sc, err := ioutil.ReadFile(testfilename)

	if err != nil {
		fmt.Println("There is something wrong")
		fmt.Println(err)
	}

	fmt.Println(sc)

	return ioutil.ReadFile(testfilename)
	// return fmt.Println(ioutil.ReadFile(testfilename))

}
