package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	aaa := map[int]string{
		123: "aaa",
		456: "bbb",
		789: "ccc",
	}

	fmt.Println(aaa[123123])

	_, ok := aaa[123]
	fmt.Printf("ok %#v\n", ok)

	if val, ok := aaa[123]; ok {
		fmt.Printf("val %#v\n", val)
	}

	aaa[456] = "zzz"
	fmt.Println(aaa)

	aaa[789] = "xxx"
	fmt.Println(aaa)
}
