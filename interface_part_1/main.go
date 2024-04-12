package main

import "fmt"

type Book struct {
	desi int
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

type Electronic struct {
	desi int
}

func (electronic Electronic) ShippingCost() int {
	return 10 + electronic.desi*3
}

func main() {
	// book1 := &Book{desi: 10}
	// book2 := &Book{desi: 10}

	// fmt.Println(book1.ShippingCost())
	// fmt.Println(book2.ShippingCost())

	// books := []Book{
	// 	Book{desi: 10},
	// 	Book{desi: 20},
	// }
	// fmt.Println(calculateTotalShippingCost(books))

	// electronic1 := &Electronic{desi: 20}
	// fmt.Println(electronic1.ShippingCost())

	electronics := []Electronic{{desi: 10}, {desi: 20}}
	fmt.Println(calculateTotalShippingCostOfElectronics(electronics))
}

func calculateTotalShippingCostOfBooks(books []Book) int {
	total := 0

	for _, book := range books {
		total += book.ShippingCost()
	}

	return total
}

func calculateTotalShippingCostOfElectronics(electronics []Electronic) int {
	total := 0

	for _, book := range electronics {
		total += book.ShippingCost()
	}

	return total
}
