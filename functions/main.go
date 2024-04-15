package main

import (
	"fmt"
)

func main() {
	//total := add(10, 20)

	//fmt.Println(total)

	// total, difference, multiply := calculation(10, 20)

	// fmt.Println(total, difference, multiply)

	// var numbers = []int{10, 20, 30, 40, 50}

	// fmt.Println(sum(numbers))

	fmt.Println(sum(3, 4, 5))
	fmt.Println(sum(3, 4, 5, 6))

}

// func add(x int, y int) int {
// 	// fmt.Println(x + y)

// 	return x + y
// }

// func calculation(x int, y int) (int, int, int) {
// 	return x + y, x - y, x * y
// }

// func sum(numbers []int) int {

// 	var sum int

// 	for _, value := range numbers {
// 		sum += value
// 	}

// 	return sum
// }

// func sum(x int, y int, z int) int {
// 	return x + y + z
// }

func sum(numbers ...int) int {
	var sum int
	for _, value := range numbers {
		sum += value
	}
	return sum
}
