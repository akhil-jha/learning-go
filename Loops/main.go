package main

import (
	"fmt"
)

func main() {
	// Simple for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("Value of i is %v\n", i)
	}

	// // for look acting as a while loop
	i := 1
	for i <= 10 {
		fmt.Printf("2 * %v = %v\n", i, (2 * i))
		i = i + 1
	}

	// for range loop
	letters := []string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz"}

	for i, j := range letters { // this has to have 2 iterators. "i" will store the index, "j" will store the value
		fmt.Println(i, j)
	}

	// for loop - map (dictionary)

	numbers := map[int]string{
		11: "one",
		22: "two",
		33: "three",
	}

	for key, value := range numbers {
		fmt.Println(key, value)
	}

	// for loop - string
	normalString := "This is a normal string"
	for i, j := range normalString {
		fmt.Printf("%v : %c\n", i, j)
	}
}
