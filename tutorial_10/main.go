/// Generics
/// https://www.youtube.com/watch?v=8uiZC0l4Ajw
// Working with Generics to allow functions to recieve different data types for the same parameter

package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice))

	var float64Slice = []float64{1, 2, 3}
	fmt.Println(sumSlice[float64](float64Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T { // Allows int, float32 or float64
	var sum T
	for _, v := range slice {
		sum += v // Actions can only be performed if they support the data types (E.g. sum cannot be performed on a string)
	}
	return sum
}
