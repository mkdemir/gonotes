package main

import "fmt"

func main() {

	var customer1 = Customer{
		id:   1,
		name: "Mustafa Kaan Demir",
		age:  24,
		address: Address{
			city:     "İstanul",
			district: "Bakırköy",
		},
	}

	// var customer2 = Customer{
	// 	id:   2,
	// 	name: "Kaan Demir",
	// 	age:  30,
	// }

	// customer1.age = 31
	// fmt.Println(customer1)
	// fmt.Println(customer2)

	customer1.ToString()
	// changeName(&customer1)
	customer1.ToString()

	customer1.changeName("mkdemir1")
	customer1.ToString()

}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}

type Address struct {
	city     string
	district string
}

func (custoomer Customer) ToString() {
	fmt.Printf("Name: %s, Age: %d\n", custoomer.name, custoomer.age)
}

// func changeName(customer *Customer) {
// 	customer.name = "mkdemir"
// }

func (customer *Customer) changeName(newName string) {
	customer.name = newName
}
