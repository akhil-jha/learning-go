package main

import (
	"fmt"
)

func main() {
	var arrayLength int = 10
	dynamicArray := make([]int, arrayLength) // slicing
	dynamicArray[9] = 1
	dynamicArray = append(dynamicArray, 111)

	var newArray = [5]int{10, 12, 13, 14, 16}
	newArray[0] = 11
	// newArray = append(newArray, 99) // Will fail becuase append() requires a slice

	var customArray = []int{123, 234, 35}
	// customArray[3] = 1 // index out of range
	customArray = append(customArray, 99)
	rangeArray := customArray[1:]

	fmt.Println(newArray, customArray, "\n", dynamicArray, "\n", rangeArray)
}
