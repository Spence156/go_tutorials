package main

import (
	"fmt"
	"unicode/utf8"
)

/// Variables and basic data types
/// https://www.youtube.com/watch?v=8uiZC0l4Ajw

func main() {

	/// Numbers

	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32) // Need to cast data types to same data type in order to perform operations on them
	fmt.Println(result)

	var result1 float32 = floatNum32 + float32(intNum) // Need to cast data types to same data type in order to perform operations on them
	fmt.Println(result1)

	/// String
	var myString string = "Hello \nWorld" // Prints onto new line
	var myString1 string = "Hello World"
	var myString2 string = "Hello" + " " + "World" // Concatante String

	fmt.Println(myString)
	fmt.Println(myString1)
	fmt.Println(myString2)

	fmt.Println(len("γ"))                    // Number of bytes not the number of characters (If using special chacaters this will be incorrect)
	fmt.Println(utf8.RuneCountInString("γ")) // Number of bytes not the number of characters (Counts the characters correctly including special characters)

	// Bool
	var myBool bool = false
	fmt.Println(myBool)

	var intNum3 int // Can create variable without specifying the value in which case the default value is applied (int = 0)
	fmt.Println(intNum3)

	myVar := "text" /// := is shorthand for declare variable
	fmt.Println(myVar)

	/// Constants (Can be set once and read many (I.e. can never change once set))
	const myConst string = "const value"
	fmt.Println(myConst)
}
