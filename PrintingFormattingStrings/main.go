package main

import (
	"fmt"
)

func main() {
	var tempA string = "I am a string"
	var numA = 100
	numB := 11
	var emptyVar string
	emptyVar = "foobar"

	fmt.Println("These are the variables I declared above:", tempA, numA, numB, "\nI can do this as well: ", emptyVar)

	fmt.Print("Hellonn") // Doesn't add new line on output
	fmt.Println("Hellonn")

	fmt.Printf("I can do this - %v - as well - %v - here which is cool - %v - too I guess\n", tempA, numA, numB)
	fmt.Printf("I can do this - %q - as well - %q - here which is cool - %q - too I guess\n", tempA, numA, numB) // %q only works for strings
	fmt.Printf("This is a float formatted string %0.3f okay?", 132.312334)

	// save formatted string
	var newString = fmt.Sprintf("I can do this - %v - as well - %v - here which is cool - %v - too I guess\n", tempA, numA, numB)
	fmt.Println(newString)
}
