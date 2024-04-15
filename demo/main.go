package main

import "fmt"

func main() {
	/*
	   GİRİŞ:
	   Modül: Birçok paketten oluşan yapılara denir.
	   Modül oluşturma:
	   	- go mod init github.com/mkdemir/onevideo
	   Gerekli paketi yükleme:
	   	- go get github.com/twpayne/go-geom
	   Modül bağımlılıklarını temizleme:
	   	- go mod tidy
	   Ana programı çalıştırma:
	   	- go run main.go
	   go.sum dosyası: Bir history gibi
	*/

	// Değişkenler ve Tip Dönüşümleri:

	// var productName string
	// var quantity int
	// var discount = 10.4
	// isInStock := true

	// productName = "Megalogon"
	// quantity = 10

	// fmt.Println("Product Name Type:", reflect.TypeOf(productName))
	// fmt.Printf("Product Name: %s\nQuantity: %d\nDiscount: %f\nIs in Stock: %t\n", productName, quantity, discount, isInStock)

	// Diziler ve Dilimler (Arrays and Slices):

	// // Array - Fixed
	// var names [4]string
	// names[0] = "Mustafa"
	// names[1] = "Kaan"
	// names[2] = "Demir"
	// names[3] = "mkdemir"

	// fmt.Println(names)

	// var name1 = [3]string{"Mustafa", "Kaan", "Demir"}
	// fmt.Println(reflect.TypeOf(name1))
	// fmt.Printf("name1: %s\n", name1)

	// // Slice
	// var names1 = []string{"test1", "test2", "test3"}
	// names1 = append(names1, "Kaan")
	// fmt.Println(names1)
	// fmt.Println(names1[0:2])

	// Koşullu İfadeler (Conditional Statements):

	// age := 27

	// if age >= 21 {
	// 	fmt.Println("Next")
	// } else if age >= 18 {
	// 	fmt.Println("Wait")
	// } else {
	// 	fmt.Println("Stop")
	// }

	// Döngüler (Loops):

	// Örnek1:
	// index := 1

	// for index <= 10 {
	// 	fmt.Println(index)
	// 	index++
	// }

	// Örnek2:
	// var total int
	// for index := 1; index <= 10; index++ {
	// 	total += index
	// }
	// fmt.Println(total)

	// // Örnek 3:
	// index := 0
	// var names = [3]string{"test1", "test2", "test3"}

	// for index < len(names) {
	// 	fmt.Println(names[index])
	// 	index++
	// }

	// // Örnek 4:
	// for index := 0; index < 10; index++ {
	// 	if index == 5 {
	// 		continue
	// 	}
	// 	if index == 8 {
	// 		break
	// 	}
	// 	fmt.Println(index)
	// }

	// Haritalar (Maps): Python'daki dictionary olarak düşünebilirsiniz - key, value şeklinde verileri saklıyor.

	// // Örnek 1:
	// names := map[string]int{
	// 	"Mustafa": 1,
	// 	"Kaan":    2,
	// 	"Demir":   3,
	// }
	// delete(names, "Mustafa")

	// fmt.Println(names)

	// // Örnek 2:
	// var names map[string]int
	// names = make(map[string]int, 0)
	// names["Mustafa"] = 1
	// names["Kaan"] = 2
	// names["Demir"] = 3

	// fmt.Println(names)

	// Foreach Loops

	// // Dilim üzerinde dolaşma
	// var numbers = []int{1, 2, 3, 4}
	// for _, value := range numbers {
	// 	fmt.Println(value)
	// }

	// // String üzerinde dolaşma
	// var language = "Go Lang"
	// for _, character := range language {
	// 	fmt.Printf("Character: %c\n", character)
	// }

	// // Harita üzerinde dolaşma
	// var names = map[string]int{
	// 	"Mustafa": 20,
	// 	"Kaan":    30,
	// 	"Demir":   40,
	// }

	// for key, value := range names {
	// 	fmt.Printf("Key: %s, Value: %d\n", key, value)
	// }

	// FONKSİYONLAR:

	//total := add(10, 20)
	//fmt.Println(total)

	// total, difference, multiply := calculation(10, 20)
	// fmt.Println(total, difference, multiply)

	// var numbers = []int{10, 20, 30, 40, 50}
	// fmt.Println(sum(numbers))

	fmt.Println(sum(3, 4, 5))
	fmt.Println(sum(3, 4, 5, 6))

	// Pointers:
	// Değişken tanımları ve işlemleri
	// var a int
	// var p *int
	// a = 10
	// p = &a
	// *p = 20
	// fmt.Println(*p)

	// Değişkenler ve işlemler
	// var a = 10
	// var b int
	// var p *int
	// b = a
	// p = &a
	// *p = 20
	// fmt.Println(a, b)

	// Değer geçirme işlemleri
	var a int = 10
	add12pointer(&a)
	fmt.Println(a)

	// // Dilim üzerinde değişiklik yapma
	// var numbers = []int{1, 2, 3}
	// fmt.Println(numbers)
	// changeValue(numbers)
	// fmt.Println(numbers)
}

//#################################################
// Fonksiyonlar:
//#################################################
// func add(x int, y int) int {
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

//#################################################

//#################################################
// Pointer:
//#################################################

// add12, değeri 12 arttıran bir fonksiyondur. Değer geçirilirken kopya alınır (pass by value).
func add12(x int) {
	x += 12
}

// add12pointer, değeri 12 arttıran bir fonksiyondur. Değer geçirilirken referans kullanılır (pass by reference).
func add12pointer(x *int) {
	*x = *x + 12
}

// changeValue, bir dilim üzerinde değişiklik yapar.
func changeValue(numbers []int) {
	numbers[0] = 1000
}

//#################################################
