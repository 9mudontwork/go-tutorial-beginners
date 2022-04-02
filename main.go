package main

import "fmt"

func main() {
	a := 11
	b := "22"

	// print
	fmt.Print("Hello, ")
	fmt.Print("Hello \n")
	fmt.Print("Hello \n")

	// print line
	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
	fmt.Println("Hello 2", a, " and ", b)

	// print format string
	fmt.Printf("Hello %v and %v \n", a, b)
	fmt.Printf("Hello %q and %q \n", a, b)
	fmt.Printf("a is type of %T \n", a)
	fmt.Printf("my b is %0.2f \n", 123.12)

	// sprintf save format string
	var str = fmt.Sprintf("my a is %v and b is %v \n", a, b)
	fmt.Println("save string is ", str)
}
