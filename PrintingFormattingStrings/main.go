package main

import (
	"fmt"
)

func main() {
	var tempA string = "String"
	var numA = 100
	numB := 11
	var emptyVar string
	emptyVar = "emptyVar"

	fmt.Println("These are the variables I declared above:", tempA, numA, numB, "\nI can do this as well: ", emptyVar)

	fmt.Print("Hellonn") // Doesn't add new line on output
	fmt.Println("Hellonn\n\n\n")

	fmt.Printf("percentage V values -> %v - %v - %v", tempA, numA, numB)
	fmt.Printf("\n\nMix value -> %q - %q - %q -> q only works for strings", tempA, numA, numB,) // %q only works for strings
	fmt.Printf("\n\nFloat string %0.3f: ", 132.312334)

	// save formatted string
	var newString = fmt.Sprintf("\n\n %v %v %v", tempA, numA, numB)
	fmt.Println(newString)
}