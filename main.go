package main

import "fmt"

func main() {
	// string
	var nameOne string = "a"
	var nameTwo = "b"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "aa"
	nameThree = "cc"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "d"

	fmt.Println(nameFour)

	// int
	var numberOne int = 1
	var numberTwo = 2
	numberThree := 3

	fmt.Println(numberOne, numberTwo, numberThree)

	// bits & memory
	var numberFour int8 = 25
	var numberFive int8 = -128
	var numberSix uint16 = 256

	fmt.Println(numberFour, numberFive, numberSix)

	var aOne float32 = 25.12
	var aTwo float64 = 123123121231231231233.123
	aThree := 1.5

	fmt.Println(aOne, aTwo, aThree)
}
