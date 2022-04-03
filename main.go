package main

import "fmt"

// func updateName(n string) {
// 	n = "wedge"
// }

func updateName(n *string) {
	*n = "wedge"
}

func main() {
	name := "tifa"

	// updateName(name)
	// fmt.Println(name)

	// fmt.Println("memory address of name is:", &name)

	m := &name
	// fmt.Println("memory address:", m)
	// fmt.Println("value at memory address:", *m)

	updateName(m)
	fmt.Printf("m %#v\n", *m)
	fmt.Println(name)

}
