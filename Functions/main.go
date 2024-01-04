package main

import (
	"fmt"
	"strings"
)

// func testFunction(firstNumber, secondNumebr int, test string) (int, int, string) {
// 	fmt.Printf("%v\n %v\n %v\n", test, firstNumber, secondNumebr)
// 	return firstNumber, secondNumebr, test
// }

// func main() {
// 	number, secondNumber, test := testFunction(11, 100, "This is a test sting passed in as function")
// 	fmt.Println(test, number, secondNumber)
// }

// Call by value -
// func swapNumbers(firstNumber, secondNumber int) {
// 	var tempNumber int
// 	tempNumber = firstNumber
// 	firstNumber = secondNumber
// 	secondNumber = tempNumber
// 	fmt.Printf("After Swaping\nFirst number: %v\nSecond number: %v\n\n", firstNumber, secondNumber)
// }

// func main() {
// 	firstNumber := 1
// 	secondNumber := 2
// 	swapNumbers(firstNumber, secondNumber)
// 	fmt.Printf("First number: %v\nSecond number: %v\n", firstNumber, secondNumber)
// }

// call by refernce
func swapNumbers(firstNumber, secondNumber *int) {
	var tempNumber int
	tempNumber = *firstNumber
	*firstNumber = *secondNumber
	*secondNumber = tempNumber
}

func main() {
	firstNumber := 1
	secondNumber := 2
	fmt.Printf("Before swapping\nFirst number: %v\nSecond number: %v\n", firstNumber, secondNumber)
	fmt.Println(strings.Repeat("-", 10))
	swapNumbers(&firstNumber, &secondNumber)
	fmt.Printf("After swapping\nFirst number: %v\nSecond number: %v\n", firstNumber, secondNumber)
}
