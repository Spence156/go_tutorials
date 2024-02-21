package main

import (
	"fmt"
)

/// Pointers
/// https://www.youtube.com/watch?v=8uiZC0l4Ajw
// Way of referencing existing memory objects to save memory instead of having duplicate memory objects passed around into functions etc...

func square(thing2 [5]float64) [5](float64) {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2) // memory address is different to thing1 meaning double the memory is consumed but the array is different so can be updated uniquely in the function
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return thing2
}

func square1(thing2 *[5]float64) [5](float64) {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", thing2) // memory address is the samne as thing1 meaning effectively the same array is re-used reducing memory (But does mean edits are applied to the same array)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}

func main() {
	var p *int32 = new(int32) //Initialize a pointer
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
	*p = 10 // Reference and update a value of a pointer

	p = &i // The value of p = the value of i
	*p = 1
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)

	var k int32 = 2
	i = k

	var thing1 = [5]float64{1, 2, 3, 4, 5} // Create a float64 array with 5 items
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is@ %v", result)

	var result1 = square1(&thing1)
	fmt.Printf("\nThe result is@ %v", result1)
}
