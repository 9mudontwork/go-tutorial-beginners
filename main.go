package main

import "fmt"

func main() {
	x := 0
	for x < 10 {
		fmt.Printf("x %#v\n", x)
		x++
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i %#v\n", i)
	}

	names := []string{"John", "Paul", "George", "Ringo"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, name := range names {
		fmt.Println(i, name)
	}

	for _, name := range names {
		fmt.Println(name)
		name = "New Name"
	}

	fmt.Println(names)
}
