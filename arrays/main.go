package main

import "fmt"

func main() {
	// Array
	// Fixed

	// var names [3]string

	// names[0] = "Mustafa"
	// names[1] = "Kaan"
	// names[2] = "Demir"
	// fmt.Println(names) // Boş bir array boş bir string'tir.

	// var names = [4]string{"Mustafa", "Kaan", "Demir", "Ömer"} // Sabit boyutlu bir indeks

	// fmt.Println(names[0:2]) // Aray ile parçalama

	// Slice

	var names = []string{"Mustafa", "Kaan", "Demir"}
	names = append(names, "Ömer")
	fmt.Println(names)
}
