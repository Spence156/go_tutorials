package main

import (
	"fmt"
	"time"
)

/// Arrays, slices, maps and loops
/// https://www.youtube.com/watch?v=8uiZC0l4Ajw

func main() {

	/// Arrays
	var intArr [3]int32 // varName [Length] datatype
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	/// Changing array values
	fmt.Println(intArr[1])
	intArr[1] = 123 // Changes the value of index 1 to 123
	fmt.Println(intArr[1])

	var intArr1 [3]int32 = [3]int32{1, 2, 3} // Initialse the values of an Array and declare it
	fmt.Println(intArr1)

	intArr2 := [3]int32{1, 2, 3} // Initialse the values of an Array and declare it
	fmt.Println(intArr2)

	intArr3 := [...]int32{1, 2, 3} // The compiler works out the max length of the array from the values passed to it (I.e. 3)
	fmt.Println(intArr3)

	/// Slices
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) //Adds a new value to the array (Either creates a new array and moves the data to it (If the length has exceeded it) or if there is capacity left in the array it adds a value to it)
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{8, 9} // Append new values to an existing array in bulk
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8) // Automatically creates a array with the number of values and the max number (E.g. 3 is the list of values to create and 8 would be the max length before expanding)
	fmt.Println(intSlice3)

	/// Maps
	/// Set of key value pairs

	var myMap map[string]uint8 = make(map[string]uint8) //This would create a map with a key being string type and value being an unsigned int8
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45} // Creates a map and assigns key/values
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])  // Gets the value of the Key named "Adam"
	fmt.Println(myMap2["Jason"]) // If you specify a value which doesn't exist you will get the data type default value (In this case it would be the default value of uint8 which would be 0)

	var age, ok = myMap2["Jason"] // This would allow you to better handle if the name doesn't exist instead of getting the default value and using it...
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	delete(myMap2, "Adam")   // Delete's a key/value from a Map.
	age, ok = myMap2["Adam"] // This would allow you to better handle if the name doesn't exist instead of getting the default value and using it...
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// Loops
	// Can loop through an array, map or slice

	// loop through a map
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	// loop through an array/slice
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// while loop
	var i int = 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i < 10; i++ { // Condensed While loop, 1st item is initialize variable, then the condition and then the post (I.e. what get's executed)
		fmt.Println(i)
	}

	/// Speed comparison for predefined length arrays vs dynamic ones...

	var n int = 1000000
	var testSlice = []int{}            // Empty array but no defined length
	var testSlice2 = make([]int, 0, n) // Empty but with defined length (1,000,000)
	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n)) // should be around 3x faster...

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
