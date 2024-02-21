package main

import (
	"fmt"
	"strings"
)

/// Strings runes and bytes
/// https://www.youtube.com/watch?v=8uiZC0l4Ajw

func main() {

	var myString = "résumé"
	var indexed = myString[0]
	fmt.Println(myString)
	fmt.Println(indexed) // Prints the utf-8 character number

	fmt.Printf("%v, %T\n", indexed, indexed) // Prints the value and the data type

	for i, v := range myString { // the é is 2 bytes which is why byte blocks 2 and 7 are missed as they are 233 as it is a special character and range is working in byte blocks not characters.
		fmt.Println(i, v)
	}

	var myString1 = []rune("résumé") // casts the string to a rune
	var indexed1 = myString1[1]
	fmt.Println(myString1)
	fmt.Println(indexed1) // Prints the utf-8 character number

	fmt.Printf("%v, %T\n", indexed1, indexed1) // Prints the value and the data type

	for i, v := range myString1 { // now loops through each character as it uses a "rune" instead of bytes
		fmt.Println(i, v)
	}

	// String building

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""

	// Using range (Not very efficient but it works)
	for n := range strSlice {
		catStr += strSlice[n]
	}
	fmt.Printf("\n%v", catStr)

	// Uses String Build package built into Go to build a string much more efficient
	var strBuilder strings.Builder

	for n := range strSlice {
		strBuilder.WriteString(strSlice[n])
	}
	catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)

}
