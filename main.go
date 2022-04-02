package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{10, 20, 30}
	var ages = [3]int{10, 20, 30}

	names := [4]string{"John", "Jane", "Joe", "Peach"}
	names[1] = "Doe"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{10, 20, 30}
	scores[2] = 40
	scores = append(scores, 50)

	fmt.Println(scores, len(scores))

	// slices ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Koopa")
	fmt.Println(rangeOne)

}
