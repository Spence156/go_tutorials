package main

import "fmt"

/// Structs and interfaces
/// https://www.youtube.com/watch?v=8uiZC0l4Ajw

type gasEngine struct { // Reusable
	mpg     uint8
	gallons uint8
	owner
	int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine // Creates variable but with default data type values (I.e. "0")
	fmt.Println(myEngine.mpg, myEngine.gallons)

	var myEngine1 gasEngine = gasEngine{mpg: 25, gallons: 15, owner: owner{"Alex"}, int: 10} // Defines the values in the struct (If fields names are not supplied it works in the order the struct is defined in)
	fmt.Println(myEngine1.mpg, myEngine1.gallons, myEngine1.owner, myEngine1.int)

	var myEngine2 = struct { // Anonomous Struct ( Not re-usable ) but defined inline with the code instead of ourside the main function
		mpg     uint8
		gallons uint8
	}{25, 15}

	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	fmt.Printf("Total miles left in tank: %v\n", myEngine1.milesLeft())

	var myEnine3 electricEngine = electricEngine{25, 15}

	canMakeIt(myEnine3, 50)
}
