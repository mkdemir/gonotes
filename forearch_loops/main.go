package main

import "fmt"

func main() {
	// var numbers = []int{1, 2, 3, 4}

	// // for index := 0; index < len(numbers); index++ {
	// // 	fmt.Println(numbers[index])
	// // }

	// // for index, value := range numbers {
	// // 	fmt.Println(index, value)

	// // }

	// for _, value := range numbers {
	// 	fmt.Println(value)
	// }

	// var language = "Go Lang"
	// for _, character := range language {
	// 	fmt.Printf("Character %c\n", character)
	// }

	var names = map[string]int {
		"Mustafa": 20,
		"Kaan": 30,
		"Demir": 40,
	}

	for key, value := range names {
		fmt.Printf("Key: %s Value: %d", key, value)
	}

}
