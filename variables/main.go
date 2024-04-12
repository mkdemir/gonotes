package main

import (
	"fmt"
)

func main() {
	// var productName string
	// var quantity int
	// var discount float32
	// var isInStock bool

	// productName = "Go Lang Dersleri"
	// quantity = 10
	// discount = 0.37
	// isInStock = true

	// fmt.Println(productName, reflect.TypeOf(productName))
	// fmt.Println(quantity, reflect.TypeOf(quantity))
	// fmt.Println(discount, reflect.TypeOf(discount))
	// fmt.Println(isInStock, reflect.TypeOf(isInStock))

	// var productName = "Go Lang Dersleri"
	// fmt.Println(productName, reflect.TypeOf(productName))

	// productName := "Go Lang Dersleri"
	// fmt.Println(productName)

	// var quantity int64 = 10
	// fmt.Println(quantity, reflect.TypeOf(quantity))

	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Go Lang Dersleri"
	quantity = 10
	discount = 0.37
	isInStock = true

	// fmt.Println("Product Name: ", productName, "Quantity: ", quantity, "Discount: ", discount, "Is in Stock:", isInStock)

	fmt.Printf("Product Name: %v,\nQuantity: %v,\nDiscount: %v,\nIs in Stock: %v \n", productName, quantity, discount, isInStock)
}
