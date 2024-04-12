package main

import "fmt"

func main() {
	// 	var a int

	// 	var p *int

	// 	a = 10

	// 	p = &a

	// 	*p = 20

	// 	fmt.Println(*p)

	// var a = 10
	// var b int
	// var p *int

	// b = a
	// p = &a

	// *p = 20

	// fmt.Println(a, b)

	// 	var a int = 10
	// 	add12pointer(&a)
	// 	fmt.Println(a)

	var numbers = []int{1, 2, 3}

	fmt.Println(numbers)
	changeValue(numbers)
	fmt.Println(numbers)

}

func add12(x int) { // pass by value
	x += 12
}

func add12pointer(x *int) { // pass by reference
	*x = *x + 12
}

func changeValue(numbers []int) {
	numbers[0] = 1000
}
