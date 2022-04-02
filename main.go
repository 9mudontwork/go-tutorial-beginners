package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hello World"

	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Goodbye"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{17, 23, 97, 21, 5, 12, 34, 2, 1, 0}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 5)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "luigi", "toad", "peach", "bowser"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "peach"))
}
