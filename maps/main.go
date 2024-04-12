package main

import "fmt"

func main() {
	// Maps: Python dict olarak düşünebilirsiniz - key, value şeklinde verileri saklıyor.

	// var names map[string]int
	// names = make(map[string]int, 0)

	// names["Mustafa"] = 1
	// names["Kaan"] = 2
	// names["Demir"] = 3

	// fmt.Println(names)
	// fmt.Println(names["1"])

	names := map[string]int{
		"Mustafa": 1,
		"Kaan":    2,
		"Demir":   3,
	}

	fmt.Println(names)
	delete(names, "Mustafa")
	fmt.Println(names)

}
