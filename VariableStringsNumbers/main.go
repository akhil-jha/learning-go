package main

import (
	"fmt"
	"reflect"
)

func main() {
	// preferred camelCase in variable names
	var MyFirstGoString = "apples"
	fmt.Println(MyFirstGoString, reflect.TypeOf(MyFirstGoString))

	var MySecondGoString string = "banana"
	fmt.Println(MySecondGoString, reflect.TypeOf(MySecondGoString))

	var NumberVar = 100
	var IntVar int = 50 // There are variations. int, int8, int16, int32, int64 which tells the range of allowed number.
	fmt.Println(NumberVar, reflect.TypeOf(NumberVar), IntVar, reflect.TypeOf(IntVar))

	var UnsignedNumber uint = 111 // Similar to int, it also has variations
	fmt.Println(UnsignedNumber, reflect.TypeOf(UnsignedNumber))

	// var UnsignedNumber uint = -111
	// fmt.Println(UnsignedNum  er, reflect.TypeOf(UnsignedNumber))

	// var UnsignedOverflow uint = 11111111111111111111111111
	// fmt.Println(UnsignedOverflow, reflect.TypeOf(UnsignedOverflow))

	// https://go.dev/ref/spec#Numeric_types

	Foo := "bar"
	fmt.Println(Foo)

}
